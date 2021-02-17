package main

import (
	"generator-helper/resolver"
	"log"
	"os"
)

func main() {
	args := os.Args
	//cut exe name
	args = args[0:]

	if len(args) < 3 {
		return
	}

	r := resolver.New(args)

	err := r.Resolve()
	if err != nil {
		log.Fatal(err)
	}

	err = r.Do()
	if err != nil {
		log.Fatal(err)
	}
}
