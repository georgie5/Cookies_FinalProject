package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/georgie5/Aldric2023/CookiesFinalProject/encoding/internal/cookies" // Import the internal/cookies package.
)

func main() {
	// Start a web server with the two endpoints.
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setCookieHandler)
	mux.HandleFunc("/get", getCookieHandler)

	log.Print("Listening...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the cookie as normal.
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello Zoë!✔",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	// Write the cookie. If there is an error (due to an encoding failure or it
	// being too long) then log the error and send a 500 Internal Server Error
	// response.
	err := cookies.Write(w, cookie)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("cookie set!"))
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Use the Read() function to retrieve the cookie value, additionally
	// checking for the ErrInvalidValue error and handling it as necessary.
	value, err := cookies.Read(r, "exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		case errors.Is(err, cookies.ErrInvalidValue):
			http.Error(w, "invalid cookie", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(value))
}
