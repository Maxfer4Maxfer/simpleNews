package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"simpleNews/pkg/service"
)

// News represents a busines logic structure in that output package
type News struct {
	UUID      uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

// NewHandler return a router for handling http request
func NewHandler(service *service.Service) http.Handler {
	// create router
	m := http.NewServeMux()

	// setup router path
	m.HandleFunc("/news", func(w http.ResponseWriter, req *http.Request) {
		handleNews(w, req, service)
	})
	m.HandleFunc("/news/", func(w http.ResponseWriter, req *http.Request) {
		handleNewsWithSlash(w, req, service)
	})
	return m
}

func handleNews(w http.ResponseWriter, req *http.Request, srv *service.Service) {
	switch req.Method {
	case "POST":
		// create new news and return UUID
		title := req.FormValue("title")
		UUID, err := srv.NewNews(title)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
			w.Write([]byte("\n"))
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(UUID))
		}
	case "GET":
		showAllNews(w, srv)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
	}

}

func handleNewsWithSlash(w http.ResponseWriter, req *http.Request, srv *service.Service) {
	switch req.Method {
	case "GET":
		// parse parameters
		params := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

		switch len(params) {
		case 1: // Show all news
			showAllNews(w, srv)
		case 2: // Show the only one news
			showOneNews(w, params[1], srv)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 bad request"))
		}
	case "DELETE":
		err := srv.DeleteAllNews()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
			w.Write([]byte("\n"))
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All news were deleted"))

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request"))
	}
}

func showAllNews(w http.ResponseWriter, srv *service.Service) {
	news, err := srv.AllNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error" + "\n" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, n := range news {
		w.Write([]byte(n.UUID.String()))
		w.Write([]byte(" | "))
		w.Write([]byte(n.Title))
		w.Write([]byte(" | "))
		w.Write([]byte(n.Timestamp.String()))
		w.Write([]byte("\n"))

	}
}

func showOneNews(w http.ResponseWriter, uuid string, srv *service.Service) {
	news, err := srv.News(uuid)
	if err != nil {
		switch {
		case err.Error() == "invalid UUID format":
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 bad request" + "\n" + err.Error()))
		case strings.Contains(err.Error(), "invalid UUID length:"):
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 bad request" + "\n" + err.Error()))
		case err == service.ErrNewsNotFound:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found" + "\n" + err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error" + "\n" + err.Error()))
		}
		return
	}

	// convert service.News to a local news structure
	lNews := News{
		UUID:      news.UUID,
		Title:     news.Title,
		Timestamp: news.Timestamp,
	}

	// convert to json and write a responce
	b, err := json.Marshal(lNews)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error" + "\n" + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
