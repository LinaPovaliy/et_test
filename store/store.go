package store

import (
	"database/sql"
	"errors"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTable() { //TODO можно было через миграции
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS thumbnails (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			video_id TEXT NOT NULL,
			image_data BLOB NOT NULL
		);`)
	if err != nil {
		fmt.Printf("Error creating table: %v", err)
	}
}

func (s *Store) SaveThumbnail(videoID string, imageData []byte) error {
	_, err := s.db.Exec(`INSERT INTO thumbnails (video_id, image_data) VALUES ($1, $2)`, videoID, imageData)
	if err != nil {
		return fmt.Errorf("failed to save thumbnail to DB: %v", err)
	}
	return nil
}

func (s *Store) CheckThumbnailInStore(videoID string) ([]byte, error) {
	var imageData []byte
	err := s.db.QueryRow(`SELECT image_data FROM thumbnails WHERE video_id = $1`, videoID).Scan(&imageData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no thumbnail found in store")
		}
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	return imageData, nil

}
