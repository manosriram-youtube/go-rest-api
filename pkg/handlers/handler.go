package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l,
	}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Print("Handling Hello requests")

	bdy, err := ioutil.ReadAll(r.Body)
	h.l.Println(bdy)
	if err != nil {
		h.l.Printf("Error: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Hello, %s", bdy)
}
