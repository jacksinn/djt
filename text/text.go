package text

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
)

type Noun struct {
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