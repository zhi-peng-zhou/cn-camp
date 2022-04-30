package main

import (
	"encoding/json"
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)
var xlogger = log.Default()

func logHandler(h http.Handler) http.Handler {
	fu := func(w http.ResponseWriter, r *http.Request) {
		// 原始请求
		h.ServeHTTP(w, r)
		reqUrl := r.URL.String()
		method := r.Method
		vals := reflect.ValueOf(w)
		iv := reflect.Indirect(vals)
		statusCode := iv.FieldByName("status")
		clientIp := exnet.ClientIP(r)
		logStr := fmt.Sprintf("[%s]:%s  statusCode:%d  ClientIp:%s  ",method, reqUrl,
		statusCode.Int(), clientIp)
		xlogger.Println(logStr)
	}
	return http.HandlerFunc(fu)
}


func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/header", setRequestHeader)
	mux.HandleFunc("/version", getVersion)
	handler := logHandler(mux)
	server := http.Server{
		Addr:":8080",
		Handler: handler,
	}
	server.ListenAndServe()
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "system is healthy, status:200")
}

func setRequestHeader(w http.ResponseWriter, r *http.Request) {
	defer func() {
		xlogger.Println(exnet.ClientIP(r))
		xlogger.Println()
	}()
	for k := range r.Header {
		w.Header().Add(k, r.Header.Get(k))
	}
	reqHeader, _ := json.Marshal(r.Header)
	respHeader, _ := json.Marshal(w.Header())
	io.WriteString(w, "reqHeader: "+fmt.Sprint(string(reqHeader)+"\n"))
	io.WriteString(w, "respHeader: "+fmt.Sprint(string(respHeader)+"\n"))
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "the version is "+os.Getenv("VERSION"))
}
