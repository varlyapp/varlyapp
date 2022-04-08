package services

import (
	"encoding/json"
	"fmt"
	"log"
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

	Theme     string     `json:"theme"`
	Documents []Document `json:"documents"`
}

type SettingsService struct {
	Path     string
	Settings *Settings
}

func GetSettingsDefaults() *Settings {
	return &Settings{
		Theme:     "auto",
		Documents: make([]Document, 0),
	}
}

func LoadSettings(path string) (*Settings, error) {
	s := &Settings{}

	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// file not found - use defaults
			return GetSettingsDefaults(), nil
		}
		return nil, errors.Wrap(err, "ReadFile")
	}
	err = json.Unmarshal(b, s)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal")
	}

	return s, nil
}

func NewSettingsService(docsdir string) *SettingsService {
	path := filepath.Join(docsdir, "settings.json")
	settings, err := LoadSettings(path)

	if err != nil {
		log.Fatalf("unable to load settings: %v\n", err)
	}

	return &SettingsService{
		Path:     path,
		Settings: settings,
	}
}

func (s *SettingsService) GetSettings() *Settings{
	return s.Settings
}
func (s *SettingsService) SaveSettings() error {
	s.Settings.Lock()
	defer s.Settings.Unlock()

	b, err := json.MarshalIndent(s.Settings, "", "\t")
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

func (s *SettingsService) SaveDocuments(newdocs []Document) {
	docs := s.Settings.Documents

	for _, newdoc := range newdocs {
		for i, doc := range docs {
			if newdoc.Path == doc.Path {
				fmt.Println("deleting ", doc.Path)
				docs = append(docs[:i], docs[i+1:]...)
			}
		}
		fmt.Println("appending ", newdoc.Path)
		docs = append(docs, newdoc)
	}

	s.Settings.Documents = docs

	s.SaveSettings()
}
