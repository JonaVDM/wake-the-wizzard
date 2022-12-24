package storage

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/google/uuid"
)

const FILENAME string = "wol-server/data.txt"

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Mac  string `json:"mac"`
}

// Get returns the whole list that is currently there
func Get() ([]Item, error) {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []Item{}, nil
		}
		return nil, errors.New("could not open file")
	}

	line := strings.TrimSpace(string(data))
	lines := strings.Split(line, "\n")
	items := make([]Item, len(lines))

	for i, l := range lines {
		spl := strings.Split(l, ";")
		if len(spl) != 3 {
			return nil, errors.New("invalid data")
		}
		items[i] = Item{
			Id:   spl[0],
			Name: spl[1],
			Mac:  spl[2],
		}
	}

	return items, nil
}

// Add a entry to the file, returns the newly created uuid
func Add(name, mac string) (string, error) {
	if strings.Contains(name, ";") {
		return "", errors.New("name cannot contain ;")
	}

	f, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", errors.New("could not open file")
	}
	defer f.Close()

	id := uuid.New().String()

	_, err = f.WriteString(fmt.Sprintf("%s;%s;%s\n", id, name, mac))
	return id, err
}

// Delete removes a entry from the storage file
func Delete(id string) error {
	items, err := Get()
	if err != nil {
		return err
	}

	out := ""
	for _, item := range items {
		if item.Id == id {
			continue
		}

		out += fmt.Sprintf("%s;%s;%s\n", item.Id, item.Name, item.Mac)
	}

	return os.WriteFile(FILENAME, []byte(out), 0644)
}
