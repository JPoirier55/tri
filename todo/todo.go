package todo

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const filePath = "/Users/jakepoirier/go/src/github.com/jpoirier55/tri/"

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filePath + filename, b, 0644)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filePath + filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}