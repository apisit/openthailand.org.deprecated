package main

import (
	"flag"
	"log"
	"net/http"
	"openthailand/middleware"
	"openthailand/router"
	"os"
	"runtime"

	"openthailand/config"
	"time"

	"github.com/gorilla/mux"
)

var (
	HTTP_PORT = os.Getenv("HTTP_PORT")
)

//Server mux
type Server struct {
	r *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, DEL")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-TC-Rest-API-Key,X-TC-Application-Id")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	record := &middleware.LogRecord{
		ResponseWriter: w,
	}

	t1 := time.Now()
	s.r.ServeHTTP(record, r)
	t2 := time.Now()
	if record.Status != 0 && record.Status != 301 {
		log.SetPrefix("\x1b[31;1m[Error] ")
		log.Printf("[%s] %q %v %v\n", r.Method, r.URL.String(), record.Status, t2.Sub(t1))
		log.Printf("%v\n", r.Header)
	} else {
		log.SetPrefix("\x1b[0m[Info]  ")
		status := record.Status
		if record.Status == 0 {
			status = 200
		}
		log.Printf("[%s] %q %v %v\n", r.Method, r.URL.String(), status, t2.Sub(t1))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mode := flag.String("mode", "", "Run in which mode. local | staging | production")
	flag.Parse()

	pathPrefix := "/"
	r := mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	http.Handle("/", &Server{r})
	config.SetMode(*mode)

	if config.IS_LOCAL {
		//in local mode we enable these paths.
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))
		//end local mode.
	}

	if HTTP_PORT == "" {
		HTTP_PORT = "8088"
	}

	//config routing here
	router.RegisterRouting(r)
	log.Println("Listening... " + HTTP_PORT)
	http.ListenAndServe(":"+HTTP_PORT, nil)
}
