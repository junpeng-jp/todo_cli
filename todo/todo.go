package todo

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Item struct {
	Description string   `yaml:"description"`
	Priority    int      `yaml:"priority"`
	Project     string   `yaml:"project"`
	Tags        []string `yaml:"tags"`
	DueDate     int64    `yaml:"due_date"`
	Done        bool     `yaml:"done"`
	CreatedAt   int64    `yaml:"created_at"`
	UpdatedAt   int64    `yaml:"updated_at"`
}

func NewItem(id int, description string, priority int, project string, tags []string, dueDate time.Time) *Item {
	if project == "" {
		project = "unclassified"
	}

	item := Item{
		Description: description,
		Priority:    priority,
		Project:     project,
		Tags:        tags,
		DueDate:     dueDate.Unix()*int64(time.Second/time.Millisecond) + int64(id),
		Done:        false,
		CreatedAt:   time.Now().UnixMilli(),
		UpdatedAt:   time.Now().UnixMilli(),
	}

	return &item
}

func SaveItems(filename string, items []Item) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	cobra.CheckErr(err)

	encoder := yaml.NewEncoder(f)
	defer func() {
		err := f.Close()

		if err != nil {
			panic(err)
		}
	}()

	for _, item := range items {
		// Marshal each item in array as separate document
		encoder.Encode(item)
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	decoder := yaml.NewDecoder(bytes.NewReader(b))

	for {

		item := Item{}

		if err := decoder.Decode(&item); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		items = append(items, item)
	}

	return items, nil
}
