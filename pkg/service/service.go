package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Storage declares methods that the real storage object should implement
type Storage interface {
	SaveNews(n News) error
	News(ID uuid.UUID) (News, error)
	AllNews() ([]News, error)
	DeleteAllNews() error
}

// News represents a business object
type News struct {
	UUID      uuid.UUID
	Title     string
	Timestamp time.Time
}

// Service is a central component of the system. It contains all business logic.
type Service struct {
	storage Storage
}

var (
	// ErrNewsNotFound show that requested news cannot be found
	ErrNewsNotFound = errors.New("News not found")

	// ErrStorageUnavailable arised if something happened with a storage subsystem
	ErrStorageUnavailable = errors.New("Storage unavailable")
)

// New retunrs new Service
func New(store Storage) *Service {
	s := &Service{
		storage: store,
	}
	return s
}

// NewNews creates a new news and save it in a storage
func (s *Service) NewNews(title string) (string, error) {
	n := News{
		UUID:      uuid.New(),
		Title:     title,
		Timestamp: time.Now(),
	}

	err := s.SaveNews(n)
	if err != nil {
		return "", err
	}

	return n.UUID.String(), nil
}

// SaveNews saves a new
func (s *Service) SaveNews(n News) error {

	err := s.storage.SaveNews(n)
	if err != nil {
		return err
	}
	return nil
}

// News returns a news with provided guid
func (s *Service) News(ID string) (News, error) {
	UUID, err := uuid.Parse(ID)
	if err != nil {
		return News{}, err
	}

	return s.storage.News(UUID)
}

// AllNews returns a news with provided guid
func (s *Service) AllNews() ([]News, error) {
	return s.storage.AllNews()
}

// DeleteAllNews delete all news in a storage
func (s *Service) DeleteAllNews() error {
	return s.storage.DeleteAllNews()
}
