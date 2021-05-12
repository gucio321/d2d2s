## Description

This package can be used for loading, modifing and saving Diablo II
character save files (.D2S).

## Dependencies

### D2S structure documentation:

- [krisives/d2s-format](https://github.com/krisives/d2s-format)
- [Diablo II Saved Game File Format](https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml)
- [nokka/d2s](https://github.com/nokka/d2s) [(see details)](./docs/base.md)

### data reader

this project uses [Stream readers based on these used in OpenDiablo2 project](https://github.com/OpenDiablo2/OpenDiablo2)
for more informations, see [here](./docs/base.md)

## Status

for now (May 2021) the project is early in development.
However, it is possible to:
- decode character save file into a (in a greater part) human-readable structure
- encoding part of structure (excepting items list)

## Goals of the project

this programm should be able to decode and encode D2S files
