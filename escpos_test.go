package escpos

import (
	"net"
	"testing"
)

func TestPrint(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.1.100:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &Config{CutOffset: 4})

	// Normal Text
	p.Write([]byte("Normal Text\n"))

	p.LineFeed(2)

	// 1-Dot Underlines Text
	p.Underline(1)
	p.Write([]byte("Underlined Text\n"))
	p.Underline(0)

	p.LineFeed(2)

	// 2-Dot Underlines Text
	p.Underline(1)
	p.Write([]byte("Underlined Text\n"))
	p.Underline(0)

	p.LineFeed(2)

	// Bold Text
	p.Bold(1)
	p.Write([]byte("Bold Text\n"))
	p.Bold(0)

	p.LineFeed(2)

	// Double Sized Text
	p.Size(1, 1)
	p.Write([]byte("Double Sized Text\n"))
	p.Size(0, 0)

	p.PrintAndCut()

}
