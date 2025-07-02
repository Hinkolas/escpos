package escpos

import (
	"image"
	_ "image/png"
	"log"
	"net"
	"os"
	"testing"
)

func TestStyling(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.69.111:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &PrinterConfig{CutOffset: 5, LineWidth: 576})

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

	p.LineFeed(2)

	// Centered Text
	p.Justify(1)
	p.Write([]byte("Centered Text\n"))
	p.Justify(0)

	p.LineFeed(2)

	p.PrintAndCut()

}

func TestImages(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.69.111:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &PrinterConfig{CutOffset: 5, LineWidth: 576})

	// Open the image file
	file, err := os.Open("test/image-256.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	p.LineFeed(4)

	// Open the image file
	file, err = os.Open("test/image-342.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	p.LineFeed(4)

	// Open the image file
	file, err = os.Open("test/image-512.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	p.LineFeed(4)

	// Open the image file
	file, err = os.Open("test/image-576.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	// p.LineFeed(4)
	p.PrintAndCut()

}

func TestLogo(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.69.111:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &PrinterConfig{CutOffset: 5, LineWidth: 576})

	// Open the image file
	file, err := os.Open("test/logo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	p.LineFeed(4)

	p.PrintAndCut()

}

func TestDither(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.69.111:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &PrinterConfig{CutOffset: 5, LineWidth: 576})

	// Open the image file
	file, err := os.Open("test/dithered.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	p.Justify(1)
	_, err = p.WriteImage(img, &ImageConfig{Threshold: 240})
	if err != nil {
		t.Fatal(err)
	}
	p.Justify(0)

	p.LineFeed(4)

	p.PrintAndCut()

}

func TestLongText(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.69.111:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &PrinterConfig{CutOffset: 5, LineWidth: 576})

	// Title: The Whispering Woods
	p.Size(1, 1)
	p.Justify(1)
	p.Bold(1)
	p.Write([]byte("The Whispering Woods\n"))
	p.Bold(0)
	p.Justify(0)
	p.Size(0, 0)
	p.LineFeed(2)

	// Paragraph 1
	p.Write([]byte("The old woman, Elara, adjusted the worn shawl around her shoulders as she surveyed the encroaching mist.\n"))
	p.Write([]byte("It curled like a lazy cat around the gnarled roots of the ancient oak at the edge of her cottage garden, seeping into the very stones of her home.\n"))
	p.Write([]byte("They called it the Whispering Mist, for it was said to carry the murmurs of forgotten things, a chorus of echoes from a time when the woods were wilder and magic more potent.\n"))
	p.LineFeed(2)

	// Paragraph 2
	p.Write([]byte("Elara had lived beside the Whispering Woods for seventy years, since she was a slip of a girl with bright eyes and a heart full of curiosity.\n"))
	p.Write([]byte("Now, her eyes held the wisdom of countless sunsets, and her heart, though still curious, beat to a slower, more deliberate rhythm.\n"))
	p.Write([]byte("Tonight, however, there was a disquiet in the air that even she, with all her years, couldn't quite decipher.\n"))
	p.LineFeed(2)

	// Paragraph 3
	p.Write([]byte("A sharp rap on her cottage door startled her.\n"))
	p.Write([]byte("It was rare to have visitors so late, especially with the mist thickening.\n"))
	p.Write([]byte("She pulled open the heavy wooden door to find a young man standing on her doorstep, his face pale and streaked with dirt, his breath coming in ragged gasps.\n"))
	p.Write([]byte("He clutched a small, leather-bound book to his chest as if it were a life raft.\n"))
	p.LineFeed(2)

	// Paragraph 4
	p.Write([]byte("Inside, by the hearth, the young man introduced himself as Liam.\n"))
	p.Write([]byte("He was a scholar, he explained, on a quest for ancient lore, drawn by whispers of a hidden library deep within the Whispering Woods.\n"))
	p.Write([]byte("Elara offered him a mug of steaming herbal tea, its warmth seeping into his trembling hands.\n"))
	p.LineFeed(2)

	p.PrintAndCut()

}
