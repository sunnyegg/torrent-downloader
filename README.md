# Torrent Downloader

A simple yet powerful command-line torrent downloader built with Go using the anacrolix's [torrent package](https://github.com/anacrolix/torrent). Download multiple torrents simultaneously with real-time progress tracking.

## Features

- Download multiple torrents simultaneously
- Support for both magnet links and torrent files
- Real-time progress tracking with visual progress bars
- Download speed monitoring
- Peer count display
- Skip already downloaded files
- Resume partial downloads
- Batch download support via text file input
- Clean and informative console interface

## Project Structure

```
.
├── cmd/
│   └── torrent-downloader/     # Main application entry point
│       └── main.go
├── internal/                   # Private application code
│   ├── config/                # Configuration
│   ├── downloader/            # Core downloader logic
│   ├── ui/                    # UI related code
│   └── utils/                 # Utility functions
├── pkg/                       # Public packages (if any)
├── go.mod
└── README.md
```

## Installation

1. Make sure you have Go 1.21 or later installed on your system
2. Clone this repository:
   ```bash
   git clone https://github.com/sunnyegg/torrent-downloader.git
   cd torrent-downloader
   ```
3. Build the application:
   ```bash
   go build -o torrent-downloader ./cmd/torrent-downloader
   ```

## Usage

There are two ways to use the downloader:

### 1. Direct Input Mode

Download one or more torrents directly from command line:

```bash
# Single torrent (magnet link or torrent file)
./torrent-downloader "magnet:?xt=urn:btih:..."
./torrent-downloader path/to/your/file.torrent

# Multiple torrents
./torrent-downloader "magnet:?xt=urn:btih:..." "magnet:?xt=urn:btih:..." path/to/file.torrent
```

### 2. File Input Mode

Download multiple torrents listed in a text file:

```bash
./torrent-downloader -f links.txt
```

The text file should contain one magnet link or torrent file path per line. Empty lines and lines starting with # are ignored.

## Output Directory

All downloads are saved in the `./downloads` directory relative to where you run the program.

## Progress Information

The program provides real-time information for each download:

- File name and size
- Visual progress bar
- Download progress percentage
- Current download speed (MB/s)
- Number of connected peers
- Total active downloads

The display automatically updates every 500ms and shows the status of all downloads simultaneously.

## Features in Detail

### Skip Completed Downloads

If a file is already fully downloaded, the program will verify it and skip the download, showing "Already completed" status.

### Resume Partial Downloads

If a file is partially downloaded, the program will automatically resume from where it left off.

### Batch Processing

Using the file input mode (-f flag), you can queue multiple downloads from a text file, making it easy to batch process many torrents.

## Development

The project follows standard Go project layout and best practices. Key components:

- `cmd/torrent-downloader/main.go`: Application entry point
- `internal/config`: Configuration management
- `internal/downloader`: Core download logic
- `internal/ui`: Terminal UI and display
- `internal/utils`: Shared utilities

To contribute:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request
