package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s field1:header field2 ... < somefile.json\n", os.Args[0])
		os.Exit(1)
	}

	// We assume that the input JSON is a single document containing an array of objects.
	var data []map[string]interface{}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("can't read input: %v", err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalf("can't unmarshal JSON: %v", err)
	}

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	err = writer.Write(headersOf(os.Args[1:]))
	if err != nil {
		log.Fatalf("can't write headers: %v", err)
	}

	keys := keysOf(os.Args[1:])
	for _, obj := range data {
		record := []string{}

		for _, name := range keys {
			record = append(record, fmt.Sprintf("%v", obj[name]))
		}

		err := writer.Write(record)
		if err != nil {
			log.Fatalf("can't write record: %v", err)
		}
	}

}

func keysOf(args []string) []string {

	keys := []string{}

	for _, h := range args {
		sep := strings.Split(h, ":")
		keys = append(keys, sep[0])
	}

	return keys
}

func headersOf(args []string) []string {

	heads := []string{}

	for _, h := range args {
		sep := strings.Split(h, ":")
		if len(sep) > 1 {
			heads = append(heads, sep[1])
		} else {
			heads = append(heads, sep[0])
		}
	}

	return heads
}
