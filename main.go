package main

import(
  // "fmt"
  "log"
  "net/http"
)

func main() {
    h:= http.NewServeMux()

    h.HandleFunc("/track", tracker)
    h.HandleFunc("/", errHandler)

    log.Println("Lilbro started.")
    err := http.ListenAndServe(":9999",h)
    if err != nil {
      log.Fatal(err)
    }
}

func tracker(w http.ResponseWriter, r *http.Request) {
  timestamp := r.URL.Query().Get("ts")
  user := r.URL.Query().Get("u")
  context := r.URL.Query().Get("ctx")
  action := r.URL.Query().Get("axn")
  log.Printf("%v,%s,%s,%s",timestamp,user,context,action)
}

func errHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404);
  log.Print("Attempt to access url: %v", r.URL )
}
