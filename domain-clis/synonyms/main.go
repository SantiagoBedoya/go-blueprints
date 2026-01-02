package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/SantiagoBedoya/go-blueprints/domain-clis/thesaurus"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms:", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for the word")
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
