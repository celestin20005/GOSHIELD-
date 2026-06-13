package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Anonymiser masque les adresses IP pour protéger la confidentialité des données
func Anonymiser(texte string) string {
	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	return re.ReplaceAllString(texte, "xxx.xxx.xxx.xxx")
}

// extraireFailles compte le nombre de "Faille" dans le rapport pour calculer le score
func extraireFailles(rapport string) int {
	return strings.Count(rapport, "Faille")
}

// AnalyserConfiguration est le point d'entrée principal pour l'audit
func AnalyserConfiguration(config string) string {
	configLower := strings.ToLower(config)
	var rapport string

	// Détection automatique de l'équipement
	if strings.Contains(configLower, "hostname") {
		rapport = AnalyserCisco(config)
	} else if strings.Contains(configLower, "/ip address") || strings.Contains(configLower, "mikrotik") {
		rapport = AnalyserMikrotik(config)
	} else {
		return "Équipement non reconnu ou format non supporté."
	}

	// Calcul du score basé sur les failles détectées
	nombreFailles := extraireFailles(rapport)
	score := 100 - (nombreFailles * 20)
	if score < 0 {
		score = 0
	}

	// Détermination de la couleur pour le rendu visuel
	couleur := "green"
	if score < 50 {
		couleur = "red"
	} else if score < 80 {
		couleur = "orange"
	}

	// Construction du rapport final
	resultat := fmt.Sprintf(`
		<h2>Score de sécurité : <span style="color:%s">%d/100</span></h2>
		<h3>Détails de l'audit :</h3>
		<p>%s</p>
	`, couleur, score, rapport)

	// Retourne le résultat anonymisé
	return Anonymiser(resultat)
}

