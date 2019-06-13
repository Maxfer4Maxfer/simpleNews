package nats

import (
	"github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/nats-io/go-nats"

	pb "simpleNews/pb"
	"simpleNews/pkg/service"
)

// Subscribers represents array of callback functions for NATS subscriptions
type Subscribers map[string]func(m *nats.Msg)

// NewHandler return a router for handling http request
func NewHandler(srv *service.Service, nc *nats.Conn) Subscribers {
	subs := make(map[string]func(m *nats.Msg))

	subs["News"] = func(m *nats.Msg) { subNews(m, srv, nc) }
	subs["SaveNews"] = func(m *nats.Msg) { subSaveNews(m, srv, nc) }
	subs["AllNews"] = func(m *nats.Msg) { subAllNews(m, srv, nc) }
	subs["DeleteAllNews"] = func(m *nats.Msg) { subDeleteAllNews(m, srv, nc) }
	return subs
}

// subNews gets a news by UUID provided in a NATS protobuf msg
func subNews(msg *nats.Msg, srv *service.Service, nc *nats.Conn) {
	// parse message
	ln := &pb.News{}
	resp := &pb.NewsResponce{}
	err := proto.Unmarshal(msg.Data, ln)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	n := service.News{}
	// Call a service method
	n, err = srv.News(ln.Uuid)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	// convert datatypes from different packages
	// service.News -> pb.News
	t, err := timestamp.TimestampProto(n.Timestamp)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}
	ln.Timestamp = t
	ln.Title = n.Title

	resp.News = ln

	// send reply
	dataPb, _ := proto.Marshal(resp)
	nc.Publish(msg.Reply, []byte(dataPb))
}

// subSaveNews save a new news in the storage
func subSaveNews(msg *nats.Msg, srv *service.Service, nc *nats.Conn) {
	// parse message
	ln := &pb.News{}
	resp := &pb.SaveNewsResponce{}
	err := proto.Unmarshal(msg.Data, ln)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}
	// convert datatypes from different packages
	// pb.News -> service.News
	id, err := uuid.Parse(ln.Uuid)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}
	ts, err := timestamp.Timestamp(ln.Timestamp)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	n := service.News{
		UUID:      id,
		Title:     ln.Title,
		Timestamp: ts,
	}

	// Call a service method
	err = srv.SaveNews(n)
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	// send reply
	dataPb, _ := proto.Marshal(resp)
	nc.Publish(msg.Reply, []byte(dataPb))
}

//subAllNews sends all news to the NATS reply message
func subAllNews(msg *nats.Msg, srv *service.Service, nc *nats.Conn) {
	resp := &pb.AllNewsResponce{}

	// Call a service method
	ns, err := srv.AllNews()
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	// convert datatypes from different packages
	// []service.News -> pb.AllNewsResponce
	lns := []*pb.News{}
	for _, n := range ns {
		ts, err := timestamp.TimestampProto(n.Timestamp)
		if err != nil {
			resp.Err = err.Error()
			dataPb, _ := proto.Marshal(resp)
			nc.Publish(msg.Reply, []byte(dataPb))
			return
		}
		lns = append(lns, &pb.News{
			Uuid:      n.UUID.String(),
			Title:     n.Title,
			Timestamp: ts,
		})
	}

	// send reply
	resp.News = lns
	dataPb, _ := proto.Marshal(resp)
	nc.Publish(msg.Reply, []byte(dataPb))
}

//subDeleteAllNews delete all news from the storage
func subDeleteAllNews(msg *nats.Msg, srv *service.Service, nc *nats.Conn) {
	resp := &pb.DeleteAllNewsResponce{}

	// Call a service method
	err := srv.DeleteAllNews()
	if err != nil {
		resp.Err = err.Error()
		dataPb, _ := proto.Marshal(resp)
		nc.Publish(msg.Reply, []byte(dataPb))
		return
	}

	// send a reply
	dataPb, _ := proto.Marshal(resp)
	nc.Publish(msg.Reply, []byte(dataPb))
}
