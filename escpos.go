package escpos

import (
	"bufio"
	"bytes"
	"image"
	"io"
)

type Printer struct {
	Config *Config

	w *bufio.Writer
}

type Config struct {
	CutOffset int
}

// NewPrinter creates a new printer with the given writer and config
func NewPrinter(w io.Writer, config *Config) *Printer {
	if config == nil {
		config = &Config{}
	}
	return &Printer{
		w:      bufio.NewWriter(w),
		Config: config,
	}
}

// Write writes data to the printer buffer
func (p *Printer) Write(data []byte) (int, error) {
	if len(data) > 0 {
		return p.w.Write(data)
	}
	return 0, nil
}

// Print flushes the buffer to output
func (p *Printer) Print() error {
	return p.w.Flush()
}

// Cut sends cut command with offset
func (p *Printer) Cut() (int, error) {
	// send the line-feeds to mitigate the cut offset
	n1, err := p.LineFeed(p.Config.CutOffset)
	if err != nil {
		return n1, err
	}
	// send the cut command
	n2, err := p.w.Write([]byte{GS, 'V', 1})
	return n1 + n2, err
}

// PrintAndCut cuts the paper and flushes output
func (p *Printer) PrintAndCut() error {
	if _, err := p.Cut(); err != nil {
		return err
	}
	return p.Print()
}

// LineFeed sends n line feed characters
func (p *Printer) LineFeed(n int) (int, error) {
	return p.w.Write(bytes.Repeat([]byte{'\n'}, n))
}

// MARK: Experimental

func (p *Printer) Underline(mode byte) (int, error) {
	return p.w.Write([]byte{ESC, '-', byte(mode)})
}

func (p *Printer) Bold(mode byte) (int, error) {
	return p.w.Write([]byte{ESC, 'E', byte(mode)})
}

func (p *Printer) Size(witdh, height byte) (int, error) {
	mode := ((witdh & 0x0F) << 4) | (height & 0x0F)
	return p.w.Write([]byte{GS, '!', byte(mode)})
}

func (p *Printer) WriteImage(img image.Image) (int, error) {
	// TODO: Implement method for writing images to the printer buffer
	return 0, nil
}
