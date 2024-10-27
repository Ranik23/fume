package handlers

import (
	"fmt"
	"net/http"
)


func HomeHandler(w http.ResponseWriter, r * http.Request) { // TODO
	fmt.Fprint(w, "Hello, World!\n")
}