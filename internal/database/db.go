package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type Client struct {
	gorm.Model
	ID_Tg    int  `json:"id_tg" gorm:"unique_index"`
	IsActive bool `json:"isActive" gorm:"default:true"`
}

func NewStore(databaseUrl string) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Client{})
	return &Store{db: db}, nil
}

func (s *Store) Create(tg_id int) error {
	client := &Client{ID_Tg: tg_id}
	tx := s.db.Create(&client)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
