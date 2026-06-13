package main

import (
	"html/template"
	"os"
	"regexp"
	"strings"
)

func Anonymiser(texte string) string {
	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	return re.ReplaceAllString(texte, "xxx.xxx.xxx.xxx")
}

func AnalyserConfiguration(config string) (int, string) {
	configLower := strings.ToLower(config)
	var rapport string

	if strings.Contains(configLower, "hostname") {
		rapport = AnalyserCisco(config)
	} else if strings.Contains(configLower, "/ip address") || strings.Contains(configLower, "mikrotik") {
		rapport = AnalyserMikrotik(config)
	} else {
		return 0, "Équipement non reconnu."
	}

	nombreFailles := strings.Count(rapport, "Faille")
	score := 100 - (nombreFailles * 20)
	if score < 0 { score = 0 }

	return score, Anonymiser(rapport)
}

func GenererRapport(score int, details string) string {
	listeFailles := strings.Split(details, " | ")
	htmlDetails := ""
	for _, faille := range listeFailles {
		htmlDetails += "<li>" + faille + "</li>"
	}

	data := struct {
		Date    string
		Score   int
		Details template.HTML
	}{
		Date:    "13 Juin 2026",
		Score:   score,
		Details: template.HTML(htmlDetails),
	}

	tmpl, _ := template.ParseFiles("rapport_template.html")
	f, _ := os.Create("rapport_final.html")
	defer f.Close()
	tmpl.Execute(f, data)
	return "rapport_final.html"
}

