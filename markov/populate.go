package markov

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

type runeClass int

const (
	runeClassSpace runeClass = iota
	runeClassBoundary
	runeClassWord
)

// Repurposed from bufio.isSpace
func classifyRune(r rune) runeClass {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case '\n', '.':
			return runeClassBoundary
		case ' ', '\t', '\v', '\f', '\r':
			return runeClassSpace
		case '\u0085', '\u00A0':
			return runeClassSpace
		}
		return runeClassWord
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return runeClassSpace
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return runeClassSpace
	}
	return runeClassWord
}

// A scanner compliant with bufio.SplitFunc
// Returns words like bufio.ScanWords, but with special treatment for newlines/periods which return Boundary
func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {

	//fmt.Printf("ScanWords given [%v] [%v]\n", string(data), atEOF)
	seenBoundary := false
	here := 0
	wordStart := -1
	for width := 0; here < len(data); here += width {

		var r rune
		r, width = utf8.DecodeRune(data[here:])
		class := classifyRune(r)
		//fmt.Printf("%v: %v (%v)\n", here, string(r), class)

		switch {
		case class == runeClassWord:
			// In a word
			if seenBoundary {
				//fmt.Printf("Returning previous Boundary\n")
				return here, []byte(Boundary), nil
			}
			if wordStart == -1 {
				// Word starts here
				wordStart = here
			}
		case wordStart > -1:
			// Not in a word anymore, but we know where the last one started
			//fmt.Printf("Returning [%v], skip %v\n", string(data[wordStart:here]), here)
			return here, data[wordStart:here], nil
		case class == runeClassBoundary:
			// At a boundary - cache for debouncing
			seenBoundary = true
		}
	}

	if wordStart > -1 && atEOF {
		//fmt.Printf("Returning last [%v], skip %v\n", string(data[wordStart:here]), here)
		return here, data[wordStart:here], nil
	} else if wordStart > -1 {
		//fmt.Printf("Returning skip %v maintain wordStart\n", wordStart)
		return wordStart, nil, nil
	} else {
		//fmt.Printf("Returning skip %v\n", here)
		return here, nil, nil
	}

}

// Populate parses the text in the supplied reader and populates the Markov matrix
func (m *Matrix) Populate(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(scanWords)
	for scanner.Scan() {
		//fmt.Printf("Adding [%v]\n", scanner.Text())
		m.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	m.Add(Boundary)
}
