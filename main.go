package main

import (
	"first/handlers"

	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "calculator", log.LstdFlags)
	cs := handlers.NewSum(l)

	sm := http.NewServeMux()
	sm.Handle("/", cs)

	http.ListenAndServe(":9090", sm)

}
