package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type words struct {
	Syn []string `json:"syn"`
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get(fmt.Sprintf("https://words.bighugelabs.com/api/2/%s/%s/json", b.APIKey, term))
	if err != nil {
		return syns, errors.New("bighuge: failed when looking for synonyms")
	}
	var data synonyms
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}
