package inmem

import (
	"github.com/google/uuid"

	"simpleNews/pkg/config"
	"simpleNews/pkg/service"
)

// Storage stores objects in memory
type Storage struct {
	news map[uuid.UUID]service.News
}

// New returns a storage object
func New(cfg config.Config) *Storage {
	s := &Storage{
		news: make(map[uuid.UUID]service.News),
	}
	return s
}

// SaveNews saves a news
func (s *Storage) SaveNews(t service.News) error {
	s.news[t.UUID] = t
	return nil
}

// News returns a news from the storage
func (s *Storage) News(ID uuid.UUID) (service.News, error) {
	_, ok := s.news[ID]
	if !ok {
		return service.News{}, service.ErrNewsNotFound
	}
	return s.news[ID], nil
}

// AllNews returns all news
func (s *Storage) AllNews() ([]service.News, error) {
	r := []service.News{}
	for _, t := range s.news {
		r = append(r, t)
	}
	return r, nil
}

// DeleteAllNews deletes all news
func (s *Storage) DeleteAllNews() error {
	s.news = make(map[uuid.UUID]service.News)
	return nil
}

// Close is an empty function for closing inmem storage
func (s *Storage) Close() {
}
