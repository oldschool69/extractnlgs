package main

import (
	"extractnlgs/nlgparser"
	"flag"
	"os"
)

func main() {

	base := flag.String("base", "", "base file with templates (required)")
	file := flag.String("file", "", "file with missing templates (required)")
	out := flag.String("out", "", "output file (optional)")

	flag.Parse()

	if *base == "" || *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *out == "" {
		*out = "out.txt"
	}

	nlgparser.Extract(*base, *file, *out)
}
