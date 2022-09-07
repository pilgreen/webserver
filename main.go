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
  flag.StringVar(&Config.Port, "p", "3000", "webserver port")
  flag.Parse()

  // Run it
  port := ":" + Config.Port
  go func() {
    root := http.Dir(Config.Directory)
    fs := http.FileServer(root)

    server := &http.Server {
      Addr: port,
      Handler: cors(fs),
    }

    log.Fatal(server.ListenAndServe())
    wg.Done()
  }()

  // Log it
  log.Printf("webserver started at http://localhost%s.", port)
  wg.Wait()
}

/*
Enable CORS for serving from multiple ports
*/

func cors(fs http.Handler) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    fs.ServeHTTP(w, r)
  }
}
