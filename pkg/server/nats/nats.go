package nats

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/nats-io/go-nats"

	"simpleNews/pkg/config"
	"simpleNews/pkg/service"
)

// Server implements NATS income serveri for storing news
type Server struct {
	nc         *nats.Conn
	natsAddr   string
	cancelConnectToNATS     context.CancelFunc
	ctxServer     context.Context
	cancelServer     context.CancelFunc
	subscribers Subscribers
	subscriptions []*nats.Subscription
	srv 	*service.Service
}

// New create in memory repository for storing nodes
func New(cfg config.Config, srv *service.Service) *Server {

	ctxConnectToNATS, cancelConnectToNATS := context.WithCancel(context.Background())
	ctxServer, cancelServer := context.WithCancel(context.Background())
	s := &Server{
		nc:       nil,
		natsAddr: cfg.NATSAddr,
		cancelConnectToNATS:   cancelConnectToNATS,
		ctxServer: ctxServer,
		cancelServer: cancelServer,
		subscriptions: []*nats.Subscription{},
		srv: srv,
	}

	// Create Subscribers


	go s.connectToNATS(ctxConnectToNATS)

	return s
}

func (s *Server) connectToNATS(ctx context.Context) {
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
					continue
				} 
				s.nc = nc
				s.subscribers =  NewHandler(s.srv, s.nc)
				log.Info("connection to NATS is established")
				for key, subcriber := range s.subscribers {
					sub, err := s.nc.Subscribe(key, subcriber)
					if err != nil {
						log.WithFields(log.Fields{
							"sub": key,
							"err": err,
						}).Error("cannot subscribe")
					}
					s.subscriptions = append(s.subscriptions, sub)
				}
			} else {
				if !s.nc.IsConnected() {
					log.Warn("lost connection to NATS")
					s.nc = nil
				}
			}
		case <-ctx.Done():
			// Stop handle a NATS connection
			ticker.Stop()
			// Unsubscribe from all subscriptions
			for _, sub := range s.subscriptions {
				if err := sub.Unsubscribe(); err != nil {
					log.WithFields(log.Fields{
						"sub": sub,
						"err": err,
					}).Error("cannot unsubscribe")
				}
			}
			// Close a connection to NATS
			s.nc.Close()
		}
	}
}

// Close closes connection to a database
func (s *Server) Close() {
	log.Warn("closing connection to NATS")
	s.cancelConnectToNATS()
	s.cancelServer()
}

// Run is a dummy fuction for that implementation 
func (s *Server) Run() {
	<-s.ctxServer.Done()
}

