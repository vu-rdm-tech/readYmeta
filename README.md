# yoda-metadata-toolkit
Software for reading, writing and manipulating Yoda metadata files.

## readYmeta 
Reads a Yoda metadata JSON file and writes it to PDF Tested with Yoda 1.7.
- [readYmeta](https://github.com/vu-rdm-tech/readYmeta/tree/main) is written in Go - main build status: [![Go](https://github.com/vu-rdm-tech/readYmeta/actions/workflows/go.yml/badge.svg)](https://github.com/vu-rdm-tech/readYmeta/actions/workflows/go.yml)

### Instructions for use
readYmeta.go reading and converting Yoda metadata from JSON to PDF formats

Usage: readYmeta.exe <yoda metadata file> filename can include a path. If no file is specified "yoda-metadata.json" is assumed in current directory.

Output: A PDF file containing the Yoda metadata with missing attributes highlighted. output <filename>.pdf is formed from input <filename>.json, defualts to current directory.

- Author: Brett G. Olivier PhD
- email: @bgoli
- licence: BSD 3 Clause
- version: 0.7-beta
- Date: 2022-08-22
(C) Brett G. Olivier, Vrije Universiteit Amsterdam, Amsterdam, The Netherlands, 2022.
