package handlers

import (
	"fmt"
	"log"
	"net/http"
)

var A, B, C int

type Sum struct {
	l *log.Logger
}

func NewSum(l *log.Logger) *Sum {
	return &Sum{l}
}

func (s *Sum) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	C = (A + B)
	fmt.Fprintf(rw, "your result: %d", C)

}
