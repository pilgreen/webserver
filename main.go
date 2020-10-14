package main

import (
  "flag"
  "log"
  "net/http"
  "os"
  "sync"
)

var Config struct {
  Directory string
  Port string
}

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)

  wd, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  flag.StringVar(&Config.Directory, "d", wd, "root directory")
  flag.StringVar(&Config.Port, "p", "1313", "webserver port")
  flag.Parse()

  // Run it
  port := ":" + Config.Port
  go func() {
    server := &http.Server {
      Addr: port,
      Handler: addCORS(http.FileServer(http.Dir(Config.Directory))),
    }

    log.Fatal(server.ListenAndServe())
    wg.Done()
  }()

  // Log it
  log.Printf("webserver started. go to http://localhost%s.", port)
  wg.Wait()
}

func addCORS(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Methods", "GET OPTIONS")
    h.ServeHTTP(w, r)
  }
}
