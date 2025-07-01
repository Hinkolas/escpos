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
	LineWidth int
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

func (p *Printer) Justify(mode byte) (int, error) {
	return p.w.Write([]byte{ESC, 'a', byte(mode)})
}

func (p *Printer) Size(witdh, height byte) (int, error) {
	mode := ((witdh & 0x0F) << 4) | (height & 0x0F)
	return p.w.Write([]byte{GS, '!', byte(mode)})
}

func (p *Printer) WriteImage(img image.Image) (int, error) {

	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	draw.Draw(grayImg, bounds, img, bounds.Min, draw.Src)
	width, height := bounds.Dx(), bounds.Dy()

	if width > p.Config.LineWidth {
		return 0, fmt.Errorf("image width exceeds maximum dots/line")
	}

	// 2. Convert the image to raster data
	// Each byte in raster data represents 8 horizontal pixels.
	bWidth := (width + 7) / 8
	rData := make([]byte, bWidth*height)

	for y := range height {
		for x := range width {
			// If the pixel is darker than the threshold, set the corresponding bit
			if grayImg.GrayAt(x+bounds.Min.X, y+bounds.Min.Y).Y < 128 {
				byteIndex := (y * bWidth) + (x / 8)
				bitIndex := 7 - uint(x%8)
				rData[byteIndex] |= 1 << bitIndex
			}
		}
	}

	// 3. Construct and send the command
	// GS v 0 m xL xH yL yH d1...dk
	// m = 0 (normal mode)
	// xL, xH = width in bytes
	// yL, yH = height in pixels
	xL := byte(bWidth % 256)
	xH := byte(bWidth / 256)
	yL := byte(height % 256)
	yH := byte(height / 256)

	header := []byte{GS, 'v', '0', 0, xL, xH, yL, yH}
	command := append(header, rData...)

	return p.Write(command)

}

func (p *Printer) WriteQRCode(data []byte) (int, error) {
	// TODO: Implement method for writing qr-codes to the printer buffer
	return 0, nil
}

func (p *Printer) WriteBarcode(data []byte) (int, error) {
	// TODO: Implement method for writing barcodes to the printer buffer
	return 0, nil
}

func (p *Printer) WriteMarkdown(data []byte) (int, error) {
	// TODO: Implement method for writing markdown to the printer buffer
	return 0, nil
}
