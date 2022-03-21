package db

import (
	"fmt"
	"time"

	"github.com/varlyapp/varlyapp/backend/fs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Base struct {
	ID uint `gorm:"primaryKey"`
}

type WithDates struct {
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Collection struct {
	Base
	Name            string
	Image           string
	Description     string
	SourceDirectory string
	OutputDirectory string
	WithDates
}

type Edition struct {
	Base
	Name        string
	Image       string
	Description string
	Size        uint
	LayerOrder  []string `gorm:"type:text"`
	WithDates
}

type Layer struct {
	Base
	EditionID uint
	Trait     string
	Name      string
	Path      string
	Weight    float64
	WithDates
}

type Database *gorm.DB

var database Database

func New() *gorm.DB {
	if database == nil {
		name := fmt.Sprintf("%s/varly.db", fs.GetApplicationDocumentsDirectory())
		db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Collection{})
		db.AutoMigrate(&Edition{})
		db.AutoMigrate(&Layer{})

		database = db
	}

	return database
}

func Seed() {
	db := New()
	db.Create(&Collection{
		Name:            "Doodles",
		Image:           "",
		Description:     "A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Community Treasury.",
		SourceDirectory: "/Users/selvinortiz/Desktop/Varly NFT Example",
		OutputDirectory: "/Users/selvinortiz/Desktop/Doodles Output",
	})

	db.Create(&Edition{
		Name:        "Primary Edition",
		Image:       "",
		Description: "10K edition that represents the entire collection",
		Size:        10000,
		LayerOrder:  []string{""},
	})

	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "black#1.png",
		Path:      "black#1.png",
		Weight:    1,
	})

	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "black#1.png",
		Path:      "black#1.png",
		Weight:    1,
	})

	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "blue#29.png",
		Path:      "blue#29.png",
		Weight:    29,
	})
	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "green#20.png",
		Path:      "green#20.png",
		Weight:    20,
	})
	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "magenta#10.png",
		Path:      "magenta#10.png",
		Weight:    10,
	})
	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "purple#10.png",
		Path:      "purple#10.png",
		Weight:    10,
	})
	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "red#20.png",
		Path:      "red#20.png",
		Weight:    20,
	})
	db.Create(&Layer{
		EditionID: 1,
		Trait:     "Backgrounds",
		Name:      "yellow#10.png",
		Path:      "yellow#10.png",
		Weight:    10,
	})
}
