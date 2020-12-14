package server

import (
	"fmt"
	"net/http"
)

func PlayServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
