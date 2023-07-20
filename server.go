package web_framework

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server interface {
	http.Handler
	// Start to start service
	Start(addr string) error
	// Stop to stop service
	Stop() error
}

type HTTPOption func(h *HTTPServer)

type HTTPServer struct {
	srv  *http.Server
	stop func() error
}

// WithHTTPServerStop Set stop function
func WithHTTPServerStop(fn func() error) HTTPOption {
	return func(h *HTTPServer) {
		if fn == nil {
			fn = func() error {
				fmt.Println("1231231312")
				quit := make(chan os.Signal)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				log.Println("Shutdown Server ...")

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				// to do before close
				if err := h.srv.Shutdown(ctx); err != nil {
					log.Fatal("Server Shutdown:", err)
				}
				// to do after close
				select {
				case <-ctx.Done():
					log.Println("timeout of 5 seconds.")
				}
				return nil
			}
		}
		h.stop = fn
	}
}

func NewHTTP(opts ...HTTPOption) *HTTPServer {
	h := &HTTPServer{}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// ServeHTTP receive request, forward request
// receive request from frontend
// forward request from frontend to the framework
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

// Start to start service
func (h *HTTPServer) Start(addr string) error {
	h.srv = &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return h.srv.ListenAndServe()
}

// Stop to stop service
func (h *HTTPServer) Stop() error {
	fmt.Println("Stop called ...")
	return h.stop()
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", nil)
}
