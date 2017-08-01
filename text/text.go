package text

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
)

type Noun struct {
	Text string
}

type Adjective struct {
	Text string
}

func SaveNouns(filename string, nouns []Noun) error {
	b, err := json.Marshal(nouns)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadNouns(filename string) ([]Noun, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Noun{}, err
	}
	var nouns []Noun
	if err := json.Unmarshal(b, &nouns); err != nil {
		return []Noun{}, err
	}
	return nouns, nil
}

func SaveAdjectives(filename string, adjectives []Adjective) error {
	b, err := json.Marshal(adjectives)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadAdjectives(filename string) ([]Adjective, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Adjective{}, err
	}
	var adjectives []Adjective
	if err := json.Unmarshal(b, &adjectives); err != nil {
		return []Adjective{}, err
	}
	return adjectives, nil
}