![CircleCI](https://img.shields.io/circleci/build/github/gucio321/d2d2s/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gucio321/d2d2s)](https://goreportcard.com/report/github.com/gucio321/d2d2s)
[![GoDoc](https://pkg.go.dev/badge/github.com/gucio321/d2d2s?utm_source=godoc)](https://pkg.go.dev/mod/github.com/gucio321/d2d2s)
[![codecov](https://codecov.io/gh/gucio321/d2d2s/branch/master/graph/badge.svg)](https://codecov.io/gh/gucio321/d2d2s)

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

        "github.com/gucio321/d2d2s/pkg/d2s"
)

func main() {
        data, err := ioutil.ReadFile("/path/to/file.d2s")
        if err != nil {
                log.Fatal(err)
        }

        character, err := d2s.Unmarshal(data)
        if err != nil {
                log.Fatal(err)
        }

        // some edits

        newData, err := d2d2s.Encode(character)
        if err != nil {
                log.fatal(err)
        }

        if err := ioutil.WriteFile("/path/to/new/file.d2s", newData, 0o600); err != nil {
                log.Fatal("error writing file")
        }
}
```

</details>

for more examples, see [here](./examples)

## Goals of the project

this programm should be able to decode, encode and create a new D2S files

## Legal Notes

Diablo 2 and its content is ©2000 Blizzard Entertainment, Inc. All rights reserved.
Diablo and Blizzard Entertainment are trademarks or registered trademarks of Blizzard Entertainment,
Inc. in the U.S. and/or other countries.

This package is ABSOLUTLY NO WARRITY!
The code developers aren't responsible
for any damage caused by this software or illegal uses of it!

## Custom D2S Files

Sometimes, you may want to parse some custom d2s files. For example
these created by certain d2 mods. They are not supported by default,
however you can do some configuration to support them.

**NOTE:** add this code in your custom implementation of package or
modify cmd/editor

**NOTE:** in case of any problems/errors please fill an issue in this repo!

- [MedianXL](https://www.median-xl.com/)
```golang
        myD2S := d2s.New()
        m := map[d2sstats.StatID]int{
                d2sstats.Strength:       11,
                d2sstats.Energy:         11,
                d2sstats.Dexterity:      11,
                d2sstats.Vitality:       11,
                d2sstats.UnusedStats:    11,
                d2sstats.UnusedSkills:   8,
                d2sstats.CurrentHP:      21,
                d2sstats.MaxHP:          21,
                d2sstats.CurrentMana:    21,
                d2sstats.MaxMana:        21,
                d2sstats.CurrentStamina: 21,
                d2sstats.MaxStamina:     21,
                d2sstats.Level:          8,
                d2sstats.Experience:     32,
                d2sstats.Gold:           22,
                d2sstats.StashedGold:    25,
                88:                      48,
        }

        myD2S.Stats.UserStatMap(m)
        myD2S.Skills.SetSkillsCount(95)
        myD2S.Items.IgnoreErrors()
```
