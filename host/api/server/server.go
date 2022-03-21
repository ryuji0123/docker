package server

import (
	"encoding/json"
	"fmt"
	"github.com/docker/host/api/types"
	"github.com/docker/host/opts"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Server struct {
}

// New returns a new instance of the server
func New() *Server {
	return &Server{}
}

func (s *Server) ServeAPI() {
	router := http.NewServeMux()
	router.HandleFunc("/commands", handler)
	logrus.Info("start server")
	if err := http.ListenAndServe("localhost:"+opts.DefaultHTTPPort, Log(router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Only POST request is acceptable\n")
		return
	}
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please set Content-type correctly\n")
		return
	}

	byteArray, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var jsonBody types.CmdJsonBody
	err = json.Unmarshal(byteArray, &jsonBody)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, jsonBody.Cmd+" command received\n")
}

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"addr":   r.RemoteAddr,
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("request received")
		h.ServeHTTP(w, r)
	})
}
