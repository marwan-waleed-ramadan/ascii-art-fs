# ASCII Art FS

Go ASCII art project focused on loading banner files and selecting a rendering style. NextEra Education × 01Edu.

## Run
```sh
go run . "Hello" shadow
go test ./...
```
Available banners: `standard`, `shadow`, `thinkertoy`. Keep the banner text files in the working directory. Requires Go compatible with `go.mod`.

## Implementation
Banner parsing, rune-to-glyph mapping, multiline rendering, command-line arguments and Go tests.

## Provenance
Source snapshot from my 01Edu Gitea repository. Original history remains on Gitea.
