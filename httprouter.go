package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type respStruct struct {
    Greet string `json:"greeting"`
}

type reqStruct struct {
	Nam string `json:"name"`
}

func helloPost(respw http.ResponseWriter, req *http.Request, param httprouter.Params) {
    decoder := json.NewDecoder(req.Body)
    var request reqStruct
    var response respStruct   
    decoder.Decode(&request)
    response.Greet = "Hello,"+request.Nam+" !"
    json.NewEncoder(respw).Encode(response)
}

func hello(respw http.ResponseWriter, req *http.Request, param httprouter.Params) {
    fmt.Fprintf(respw, "Hello, %s!\n", param.ByName("name"))
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", helloPost)
    serve := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    serve.ListenAndServe()
}