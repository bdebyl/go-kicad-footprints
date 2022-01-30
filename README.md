# go-kicad-footprints

A super simple, lazily written, KiCAD config generator for footprint libraries.
This was made because KiCAD does not allow you to select multiple directories,
and selecting 20+ footprint library directories one after the other is tedious.

## Usage
Pass in the `-path` of where you store your `*.pretty` KiCAD Footprint
libraries (e.g. Local git clone of `kicad-footprints` from official)

```
  -path string
          Path to look for .pretty footprints
```
