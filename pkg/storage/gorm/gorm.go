package gorm

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"simpleNews/pkg/config"
	"simpleNews/pkg/service"
)

// News represents a business object
type News struct {
	UUID      uuid.UUID `gorm:"primary_key"`
	Title     string
	Timestamp time.Time
}

// Storage implements SQL storage for news
type Storage struct {
	db     *gorm.DB
	dsn    string
	cancel context.CancelFunc
}

// New create in memory repository for storing nodes
func New(cfg config.Config) *Storage {

	ctx, cancel := context.WithCancel(context.Background())
	s := &Storage{
		db:     nil,
		dsn:    cfg.DSN,
		cancel: cancel,
	}

	go s.connectToDB(ctx)

	return s
}

func (s *Storage) connectToDB(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if s.db == nil {
				db, err := gorm.Open("mysql", s.dsn)
				if err != nil {
					fmt.Println("db", "mysql", "message", "got an error", "err", err)
				} else {
					fmt.Println("db", "mysql", "message", "connection is established")
					s.db = db
					s.db.AutoMigrate(&News{})
				}

			}
			if s.db != nil {
				err := s.db.DB().Ping()
				if err != nil {
					fmt.Println("db", "mysql", "message", "lost connection to the db", "err", err)
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			s.db.Close()
		}
	}
}

// Close closes connection to a database
func (s *Storage) Close() {
	s.cancel()
}

// SaveNews creates or saves a news in a database
func (s *Storage) SaveNews(n service.News) error {
	// convert datatypes from different packages
	ln := News{
		UUID:      n.UUID,
		Title:     n.Title,
		Timestamp: n.Timestamp,
	}
	if s.db != nil {
		s.db.Save(&ln)
		return nil
	}
	return service.ErrStorageUnavailable
}

// News returns a news by UUID
func (s *Storage) News(ID uuid.UUID) (service.News, error) {
	n := News{}
	if s.db != nil {
		var count int
		s.db.Where("UUID = ?", ID).First(&n).Count(&count)
		if count == 0 {
			return service.News{}, service.ErrNewsNotFound
		}
		// convert datatypes from different packages
		r := service.News{
			UUID:      n.UUID,
			Title:     n.Title,
			Timestamp: n.Timestamp,
		}
		return r, nil
	}
	return service.News{}, service.ErrStorageUnavailable
}

// AllNews returns all news
func (s *Storage) AllNews() ([]service.News, error) {
	lnews := []News{}
	rnews := []service.News{}

	if s.db != nil {
		s.db.Find(&lnews)
		// convert datatypes from different packages
		for _, n := range lnews {
			rnews = append(rnews, service.News{
				UUID:      n.UUID,
				Title:     n.Title,
				Timestamp: n.Timestamp,
			})
		}

		return rnews, nil
	}
	return rnews, service.ErrStorageUnavailable
}

// DeleteAllNews deletes all news
func (s *Storage) DeleteAllNews() error {
	s.db.Where("UUID LIKE ?", "%").Delete(service.News{})
	return nil
}
