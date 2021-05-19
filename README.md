![CircleCI](https://img.shields.io/circleci/build/github/gucio321/d2d2s/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gucio321/d2d2s)](https://goreportcard.com/report/github.com/gucio321/d2d2s)
[![GoDoc](https://pkg.go.dev/badge/github.com/gucio321/d2d2s?utm_source=godoc)](https://pkg.go.dev/mod/github.com/gucio321/d2d2s)

## Description

This package contains a Diablo II save files (D2S)
decoder and encoder written in [golang](https://golang.org)

## Dependencies / Documentation

### D2S structure documentation:

*   [krisives/d2s-format](https://github.com/krisives/d2s-format)
*   [Diablo II Saved Game File Format](https://user.xmission.com/~trevin/DiabloIIv1.09\_File_Format.shtml)
*   [nokka/d2s](https://github.com/nokka/d2s) [(see details)](./docs/base.md)
*   [pairofdocs/d2s_edit_recalc](https://github.com/pairofdocs/d2s_edit_recalc)

### data reader

this project uses [Stream readers based on these used in OpenDiablo2 project](https://github.com/OpenDiablo2/OpenDiablo2)
for more informations, see [here](./docs/base.md)

## Status

for now (May 2021) the project is early in development.
However, it is possible to:

*   decode character save file into a (in a greater part) human-readable structure
*   encoding file

## Example

for some reasons, this project isn't able to create a whole new
file yet, but it is able to load modify and than encode D2S format

<details><summary>simple example</summary>

```golang
package main

import (
        "fmt"
        "ioutil"
        "log"

        "github.com/gucio321/d2d2s"
)

func main() {
        data, err := ioutil.ReadFile("/path/to/file.d2s")
        if err != nil {
                log.Fatal(err)
        }

        character, err := d2d2s.Unmarshal(data)
        if err != nil {
                log.Fatal(err)
        }

        // some edits

        newData, err := d2d2s.Encode(character)
        if err != nil {
                log.fatal(err)
        }

        ioutil.WriteFile("/path/to/new/file.d2s", newData, 0o600)
}
```

</details>

for more examples, see [here](./examples)

## Goals of the project

this programm should be able to decode and encode D2S files
