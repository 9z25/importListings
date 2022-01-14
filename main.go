package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var i int

	f, err := os.Open("product_data.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)

		for value := range record {
			if i == 1 {
				fmt.Println(i)
				fmt.Printf("%s\n", record[value])
			}
		}
		i++
	}
}
