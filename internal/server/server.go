package server

import (
	"crypto/tls"
	"forum/pkg"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	cer, err := tls.LoadX509KeyPair("cert/forum.pem", "cert/forum_key.pem")
	if err != nil {
		log.Fatal("SSL error: ", err)
		return err
	}

	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{cer},
			InsecureSkipVerify: true,
		},
	}

	pkg.InfoLog.Printf("Server run on https://localhost%s", port)
	return s.httpServer.ListenAndServeTLS("", "")
}
