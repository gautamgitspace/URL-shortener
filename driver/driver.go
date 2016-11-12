package main

import (
	"flag"
	"fmt"
	"urlshortener"
	"os"
	
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Expecting: Long URI as an argumnet")
		return
	}

	uri, err := goisgd.Shorten(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	fmt.Println(uri)
	return
}
