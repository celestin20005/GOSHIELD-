package main

import "strings"

func AnalyserCisco(config string) string {
    // Spécifique à Cisco
    if !strings.Contains(config, "service password-encryption") {
        return "<li>Cisco : Mots de passe en clair détectés !</li>"
    }
    return "<li>Cisco : Sécurité des mots de passe OK.</li>"
}

