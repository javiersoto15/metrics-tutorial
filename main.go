package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"go.elastic.co/apm/module/apmchi/v2"

	"github.com/javiersoto15/metrics/kibana-apm"
)

func main() {
	go func() {
		for {
			iterations := generateRandomInt(2000, 5000)
			for i := 0; i < iterations; i++ {
				kibana.ProcessExample()
			}
			sec := generateRandomInt(0, 50)
			<-time.After(time.Duration(sec) * time.Second)

		}
	}()
	go func() {
		for {
			iterations := generateRandomInt(2000, 5000)
			for i := 0; i < iterations; i++ {
				kibana.ReadExample()
			}
			sec := generateRandomInt(0, 50)
			<-time.After(time.Duration(sec) * time.Second)
		}
	}()
	r := chi.NewRouter()
	r.Use(apmchi.Middleware())
	http.ListenAndServe(":3000", r)
}

func generateRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
