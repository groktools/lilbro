package main

import(
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  // open log file
  tlog,err := os.OpenFile("lilbro.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
  if err!=nil{
    panic(err)
  }
  defer tlog.Close()

  h:= http.NewServeMux()

  h.HandleFunc("/track", tracker(tlog))
  h.HandleFunc("/", errHandler)

  log.Println("Lilbro started.")
  err = http.ListenAndServe(":9999",h)
  if err != nil {
    log.Fatal(err)
  }
}

func tracker(f *os.File) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request){
    timestamp := r.URL.Query().Get("ts")
    user := r.URL.Query().Get("u")
    context := r.URL.Query().Get("ctx")
    action := r.URL.Query().Get("axn")
    logTxt :=fmt.Sprintf("%v,%s,%s,%s\n",timestamp,user,context,action)
    f.WriteString(logTxt)
  }
}

func errHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404);
  log.Print("Attempt to access url: %v", r.URL )
}
