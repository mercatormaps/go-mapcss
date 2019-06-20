# go-mapcss #

## Features ##

### `canvas` ###

| Properties     | Supported              |
| -------------- |:----------------------:|
| `antialiasing` | `full`, `text`, `none` |
| `fill-color`   |                        |
| `fill-opacity` |                        |
| `fill-image`   |                        |

## Development ##

The included ANTLR grammar file ([`MapCSS.g4`](MapCSS.g4)) is used to produce the [`parser`](internal) package. The Go target was added to ANTLR v4.6, so the command below downloads the latest version at the time of writing:

```bash
curl -O https://www.antlr.org/download/antlr-4.7.2-complete.jar
```

This command generates the Go target package:

```bash
java -jar /path/to/antlr-4.7.2-complete.jar -Dlanguage=Go -o internal/ MapCSS.g4
```