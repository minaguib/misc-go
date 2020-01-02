package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minaguib/misc-go/markov"
)

func main() {

	m := markov.NewMatrix()

	// Populate from CLI arg files or stdin
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			fh, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			m.Populate(fh)
			fh.Close()
		}
	} else {
		m.Populate(os.Stdin)
	}

	// Generate to stdout
	w := markov.Boundary
	for {
		w = m.Get(w)
		if w == markov.Boundary {
			fmt.Println("")
		} else {
			fmt.Print(w, " ")
		}
	}

}
