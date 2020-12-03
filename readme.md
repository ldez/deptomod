# dep to go modules with sugar

[![GitHub release](https://img.shields.io/github/release/ldez/deptomod.svg)](https://github.com/ldez/deptomod/releases/latest)
[![Build Status](https://github.com/ldez/deptomod/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/deptomod/actions)

A simple tool to convert a project from dep to go modules, respect `source` attribute.

If you appreciate this project:

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)

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
