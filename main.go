/*
 */
package main

import (
	"log"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("COSMiddleware...")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Yi-Auth-Token")
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("COSMiddleware2...")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("COSMiddleware3...")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Yi-Auth-Token")
		next.ServeHTTP(w, r)
	})
}

type Middleware func(handler http.Handler) http.Handler

func NewMiddlewareChain(middlewares ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	middlewareChain := NewMiddlewareChain(CORSMiddleware, CORSMiddleware2, CORSMiddleware3)

	mux := http.NewServeMux()
	mux.Handle("/", middlewareChain(http.HandlerFunc(DefaultHandler)))

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
