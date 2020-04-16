package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("file", "ex1.html", "name of the html file to be parsed")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	l, err := ParseLinks(f)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range l {
		fmt.Println(l)
	}

}
