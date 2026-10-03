# PicoConv

**PicoConv** is a lightweight command-line converter written in Go. The compiled binary is `pconv`.

It converts video, audio, images, and configuration files through one CLI. Media conversion uses [FFmpeg](https://ffmpeg.org/); image and configuration conversion use Go libraries.

## Supported formats

- **Video:** MP4, MKV, MOV, AVI, WebM, WMV, FLV
- **Audio:** MP3, FLAC, WAV, OGG, AAC, Opus
- **Images:** JPEG, PNG, WebP, AVIF, BMP, TIFF, HEIC, SVG (depending on decoder/encoder support)
- **Configuration:** INI, JSON, TOML, YAML

PicoConv can convert between distinct formats within each category. Video and audio conversion is delegated to FFmpeg.

## Requirements

- Go **1.27.1+** to build PicoConv
- FFmpeg available in `PATH` for audio and video conversion

```bash
ffmpeg -version
```

## Build

Build the optimized executable:

```bash
go build -ldflags '-s -w' -o pconv ./cmd/converter
```

Install it into the Go binary directory instead:

```bash
go install -ldflags '-s -w' ./cmd/converter
```

Verify the result:

```bash
./pconv version
./pconv --help
```

## Usage

### Convert

```bash
pconv convert --to <format> <input>
```

The source format is detected automatically by default. Use `--from` to specify it explicitly. Without `--output`, the output keeps the input name and receives the target extension.

```bash
# Creates photo.webp
pconv convert --to webp photo.png

# Explicit source and destination
pconv convert --from png --to jpeg --output photo.jpg photo.png

# Audio and video
pconv convert --from wav --to flac recording.wav
pconv convert --from mkv --to mp4 movie.mkv

# Configuration files
pconv convert --from yaml --to json config.yaml
```

### Inspect a file

```bash
pconv inspect photo.png
```

This prints the detected file type.

### List conversions

```bash
pconv list                    # all registered conversions
pconv list --flat             # unique format names
pconv list --src png          # filter by source format
pconv list --trg webp         # filter by target format
pconv list --json             # JSON output
pconv list --flat --json      # flat JSON output
```

### Version

```bash
pconv version
```

## Conversion options

These flags are available on `convert`:

| Flag | Default | Description |
|---|---:|---|
| `--from` | `auto` | Input format |
| `--to` | — | Required output format |
| `--output` | Derived | Output path |
| `--quality` | `80` | Output quality |
| `--alpha-quality` | `80` | Alpha-channel quality |
| `--compress-lvl` | `0` | Compression level |
| `--method` | `4` | Encoder method |
| `--speed` | `6` | Encoder speed/preset |
| `--indent` | `0` | Structured output indentation |
| `--comma` | `;` | CSV delimiter option |
| `--comment` | — | Comment option |
| `--lossless` | `false` | Lossless output where supported |

Options may be ignored when they are not applicable to the selected converter.

## Development

```bash
go test ./...
go build -o pconv ./cmd/converter
```

## License

PicoConv is free software released under the **GNU General Public License v3.0**.

Copyright (C) 2026 Gorozhankin Andrey

See the full license text in [`LICENSE`](LICENSE).
