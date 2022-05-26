package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/javiersoto15/metrics/kibana-apm"
	"go.elastic.co/apm/module/apmhttp"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Test")
		w.WriteHeader(200)
	})
	http.ListenAndServe(":8080", apmhttp.Wrap(mux))
}

func generateRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
