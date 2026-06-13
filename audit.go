package main

import (
	"fmt"
	"strings"
)

// AnalyserConfiguration gère le score et le rapport
func AnalyserConfiguration(config string) string {
	config = strings.ToLower(config)
	var failles []string
	
	// On détecte la marque pour choisir le driver
	var rapport string
	if strings.Contains(config, "cisco") || strings.Contains(config, "hostname") {
		rapport = AnalyserCisco(config)
		failles = extraireFailles(rapport)
	} else if strings.Contains(config, "mikrotik") || strings.Contains(config, "/ip address") {
		rapport = AnalyserMikrotik(config)
		failles = extraireFailles(rapport)
	} else {
		return "<h2>Équipement non reconnu.</h2>"
	}

	// Calcul du score (Base 100, chaque faille enlève 20 points)
	score := 100 - (len(failles) * 20)
	if score < 0 { score = 0 }

	// Mise en forme du résultat
	couleur := "green"
	if score < 50 { couleur = "red" } else if score < 80 { couleur = "orange" }

	return fmt.Sprintf(`
		<h2>Score de sécurité : <span style="color:%s">%d/100</span></h2>
		<h3>Détails de l'audit :</h3>
		%s
	`, couleur, score, rapport)
}

// Petite fonction utilitaire pour compter les failles
func extraireFailles(rapport string) []string {
	// On compte les occurrences de "Faille" ou "Attention" ou "Alerte"
	var list []string
	if strings.Contains(rapport, "Faille") { list = append(list, "faille") }
	if strings.Contains(rapport, "Attention") { list = append(list, "attention") }
	if strings.Contains(rapport, "Alerte") { list = append(list, "alerte") }
	return list
}

