package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "index.html")
		return
	}

	// Limiter la taille à 1 Mo
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
	
	file, header, err := r.FormFile("config")
	if err != nil {
		http.Error(w, "Erreur d'upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Vérification de l'extension
	if !strings.HasSuffix(header.Filename, ".txt") && !strings.HasSuffix(header.Filename, ".cfg") {
		http.Error(w, "Format non autorisé", http.StatusBadRequest)
		return
	}

	// Lecture en RAM uniquement
	content, _ := io.ReadAll(file)
	configString := string(content)

	// Lancer l'audit
	rapport := AnalyserConfiguration(configString)
	
	fmt.Fprintln(w, rapport)
}

func main() {
	http.HandleFunc("/", uploadHandler)
	fmt.Println("Serveur GoShield démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

