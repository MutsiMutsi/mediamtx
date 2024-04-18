// main executable.
package main

import (
	"log"
	"os"

	"github.com/bluenviron/mediamtx/internal/core"
)

func main() {
	s, ok := core.New(os.Args[1:], publishTSPart)
	if !ok {
		os.Exit(1)
	}

	s.Wait()
}

func publishTSPart(segment []byte) {
	log.Println("SHOW TIME", len(segment))
}
