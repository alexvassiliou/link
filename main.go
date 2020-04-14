package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileName := flag.String("file", "ex1.html", "name of the html file to be parsed")
	flag.Parse()

	f, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(f))

}
