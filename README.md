# go-mapcss #

`go-mapcss` is a Go parser for MapCSS.

## Features ##

### Rules ###

| Selectors           | Supported                                                                |
| ------------------- | ------------------------------------------------------------------------ |
| multiple selectors  | :heavy_check_mark:                                                       |
| multiple attributes | :heavy_check_mark:                                                       |
| negated attributes  | :heavy_check_mark:                                                       |
| zoom level (`|z`)   | :heavy_check_mark: range (`|z1-9`), exact (`|z1`), minimum (`|z1-`)      |

### `canvas` rule properties ###

| Properties     | Supported                                                   |
| -------------- | ----------------------------------------------------------- |
| `antialiasing` | :heavy_check_mark: `full`, `text`, `none`                   |
| `fill-color`   | :heavy_check_mark: `rgb`, `rgba`, hex (3, 4, 6 and 8-digit) |
| `fill-opacity` | :heavy_check_mark: `0` (transparent) — `1` (opaque)         |
| `fill-image`   |                                                             |

## Development ##

The included ANTLR grammar file ([`MapCSS.g4`](MapCSS.g4)) is used to produce the [`parser`](internal) package. The Go target was added to ANTLR v4.6, so the command below downloads the latest version at the time of writing:

```bash
curl -O https://www.antlr.org/download/antlr-4.7.2-complete.jar
```

This command generates the Go target package:

```bash
java -jar /path/to/antlr-4.7.2-complete.jar -Dlanguage=Go -o internal/ MapCSS.g4
```