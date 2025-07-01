package escpos

import (
	"net"
	"testing"
)

func TestStyling(t *testing.T) {

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

	p.LineFeed(2)

	// Centered Text
	p.Justify(1)
	p.Write([]byte("Centered Text\n"))
	p.Justify(0)

	p.LineFeed(2)

	p.PrintAndCut()

}

func TestLongText(t *testing.T) {

	socket, err := net.Dial("tcp", "192.168.1.100:9100")
	if err != nil {
		t.Fatal(err)
	}
	defer socket.Close()

	p := NewPrinter(socket, &Config{CutOffset: 4})

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

	// Paragraph 5
	p.Write([]byte("As the tea worked its magic, Liam's eyes fell upon a peculiar carving on Elara's mantelpiece – a coiled serpent entwined around a blooming rose.\n"))
	p.Write([]byte("His eyes widened.\n"))
	p.Write([]byte("\"The Serpent and the Rose!\" he breathed, a thrill in his voice.\n"))
	p.Write([]byte("\"It's said to be the symbol of the Elder Keepers, guardians of the forgotten knowledge!\"\n"))
	p.LineFeed(2)

	// Paragraph 6
	p.Write([]byte("Elara merely smiled, a knowing glint in her eyes.\n"))
	p.Write([]byte("\"Some things are best left forgotten, young one.\"\n"))
	p.LineFeed(2)

	// Paragraph 7
	p.Write([]byte("But Liam, emboldened by the tea and the warmth of the fire, pressed on.\n"))
	p.Write([]byte("He spoke of the legends he'd read, of powerful spells and ancient truths hidden from mankind.\n"))
	p.Write([]byte("He showed her the book he carried, its pages brittle with age, filled with strange symbols and intricate diagrams.\n"))
	p.LineFeed(2)

	// Paragraph 8
	p.Write([]byte("As he spoke, the whispers from the woods grew louder, seeming to press against the very walls of the cottage.\n"))
	p.Write([]byte("Elara felt a chill that had nothing to do with the mist.\n"))
	p.Write([]byte("She remembered the stories her grandmother told, of those who ventured too deep, lured by promises of power, only to be consumed by the woods' embrace.\n"))
	p.LineFeed(2)

	// Paragraph 9
	p.Write([]byte("Suddenly, a gust of wind rattled the windows, and the fire in the hearth flickered, casting dancing shadows across the room.\n"))
	p.Write([]byte("The book in Liam's hand flew open to a page adorned with a particularly elaborate symbol, identical to one he’d seen repeatedly in the carvings around Elara’s cottage.\n"))
	p.Write([]byte("A faint, almost imperceptible hum filled the air.\n"))
	p.LineFeed(2)

	// Paragraph 10
	p.Write([]byte("Liam, oblivious to the change in the atmosphere, traced the symbol with his finger, his eyes alight with a feverish excitement.\n"))
	p.Write([]byte("\"This is it! The key! This will lead me to the library!\"\n"))
	p.LineFeed(2)

	// Paragraph 11
	p.Write([]byte("Elara knew, with a certainty that chilled her to the bone, that this was a trap.\n"))
	p.Write([]byte("The woods were not just a place of forgotten knowledge, but also of forgotten dangers.\n"))
	p.Write([]byte("She looked at the symbol, then at Liam's eager face, and a decision formed in her mind.\n"))
	p.LineFeed(2)

	// Paragraph 12
	p.Write([]byte("\"Liam,\" she said, her voice calm but firm, \"there are some secrets that should remain buried.\"\n"))
	p.Write([]byte("\"The woods protect themselves.\"\n"))
	p.LineFeed(2)

	// Paragraph 13
	p.Write([]byte("But he wasn't listening.\n"))
	p.Write([]byte("His eyes were fixed on the ancient text, his fingers trembling as he prepared to read aloud.\n"))
	p.Write([]byte("Elara moved swiftly, her movements surprisingly agile for her age.\n"))
	p.Write([]byte("She snatched the book from his grasp.\n"))
	p.LineFeed(2)

	// Paragraph 14
	p.Write([]byte("\"No!\" Liam cried, bewildered.\n"))
	p.Write([]byte("\"What are you doing?\"\n"))
	p.LineFeed(2)

	// Paragraph 15
	p.Write([]byte("Elara didn't answer.\n"))
	p.Write([]byte("With a strength born of desperation and ancient wisdom, she threw the book into the roaring flames of the hearth.\n"))
	p.Write([]byte("The brittle pages caught instantly, curling and blackening, the symbols dissolving into ash.\n"))
	p.LineFeed(2)

	// Paragraph 16
	p.Write([]byte("Liam stared, aghast, then turned a furious face to Elara.\n"))
	p.Write([]byte("\"Why? Why would you do that? You've destroyed centuries of knowledge!\"\n"))
	p.LineFeed(2)

	// Paragraph 17
	p.Write([]byte("Elara watched the last embers of the book fade.\n"))
	p.Write([]byte("\"Some knowledge, Liam,\" she said softly, \"is a burden too heavy to bear.\"\n"))
	p.Write([]byte("\"The Whispering Woods have whispered to me for decades.\"\n"))
	p.Write([]byte("\"They speak of balance, of respect, and of the profound dangers of disturbing what lies truly dormant.\"\n"))
	p.LineFeed(2)

	// Paragraph 18
	p.Write([]byte("The mist outside began to recede, slowly at first, then with more purpose.\n"))
	p.Write([]byte("The whispers quieted.\n"))
	p.Write([]byte("The cottage, once filled with an unsettling hum, now felt simply warm and safe.\n"))
	p.LineFeed(2)

	// Paragraph 19
	p.Write([]byte("Liam, still angry, but with a dawning comprehension in his eyes, looked from the fading mist to Elara's serene face.\n"))
	p.Write([]byte("He had come seeking power, but perhaps, he realized, he had found something far more valuable: a quiet wisdom, and a reminder that some mysteries are best left to the whispers of the woods.\n"))
	p.Write([]byte("He looked out into the clearing night, then back at Elara, and for the first time, truly saw not just an old woman, but a guardian of true, profound, and often silent knowledge.\n"))
	p.Write([]byte("He still had much to learn.\n"))
	p.LineFeed(2)

	p.PrintAndCut()

}
