package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// detecterEtAnalyser est le "cerveau" qui choisit le bon driver
func detecterEtAnalyser(contenu string) string {
	contenu = strings.ToLower(contenu)

	// Logique de détection de l'équipement
	if strings.Contains(contenu, "cisco") || strings.Contains(contenu, "hostname") {
		return AnalyserCisco(contenu)
	} else if strings.Contains(contenu, "mikrotik") || strings.Contains(contenu, "/ip address") {
		return AnalyserMikrotik(contenu)
	}
	
	return "Équipement non reconnu. Veuillez vérifier que le fichier est bien une configuration valide."
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// 1. Récupérer le fichier
	file, _, err := r.FormFile("configfile")
	if err != nil {
		http.Error(w, "Erreur lors de l'envoi du fichier", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. Lire le contenu
	buf := new(strings.Builder)
	io.Copy(buf, file)
	
	// 3. Analyser en appelant la fonction centralisée
	resultat := detecterEtAnalyser(buf.String())

	// 4. Afficher le résultat
	fmt.Fprintf(w, "<html><head><meta charset='UTF-8'></head><body><h1>Résultat de l'Audit</h1><p>%s</p><br><a href='/'>Retour</a></body></html>", resultat)
}

func main() {
	// Servir le fichier index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	
	// Gérer l'upload
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("Serveur GoShield prêt sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

