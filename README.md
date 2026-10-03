# PicoConv

**PicoConv** is a small, fast command-line converter written in Go. Its compiled binary is called `pconv`.

It provides one consistent interface for converting:

- **Video** — MP4, MKV, MOV, AVI, WebM, WMV, and FLV
- **Audio** — MP3, FLAC, WAV, OGG, AAC, and Opus
- **Images** — JPEG, PNG, WebP, AVIF, BMP, TIFF, HEIC, and SVG as supported by the selected decoder/encoder
- **Configuration files** — INI, JSON, TOML, and YAML

Media conversion is powered by FFmpeg, while image and configuration conversion use native Go implementations and format-specific Go libraries.

## Features

- One command for media, image, and configuration conversion
- Automatic input type detection when possible
- Explicit source and target format selection
- Automatic output filename generation
- Format discovery and filtering through `list`
- JSON output for scripting and automation
- File type inspection through `inspect`
- Image quality, compression, alpha, lossless, and format-specific options
- Minimal standalone CLI with no runtime service required

## Requirements

- Go **1.27.1** or newer to build PicoConv
- [FFmpeg](https://ffmpeg.org/) available in `PATH` for video and audio conversion

Check the FFmpeg installation with:

```bash
ffmpeg -version
```

## Installation

### Build from source

Clone or download the project, then run:

```bash
go build -ldflags '-s -w' -o pconv ./cmd/converter
```

The optimized executable will be created as `./pconv`.

Optionally install it into your Go binary directory:

```bash
go install -ldflags '-s -w' ./cmd/converter
```

### Verify the build

```bash
./pconv version
./pconv --help
```

## Usage

### Convert a file

```bash
pconv convert --to <format> <input>
```

If `--output` is omitted, PicoConv keeps the input filename and replaces its extension:

```bash
pconv convert --to webp photo.png
# Creates: photo.webp
```

Choose an explicit output path with `--output`:

```bash
pconv convert --to mp3 --output soundtrack.mp3 video.mp4
```

The source format can be detected automatically. To specify it explicitly, use `--from`:

```bash
pconv convert --from png --to jpeg --output photo.jpg photo.png
```

> **Note:** `--from` defaults to `auto`. Automatic detection uses the file contents rather than relying only on the filename extension. For formats that cannot be recognized automatically, provide `--from` explicitly.

### Inspect a file type

```bash
pconv inspect photo.png
```

This prints the detected media type, for example:

```text
png
```

### List available conversions

Show the registered source-to-target conversions:

```bash
pconv list
```

Show only the unique format names:

```bash
pconv list --flat
```

Filter by source or target format:

```bash
pconv list --src png
pconv list --trg webp
```

Return machine-readable JSON:

```bash
pconv list --json
pconv list --flat --json
```

### Show the version

```bash
pconv version
```

## Conversion options

The following flags are available on `convert`:

| Flag | Default | Purpose |
|---|---:|---|
| `--from` | `auto` | Input format; automatically detected by default |
| `--to` | — | Required output format |
| `--output` | Derived from input | Output file path |
| `--quality` | `80` | General output quality |
| `--alpha-quality` | `80` | Alpha-channel quality |
| `--compress-lvl` | `0` | Compression level |
| `--method` | `4` | Encoder method |
| `--speed` | `6` | Encoder speed/preset |
| `--indent` | `0` | Indentation for supported structured output |
| `--comma` | `;` | CSV delimiter option for supported configuration handling |
| `--comment` | — | Comment option for supported configuration handling |
| `--lossless` | `false` | Enable lossless output where supported |

Not every option applies to every conversion. Unsupported options are ignored by converters that do not use them.

## Examples

### Images

```bash
# PNG to WebP
pconv convert --from png --to webp image.png

# JPEG to AVIF with lossless output
pconv convert --from jpeg --to avif --lossless image.jpg

# SVG to PNG with an explicit destination
pconv convert --from svg --to png --output rendered.png artwork.svg
```

### Audio

```bash
pconv convert --from wav --to flac recording.wav
pconv convert --from flac --to mp3 --output podcast.mp3 recording.flac
```

### Video

```bash
pconv convert --from mkv --to mp4 movie.mkv
pconv convert --from mov --to webm --output clip.webm clip.mov
```

### Configuration files

```bash
pconv convert --from yaml --to json config.yaml
pconv convert --from json --to toml --output config.toml config.json
pconv convert --from ini --to yaml settings.ini
```

## Supported formats

PicoConv registers conversions between distinct formats in each category.

### Images

| Decodable input | Encodable output |
|---|---|
| JPEG, PNG, WebP, AVIF, BMP, HEIC, TIFF, SVG | JPEG, PNG, WebP, AVIF, BMP, TIFF |

### Video

MP4, MKV, MOV, AVI, WebM, WMV, and FLV are registered as video formats. Conversion is delegated to FFmpeg.

### Audio

MP3, FLAC, WAV, OGG, AAC, and Opus are registered as audio formats. Conversion is delegated to FFmpeg.

### Configuration

INI, JSON, TOML, and YAML can be converted between one another.

## Development

Run the test suite:

```bash
go test ./...
```

Build a development binary:

```bash
go build -o pconv ./cmd/converter
```

## License

No license file is currently included in this repository.
