# dep to go modules with sugar

[![GitHub release](https://img.shields.io/github/release/ldez/deptomod.svg)](https://github.com/ldez/deptomod/releases/latest)
[![Build Status](https://travis-ci.com/ldez/deptomod.svg?branch=master)](https://travis-ci.com/ldez/deptomod)

Simple tool to convert a project from dep to go modules, respect `source` attribute.

If you appreciate this project:

[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg?style=for-the-badge)](https://saythanks.io/to/ldez)

## How to use

I recommend to export the status instead of using the `Gopkg.lock` from the project:

```bash
dep status -lock > Gopkg.lock.toml
```

Command line:

```yaml
Enhanced migration from dep to go modules.

Usage:
  deptomod [flags]

Flags:
  -h, --help            help for deptomod
  -i, --input string    The input directory. (default "./fixtures")
  -m, --module string   The future module name. (default "github.com/user/repo")
  -o, --output string   The output file. (default "./go.mod.txt")
```
