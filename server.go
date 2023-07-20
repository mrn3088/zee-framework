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

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type server interface {
	http.Handler
	// Start to start service
	Start(addr string) error
	// Stop to stop service
	Stop() error
	// addRouter a very core API, cannot be accessed by clients
	addRouter(method string, pattern string, handleFunc HandleFunc)
}

type HTTPOption func(h *HTTPServer)

type HTTPServer struct {
	srv  *http.Server
	stop func() error
	// routers store routes (temp!)
	routers map[string]HandleFunc
}

/*
{
	"GET-login": HandleFunc1,
	"POST-login": HandleFunc2,
	...
	...
}
*/

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
	h := &HTTPServer{
		routers: map[string]HandleFunc{},
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// ServeHTTP receive request, forward request
// receive request from frontend
// forward request from frontend to the framework
// connects frontend and backend
func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// 1. match route
	key := fmt.Sprintf("%s-%s", request.Method, request.URL.Path)
	handler, ok := h.routers[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 NOT FOUND"))
		return
	}
	// 2. send request
	handler(w, request)
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

// addRouter core route registration
// register routes when start
func (h *HTTPServer) addRouter(method string, pattern string, handleFunc HandleFunc) {
	// construct unique key
	key := fmt.Sprintf("%s-%s", method, pattern)
	fmt.Printf("add route %s-%s", method, pattern)
	h.routers[key] = handleFunc
}

// GET request
func (h *HTTPServer) GET(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodGet, pattern, handleFunc)
}

// POST request
func (h *HTTPServer) POST(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPost, pattern, handleFunc)
}

// DELETE request
func (h *HTTPServer) DELETE(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodDelete, pattern, handleFunc)
}

// PUT request
func (h *HTTPServer) PUT(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPut, pattern, handleFunc)
}
func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", nil)
}
