package main

import (
	"fmt"
	"github.com/soffes/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
