package main

import (
	"flag"
	"encoding/json"
	"os"
	"log"
)

func main() {

	flag.Parse()

	uris := flag.Args()

	for _, path := range uris {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open '%s', %v", path, err)
		}

		defer fh.Close()

		var tmp interface{}

		dec := json.NewDecoder(fh)
		err = dec.Decode(&tmp)

		if err != nil {
			log.Fatalf("Failed to decode '%s', %v", path, err)
		}

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(tmp)

		if err != nil {
			log.Fatalf("Failed to encode '%s', %v", path, err)
		}
	}
}
