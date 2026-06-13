package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
		file, header, err := r.FormFile("config")
		if err != nil {
			http.Error(w, "Erreur d'upload", http.StatusBadRequest)
			return
		}
		defer file.Close()

		if !strings.HasSuffix(header.Filename, ".txt") && !strings.HasSuffix(header.Filename, ".cfg") {
			http.Error(w, "Format non autorisé", http.StatusBadRequest)
			return
		}

		content, _ := io.ReadAll(file)
		// On lance l'audit et on récupère le score et le rapport
		score, rapportBrut := AnalyserConfiguration(string(content))
		
		// On génère le fichier HTML
		fichierRapport := GenererRapport(score, rapportBrut)
		
		fmt.Fprintf(w, "<h2>Audit terminé avec un score de %d/100</h2>", score)
		fmt.Fprintf(w, "<a href='/%s'>Télécharger/Imprimer le rapport PDF</a>", fichierRapport)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", uploadHandler)
	// Permet d'accéder au rapport généré
	http.Handle("/rapport_final.html", http.FileServer(http.Dir(".")))
	fmt.Println("GoShield démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

