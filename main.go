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
    log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(Config.Directory))))
    wg.Done()
  }()

  // Log it
  log.Printf("webserver started. go to http://localhost%s.", port)
  wg.Wait()
}
