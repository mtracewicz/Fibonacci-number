package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", fib)
	fmt.Println("Listening on port: 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error")
	}
}

func fib(w http.ResponseWriter, r *http.Request) {
	var resultForUser result

	n, err := strconv.Atoi(r.URL.String()[1:])
	if err != nil || n < 0 {
		handleInvalidRequest(w, r.URL.String())
		return
	}

	calculateFib(&resultForUser, n)
	respond(w, &resultForUser)
}

func handleInvalidRequest(w http.ResponseWriter, url string) {
	w.WriteHeader((http.StatusBadRequest))
	io.WriteString(w, "Must be a positive integer!")
	log.Printf("Invalid request URL: %s!", url)
}

func calculateFib(resultForUser *result, n int) {
	var wg sync.WaitGroup
	wg.Add(4)
	startI := time.Now()
	go func() {
		resultForUser.Iterative = fmt.Sprintf("%d It took me:%s", iterativeFib(n), time.Since(startI))
		wg.Done()
	}()
	startI2 := time.Now()
	go func() {
		resultForUser.IterativeFibV2 = fmt.Sprintf("%d It took me:%s", iterativeFibV2(n), time.Since(startI2))
		wg.Done()
	}()
	startR := time.Now()
	go func() {
		resultForUser.Recursive = fmt.Sprintf("%d It took me:%s", recursiveFib(n), time.Since(startR))
		wg.Done()
	}()
	startS := time.Now()
	go func() {
		resultForUser.Slices = fmt.Sprintf("%d it took me: %s ", slicesFib(n), time.Since(startS))
		wg.Done()
	}()
	wg.Wait()
}

func respond(w http.ResponseWriter, resultForUser *result) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusAccepted))
	marshaled, _ := json.Marshal(resultForUser)
	w.Write(marshaled)
	log.Println(string(marshaled))
}
