package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ldez/deptomod/goproxy"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

// Version application version.
var Version = "dev"

type modFile struct {
	Require []item
	Replace []item
}

type item struct {
	Name    string
	Source  string
	Version string
}

type config struct {
	ModuleName string
	InputDir   string
	Output     string
}

func main() {
	log.SetFlags(log.Lshortfile)

	cfg := config{}

	rootCmd := &cobra.Command{
		Use:     "deptomod",
		Short:   "Enhanced migration from dep to go modules.",
		Long:    `Enhanced migration from dep to go modules.`,
		Version: Version,
		RunE: func(_ *cobra.Command, _ []string) error {
			return run(cfg)
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&cfg.ModuleName, "module", "m", "github.com/user/repo", "The future module name.")
	flags.StringVarP(&cfg.InputDir, "input", "i", "./fixtures", "The input directory.")
	flags.StringVarP(&cfg.Output, "output", "o", "./go.mod.txt", "The output file.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	mod, err := convert(cfg.InputDir)
	if err != nil {
		return err
	}

	file, err := os.Create(cfg.Output)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	ew := &errWriter{w: file}

	ew.writef("module %s\n", cfg.ModuleName)
	ew.writeln()
	ew.writeln("go 1.13")
	ew.writeln()
	ew.writeln("require (")

	for _, v := range mod.Require {
		ew.writef("\t%s\t%s\n", v.Name, v.Version)
	}

	for _, v := range mod.Replace {
		ew.writef("\t%s\t%s\n", v.Name, "v0.0.0-00010101000000-000000000000")
	}

	ew.writeln(")")
	ew.writeln()

	if len(mod.Replace) != 0 {
		ew.writeln("replace (")

		for _, v := range mod.Replace {
			ew.writef("\t%s => %s %s\n", v.Name, v.Source, v.Version)
		}

		ew.writeln(")")
		ew.writeln()
	}

	if ew.err != nil {
		return ew.err
	}

	return nil
}

func convert(dir string) (modFile, error) {
	lock := rawLock{}
	err := readTomlFile(filepath.Join(dir, "Gopkg.lock"), &lock)
	if err != nil {
		return modFile{}, err
	}

	client := goproxy.NewClient("")

	mod := modFile{}
	for _, project := range lock.Projects {
		if project.Source == "" {
			version := getVersion(project)
			info, err := client.GetInfo(project.Name, version)
			if err != nil {
				fmt.Println(project.Name, err)
			} else {
				version = info.Version
			}

			mod.Require = append(mod.Require, item{
				Name:    project.Name,
				Version: version,
			})
		} else {
			version := getVersion(project)
			source := cleanSource(project.Source)

			info, err := client.GetInfo(source, version)
			if err != nil {
				fmt.Println(project.Name, project.Source, err)
			} else {
				version = info.Version
			}

			mod.Replace = append(mod.Replace, item{
				Name:    project.Name,
				Source:  source,
				Version: version,
			})
		}
	}

	return mod, nil
}

func readTomlFile(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	return toml.NewDecoder(file).Decode(data)
}

func getVersion(lock rawLockedProject) string {
	specials := []string{
		"github.com/transip/gotransip",
		"github.com/coreos/go-systemd",
		"github.com/gogo/protobuf",
	}

	if contains(specials, lock.Name) {
		return lock.Revision
	}

	if lock.Version != "" {
		return lock.Version
	}
	return lock.Revision
}

func contains(values []string, val string) bool {
	for _, v := range values {
		if v == val {
			return true
		}
	}
	return false
}

func cleanSource(src string) string {
	return strings.NewReplacer("https://", "", ".git", "").Replace(src)
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) writeln(a ...interface{}) {
	if ew.err != nil {
		return
	}

	_, ew.err = fmt.Fprintln(ew.w, a...)
}

func (ew *errWriter) writef(format string, a ...interface{}) {
	if ew.err != nil {
		return
	}

	_, ew.err = fmt.Fprintf(ew.w, format, a...)
}
