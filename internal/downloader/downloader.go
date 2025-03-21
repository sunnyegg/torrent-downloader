package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/anacrolix/torrent"
)

// Downloader manages torrent downloads
type Downloader struct {
	client *torrent.Client
	wg     sync.WaitGroup
}

// New creates a new Downloader instance
func New(client *torrent.Client) *Downloader {
	return &Downloader{
		client: client,
	}
}

// checkExistingDownload checks if a torrent is already downloaded
func (d *Downloader) checkExistingDownload(t *torrent.Torrent) (bool, bool) {
	<-t.GotInfo()

	downloadPath := filepath.Join("downloads", t.Name())
	info, err := os.Stat(downloadPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false
		}
		return false, false
	}

	if info.Size() == t.Length() {
		t.VerifyData()
		time.Sleep(time.Second)

		if t.BytesCompleted() == t.Length() {
			return true, true
		}
	}

	return true, false
}

// Download starts downloading a torrent from a magnet link or torrent file
func (d *Downloader) Download(input string, status *Status) {
	d.wg.Add(1)
	go d.download(input, status)
}

// download handles the actual download process
func (d *Downloader) download(input string, status *Status) {
	defer d.wg.Done()

	var t *torrent.Torrent
	var err error

	if _, err := os.Stat(input); err == nil {
		t, err = d.client.AddTorrentFromFile(input)
	} else {
		t, err = d.client.AddMagnet(input)
	}

	if err != nil {
		status.Error = fmt.Sprintf("Error adding torrent: %v", err)
		return
	}

	exists, isComplete := d.checkExistingDownload(t)
	status.Name = t.Name()

	if exists {
		if isComplete {
			status.Completed = t.Length()
			status.Total = t.Length()
			status.IsComplete = true
			status.Skipped = true
			return
		}
	}

	t.DownloadAll()

	for {
		stats := t.Stats()
		status.Completed = t.BytesCompleted()
		status.Total = t.Length()

		if status.Completed == status.Total {
			status.IsComplete = true
			return
		}

		currentBytes := int64(stats.BytesReadUsefulData.Int64())
		status.Speed = float64(currentBytes-status.LastBytes) / 1024 // KB/s
		status.LastBytes = currentBytes
		status.Peers = stats.TotalPeers

		time.Sleep(time.Second)
	}
}

// Wait waits for all downloads to complete
func (d *Downloader) Wait() {
	d.wg.Wait()
}
