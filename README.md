# yoda-metadata-toolkit
Software for reading, writing and manipulating Yoda metadata files.

## readYmeta 
Reads a Yoda metadata JSON file and writes it to PDF  
- [readYmeta](https://github.com/bgoli/readYmeta2/tree/main) is written in Go 
- main build status: [![Go](https://github.com/bgoli/readYmeta2/actions/workflows/go.yml/badge.svg)](https://github.com/bgoli/readYmeta2/actions/workflows/go.yml)

## Usage 

### Windows 
`readYmeta.exe <filename>` 

### Linux 
`readYmeta <filename>` 

The filename can include a path specification. If no file is specified "yoda-metadata.json" is assumed as default filename using the current directory.

## Output 
A PDF file containing the Yoda metadata with missing attributes highlighted. <name>.pdf is formed from <name>.json, defaults to current directory.

## Admin stuff
- Author: Brett G. Olivier PhD
- email: @bgoli
- licence: BSD 3 Clause
- version: 0.8
- Date: 2023-01-25
(C) Brett G. Olivier, Vrije Universiteit Amsterdam, Amsterdam, The Netherlands, 2022.
