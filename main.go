package main

import (
	"net/http"
	"fmt"
	// "log"
)

const (
	PORT string = ":8080"
	ASSETS string = "./web/assets"
	WEB string = "./web/index.html"
)

func ServeAssets() {
	assetsDir := http.Dir(ASSETS)

	fileHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))

	http.Handle("/assets/", fileHandler)
}

func ServeHTML() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, WEB)
	})
}

func main() {
	ServeAssets()
	ServeHTML()
	
	fmt.Printf("Server starting, Listening on port: %s \n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if (err != nil) {
		fmt.Printf("Error: %s", err)
	}
}