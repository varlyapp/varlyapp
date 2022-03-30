package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
)

type Document struct {
	Title   string `json:"title"`
	Path    string `json:"path"`
	Preview string `json:"preview"`
}

type Settings struct {
	sync.Mutex

	Path      string      `json:"path"`
	Theme     string      `json:"theme"`
	Documents []Document `json:"documents"`
}

func GetDefaults(path string) *Settings {
	return &Settings{
		Path: path,
		Theme: "auto",
		Documents: make([]Document, 0),
	}
}

func LoadSettings(path string) (*Settings, error) {
	fmt.Println(path)
	s := &Settings{
		Path: path,
	}
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// file not found - use defaults
			return GetDefaults(path), nil
		}
		return nil, errors.Wrap(err, "ReadFile")
	}
	err = json.Unmarshal(b, s)
	fmt.Println(s)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal")
	}
	return s, nil
}

func (s *Settings) SaveDocuments(newdocs []Document) {
	docs := s.Documents

	for _, newdoc := range newdocs {
		for i, olddoc := range docs {
			if newdoc.Path == olddoc.Path {
				fmt.Println("deleting ", olddoc.Path)
				docs = append(docs[:i], docs[i+1:]...)
			}
		}
		fmt.Println("appending ", newdoc.Path)
		docs = append(docs, newdoc)
	}

	s.Documents = docs

	s.Save()
}

func (s *Settings) Save() error {
	s.Lock()
	defer s.Unlock()
	// s.setDefaults()
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return errors.Wrap(err, "MarshalIndent")
	}
	err = os.MkdirAll(filepath.Dir(s.Path), 0777)
	if err != nil {
		return errors.Wrap(err, "MkdirAll")
	}
	err = os.WriteFile(s.Path, b, 0777)
	if err != nil {
		return errors.Wrap(err, "WriteFile")
	}
	return nil
}
