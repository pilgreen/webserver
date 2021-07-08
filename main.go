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

    server := &http.Server {
      Addr: port,
      Handler: http.FileServer(root),
    }

    log.Fatal(server.ListenAndServe())
    wg.Done()
  }()

  // Log it
  log.Printf("webserver started at http://localhost%s.", port)
  wg.Wait()
}
