package downloader

// Status represents the current state of a torrent download
type Status struct {
	Name       string
	Completed  int64
	Total      int64
	LastBytes  int64
	Speed      float64
	Peers      int
	IsComplete bool
	Error      string
	Skipped    bool
}

// NewStatus creates a new download status
func NewStatus() *Status {
	return &Status{}
}

// IsError returns true if the status contains an error
func (s *Status) IsError() bool {
	return s.Error != ""
}

// Progress returns the download progress as a percentage
func (s *Status) Progress() float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Completed) / float64(s.Total) * 100
}
