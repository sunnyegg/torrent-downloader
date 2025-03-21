package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/sunnyegg/torrent-downloader/internal/downloader"
	"github.com/sunnyegg/torrent-downloader/internal/utils"
)

// Display manages the UI display
type Display struct {
	statuses []*downloader.Status
}

// New creates a new Display instance
func New(statuses []*downloader.Status) *Display {
	return &Display{
		statuses: statuses,
	}
}

// Start begins the display loop
func (d *Display) Start() {
	var lastUpdate time.Time
	for {
		if time.Since(lastUpdate) < 500*time.Millisecond {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		lastUpdate = time.Now()

		d.clear()
		d.printHeader()
		d.printStatuses()
		d.printSummary()
	}
}

// clear clears the screen and moves cursor to top
func (d *Display) clear() {
	fmt.Print("\033[2J\033[H")
}

// printHeader prints the header section
func (d *Display) printHeader() {
	fmt.Println("Torrent Download Status")
	fmt.Println("======================")
}

// printStatuses prints the status of each download
func (d *Display) printStatuses() {
	for _, status := range d.statuses {
		if status.IsError() {
			fmt.Printf("\n%s\nError: %s\n", status.Name, status.Error)
			continue
		}

		if status.Total > 0 {
			if !status.IsComplete {
				d.printActiveDownload(status)
			} else {
				d.printCompletedDownload(status)
			}
		} else if status.Name != "" {
			fmt.Printf("\n%s: Waiting for metadata...\n", status.Name)
		}
	}
}

// printActiveDownload prints information for an active download
func (d *Display) printActiveDownload(status *downloader.Status) {
	width := 40
	progress := status.Progress()
	completed := int(float64(width) * float64(status.Completed) / float64(status.Total))
	bar := strings.Repeat("=", completed) + strings.Repeat("-", width-completed)

	fmt.Printf("\n%s\n", status.Name)
	fmt.Printf("[%s] %.1f%%\n", bar, progress)
	fmt.Printf("Speed: %.1f MB/s | Peers: %d\n", status.Speed/1024, status.Peers)
	fmt.Printf("Downloaded: %s / %s\n",
		utils.FormatBytes(status.Completed),
		utils.FormatBytes(status.Total))
}

// printCompletedDownload prints information for a completed download
func (d *Display) printCompletedDownload(status *downloader.Status) {
	if status.Skipped {
		fmt.Printf("\n%s: Already completed (%s)\n",
			status.Name,
			utils.FormatBytes(status.Total))
	} else {
		fmt.Printf("\n%s: Download completed (%s)\n",
			status.Name,
			utils.FormatBytes(status.Total))
	}
}

// printSummary prints the summary section
func (d *Display) printSummary() {
	fmt.Println("\n======================")

	allComplete := true
	activeDownloads := 0

	for _, status := range d.statuses {
		if !status.IsComplete && !status.IsError() {
			allComplete = false
			activeDownloads++
		}
	}

	if allComplete {
		fmt.Println("All downloads completed!")
	} else {
		fmt.Printf("Active downloads: %d\n", activeDownloads)
	}
}
