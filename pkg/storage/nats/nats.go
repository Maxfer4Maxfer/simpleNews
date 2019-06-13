package nats

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/nats-io/go-nats"
	log "github.com/sirupsen/logrus"

	pb "simpleNews/pb"
	"simpleNews/pkg/config"
	"simpleNews/pkg/service"
)

// Storage implements NATS storage for news
type Storage struct {
	nc       *nats.Conn
	natsAddr string
	cancel   context.CancelFunc
}

// New create in memory repository for storing nodes
func New(cfg config.Config) *Storage {

	ctx, cancel := context.WithCancel(context.Background())
	s := &Storage{
		nc:       nil,
		natsAddr: cfg.NATSAddr,
		cancel:   cancel,
	}

	go s.connectToNATS(ctx)

	return s
}

func (s *Storage) connectToNATS(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if s.nc == nil {
				// tryint to connect to NATS
				log.Info("trying connect to the NATS...")
				nc, err := nats.Connect(s.natsAddr)
				if err != nil {
					log.WithFields(log.Fields{
						"err": err,
					}).Error("got an error while trying to Conncet to NATS", "err")
				} else {
					log.Info("connection to NATS is established")
					s.nc = nc
				}
			} else {
				if !s.nc.IsConnected() {
					log.Warn("lost connection to NATS")
					s.nc = nil
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			s.nc.Close()
		}
	}
}

// Close closes connection to a database
func (s *Storage) Close() {
	log.Warn("closing connection to NATS")
	s.cancel()
}

// SaveNews creates or saves a news in a database
func (s *Storage) SaveNews(n service.News) error {
	// convert datatypes from different packages
	t, err := timestamp.TimestampProto(n.Timestamp)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("error while parsint Timestamp in SaveNews")
		return err
	}
	ln := &pb.News{
		Uuid:      n.UUID.String(),
		Title:     n.Title,
		Timestamp: t,
	}
	if s.nc.IsConnected() {
		dataPb, _ := proto.Marshal(ln)
		r, err := s.nc.Request("SaveNews", []byte(dataPb), 2*time.Second)

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("error while calling SaveNews")
			return err
		}
		// creating a structure for parsing a responce
		resp := &pb.SaveNewsResponce{}
		// encoding the response
		err = proto.Unmarshal(r.Data, resp)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("unmarshal a SaveNews responce")
			return err
		}
		// something wrong on a storage site
		if resp.Err != "" {
			log.WithFields(log.Fields{
				"err": resp.Err,
			}).Error("something wrong on the repository site")
			return service.ErrStorageUnavailable
		}

		log.WithFields(log.Fields{
			"uuid": n.UUID.String(),
		}).Info("the news save succesfully")
		return nil
	}
	return service.ErrStorageUnavailable
}

// News returns a news by UUID
func (s *Storage) News(ID uuid.UUID) (service.News, error) {
	if s.nc.IsConnected() {
		n := &pb.News{
			Uuid: ID.String(),
		}
		dataPb, _ := proto.Marshal(n)
		r, err := s.nc.Request("News", []byte(dataPb), 2*time.Second)

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("error while calling News")
			return service.News{}, err
		}
		// creating a structure for parsing a responce
		resp := &pb.NewsResponce{}
		// encoding the response
		err = proto.Unmarshal(r.Data, resp)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("unmarshal a News responce")
			return service.News{}, err
		}

		// The news is not found
		if resp.Err == service.ErrNewsNotFound.Error() {
			log.WithFields(log.Fields{
				"err": resp.Err,
			}).Error("The news is not found")
			return service.News{}, service.ErrNewsNotFound
		}

		// something wrong on a storage site
		if resp.Err != "" {
			log.WithFields(log.Fields{
				"err": resp.Err,
			}).Error("something wrong on the repository site")
			return service.News{}, service.ErrStorageUnavailable
		}

		// convert datatypes from different packages
		id, err := uuid.Parse(resp.News.Uuid)
		if err != nil {
			return service.News{}, err
		}
		ts, err := timestamp.Timestamp(resp.News.Timestamp)
		if err != nil {
			return service.News{}, err
		}

		result := service.News{
			UUID:      id,
			Title:     resp.News.Title,
			Timestamp: ts,
		}

		log.WithFields(log.Fields{
			"uuid": id,
		}).Info("the news requested succesfully")

		return result, nil
	}
	return service.News{}, service.ErrStorageUnavailable
}

// AllNews returns all news
func (s *Storage) AllNews() ([]service.News, error) {
	if s.nc.IsConnected() {
		r, err := s.nc.Request("AllNews", []byte(""), 2*time.Second)

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("error while calling AllNews")
			return []service.News{}, err
		}
		// creating a structure for parsing a responce
		resp := &pb.AllNewsResponce{}
		// encoding the response
		err = proto.Unmarshal(r.Data, resp)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("unmarshal a News responce")
			return []service.News{}, err
		}

		// something wrong on a storage site
		if resp.Err != "" {
			log.WithFields(log.Fields{
				"err": resp.Err,
			}).Error("something wrong on the repository site")
			return []service.News{}, service.ErrStorageUnavailable
		}

		// convert datatypes from different packages
		result := []service.News{}
		for _, n := range resp.News {
			id, err := uuid.Parse(n.Uuid)
			if err != nil {
				return []service.News{}, err
			}
			ts, err := timestamp.Timestamp(n.Timestamp)
			if err != nil {
				return []service.News{}, err
			}
			result = append(result, service.News{
				UUID:      id,
				Title:     n.Title,
				Timestamp: ts,
			})
		}

		log.Info("All news requested succesfully")

		return result, nil
	}
	return []service.News{}, service.ErrStorageUnavailable
}

// DeleteAllNews deletes all news
func (s *Storage) DeleteAllNews() error {
	if s.nc.IsConnected() {
		r, err := s.nc.Request("DeleteAllNews", []byte(""), 2*time.Second)

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("error while calling DeleteAllNews")
			return err
		}
		// creating a structure for parsing a responce
		resp := &pb.DeleteAllNewsResponce{}
		// encoding the response
		err = proto.Unmarshal(r.Data, resp)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("unmarshal a DeleteAllNews responce")
			return err
		}

		// something wrong on a storage site
		if resp.Err != "" {
			log.WithFields(log.Fields{
				"err": resp.Err,
			}).Error("something wrong on the repository site")
			return service.ErrStorageUnavailable
		}

		// All nodes have been deleted
		log.Info("All news deleted succesfully")

		return nil
	}
	return service.ErrStorageUnavailable
}
