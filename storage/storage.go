package storage

import (
	"errors"
	"os"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const FILENAME string = "wol-server/data.yml"
const FOLDER string = "wol-server"

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Mac  string `json:"mac"`
}

// Get returns the whole list that is currently there
func Get() ([]Item, error) {
	file, err := os.ReadFile(FILENAME)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return []Item{}, nil
	} else if err != nil {
		return nil, err
	}

	var item []Item = make([]Item, 0)
	if err := yaml.Unmarshal(file, &item); err != nil {
		return nil, err
	}

	return item, nil
}

// Add a entry to the file, returns the newly created uuid
func Add(name, mac string) (string, error) {
	items, err := Get()
	if err != nil {
		return "", err
	}

	item := Item{
		Id:   uuid.New().String(),
		Name: name,
		Mac:  mac,
	}

	if err := save(append(items, item)); err != nil {
		return "", err
	}

	return item.Id, nil
}

// Delete removes a entry from the storage file
func Delete(id string) error {
	items, err := Get()
	if err != nil {
		return err
	}

	newList := make([]Item, 0)
	for _, item := range items {
		if item.Id == id {
			continue
		}

		newList = append(newList, item)
	}

	return save(newList)
}

// save all the data, by writing it to the disk
func save(list []Item) error {
	body, err := yaml.Marshal(list)
	if err != nil {
		return err
	}

	if _, err := os.Stat(FOLDER); err != nil && errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(FOLDER, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return os.WriteFile(FILENAME, body, 0644)
}
