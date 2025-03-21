package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/sunnyegg/torrent-downloader/internal/config"
	"github.com/sunnyegg/torrent-downloader/internal/downloader"
	"github.com/sunnyegg/torrent-downloader/internal/ui"
	"github.com/sunnyegg/torrent-downloader/internal/utils"
)

func main() {
	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  1. Direct input: torrent-downloader <magnet link or torrent file path> [more magnet links or torrent files...]")
		fmt.Println("  2. From file: torrent-downloader -f <links_file.txt>")
		os.Exit(1)
	}

	// Get configuration
	cfg := config.NewDefaultConfig()

	// Create downloads directory if it doesn't exist
	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Create torrent client
	client, err := torrent.NewClient(cfg.Client)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Get input links
	var inputs []string
	if os.Args[1] == "-f" {
		if len(os.Args) != 3 {
			fmt.Println("Usage for file input: torrent-downloader -f <links_file.txt>")
			os.Exit(1)
		}
		var err error
		inputs, err = utils.ReadMagnetLinks(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if len(inputs) == 0 {
			log.Fatal("No valid links found in the file")
		}
	} else {
		inputs = os.Args[1:]
	}

	// Create downloader
	dl := downloader.New(client)

	// Initialize status slice
	statuses := make([]*downloader.Status, len(inputs))
	for i := range statuses {
		statuses[i] = downloader.NewStatus()
	}

	// Start downloads
	for i, input := range inputs {
		dl.Download(input, statuses[i])
	}

	// Start display
	display := ui.New(statuses)
	go display.Start()

	// Wait for downloads to complete
	dl.Wait()
	time.Sleep(time.Second) // Give display time to show final status
}
