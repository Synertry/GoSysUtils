package UI

import (
	"fmt"
	"time"
)

// ProgressBar holds the ui values
// Initial source: https://www.pixelstech.net/article/1596946473-A-simple-example-on-implementing-progress-bar-in-GoLang
type ProgressBar struct {
	timer   time.Time // holds elapsed time after Start
	percent int64     // progress percentage
	cur     int64     // current progress
	total   int64     // total value for progress
	label   string    // prefix of progress bar
	rate    string    // the actual progress bar to be printed
	fill    string    // the fill value for progress bar
	head    string    // the head value for progress bar
}

// NewProgressBar sets the graph values for the progress bar and returns a pointer to ProgressBar
func NewProgressBar(label, fill, head string) *ProgressBar {
	return &ProgressBar{
		label: label,
		fill:  fill,
		head:  head,
	}
}

// Set adjusts the graph values
func (bar *ProgressBar) Set(label, fill, head string) {
	bar.label = label
	bar.fill = fill
	bar.head = head
}

// Start is the initializer
func (bar *ProgressBar) Start(label string, total int64) {
	bar.total = total
	if label != "" {
		bar.label = label + " "
	}
	if bar.fill == "" {
		bar.fill = "="
	}
	if bar.head == "" {
		bar.head = ">"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.fill // initial progress position
	}
	bar.timer = time.Now()
}

// getPercent helper function to calculate the percentage
func (bar *ProgressBar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

// Prog updates the ProgressBar.
// This is the core method of the ProgressBar
func (bar *ProgressBar) Prog(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.fill
	}
	fmt.Printf("\r%s[%-50s]%3d%% %8d/%d elapsed %v\r", bar.label, bar.rate+bar.head, bar.percent, bar.cur, bar.total, time.Since(bar.timer))
}

// Finish resets ProgressBar and prints a new line
func (bar *ProgressBar) Finish() {
	*bar = ProgressBar{}
	fmt.Println()
}

// Reset nils the ProgressBar
func (bar *ProgressBar) Reset() {
	*bar = ProgressBar{}
}
