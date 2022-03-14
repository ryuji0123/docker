package main

import (
	"docker/host/opts"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
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

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	logrus.Info("start server")
	if err := http.ListenAndServe("localhost:"+opts.DefaultHTTPPort, Log(router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
}
