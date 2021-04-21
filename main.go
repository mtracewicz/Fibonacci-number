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

	"github.com/mtracewicz/fibonacci-number/fib"
)

func main() {
	http.HandleFunc("/", fibonacci)
	fmt.Println("Listening on port: 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error")
	}
}

func fibonacci(w http.ResponseWriter, r *http.Request) {
	var resultForUser result

	n, err := strconv.Atoi(r.URL.String()[1:])
	if err != nil || n < 0 {
		handleInvalidRequest(w, r.URL.String())
		return
	}

	calculateFib(&resultForUser, n)
	if !isResultValid(resultForUser) {
		handleInvalidRequest(w, r.URL.String())
		return
	}
	respond(w, &resultForUser)
}

func calculateFib(resultForUser *result, n int) {
	var wg sync.WaitGroup
	wg.Add(4)

	startI := time.Now()
	go func() {
		defer completeCalculation(&wg)
		res, err := fib.IterativeFib(n)
		if err != nil {
			panic("Error: Must be a positive integer!")
		}
		resultForUser.Iterative = fmt.Sprintf("%d It took me:%s", res, time.Since(startI))
	}()

	startI2 := time.Now()
	go func() {
		defer completeCalculation(&wg)
		res, err := fib.IterativeFibV2(n)
		if err != nil {
			panic("Error: Must be a positive integer!")
		}
		resultForUser.IterativeFibV2 = fmt.Sprintf("%d It took me:%s", res, time.Since(startI2))
	}()

	startR := time.Now()
	go func() {
		defer completeCalculation(&wg)
		res, err := fib.RecursiveFib(n)
		if err != nil {
			panic("Error: Must be a positive integer!")
		}
		resultForUser.Recursive = fmt.Sprintf("%d It took me:%s", res, time.Since(startR))
	}()

	startS := time.Now()
	go func() {
		defer completeCalculation(&wg)
		res, err := fib.SlicesFib(n)
		if err != nil {
			panic("Error: Must be a positive integer!")
		}
		resultForUser.Slices = fmt.Sprintf("%d it took me: %s ", res, time.Since(startS))
	}()

	wg.Wait()
}

func completeCalculation(wg *sync.WaitGroup) {
	wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func respond(w http.ResponseWriter, resultForUser *result) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusAccepted))
	marshaled, _ := json.Marshal(resultForUser)
	w.Write(marshaled)
	log.Println(string(marshaled))
}

func isResultValid(res result) bool {
	return res.Iterative != "" && res.IterativeFibV2 != "" && res.Recursive != "" && res.Slices != ""
}

func handleInvalidRequest(w http.ResponseWriter, url string) {
	w.WriteHeader((http.StatusBadRequest))
	io.WriteString(w, "Error: Must be a positive integer!")
	log.Printf("Invalid request URL: %s!", url)
}
