package main

// Extracted from https://github.com/golang/dep/

type rawLock struct {
	SolveMeta solveMeta          `toml:"solve-meta"`
	Projects  []rawLockedProject `toml:"projects"`
}

type solveMeta struct {
	AnalyzerName    string   `toml:"analyzer-name"`
	AnalyzerVersion int      `toml:"analyzer-version"`
	SolverName      string   `toml:"solver-name"`
	SolverVersion   int      `toml:"solver-version"`
	InputImports    []string `toml:"input-imports"`
}

type rawLockedProject struct {
	Name      string   `toml:"name"`
	Branch    string   `toml:"branch,omitempty"`
	Revision  string   `toml:"revision"`
	Version   string   `toml:"version,omitempty"`
	Source    string   `toml:"source,omitempty"`
	Packages  []string `toml:"packages"`
	PruneOpts string   `toml:"pruneopts"`
	Digest    string   `toml:"digest"`
}
