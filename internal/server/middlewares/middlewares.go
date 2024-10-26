package middlewares

import (
	"net/http"
	"time"
	"log"
)


func TimeMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // перед выполнением хендлера 
        h.ServeHTTP(w, r)
		end := time.Now() // после выполнения хендлера
		log.Print("Completed in ", end.Second() - start.Second())
    })
}