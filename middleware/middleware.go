package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"openthailand/constant"
	"openthailand/errors"
	"strings"
	"time"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func UseWithHook(handler http.HandlerFunc, hook func(http.HandlerFunc) http.HandlerFunc, middlewares []func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	if hook != nil {
		handler = hook(handler)
	}

	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func Use(handler http.HandlerFunc, middlewares []func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func GzipResult(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		h(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	}
}

func CheckApplicationKey(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		applicationID := r.Header.Get(constant.HTTP_HEADER_KEY_APPLICATION_ID)
		restAPIKey := r.Header.Get(constant.HTTP_HEADER_KEY_REST_API)
		if applicationID == "" || restAPIKey == "" {
			errors.NoApplicationKey().Write(w)
			return
		}
		h(w, r)
	}
}

func CachedResult(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", 3600))
		h(w, r)
	}
}

func JsonResult(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h(w, r)
	}
}

func JavascriptResult(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-javascript; charset=utf-8")
		h(w, r)
	}
}

func LoggingHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

type LogRecord struct {
	http.ResponseWriter
	Status  int
	Message string
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.Status = status
	//error is always in json format.
	if status != 0 && status != 200 {
		r.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	r.ResponseWriter.WriteHeader(status)
}
