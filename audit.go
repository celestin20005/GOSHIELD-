package main

import "strings"

func AnalyserConfiguration(config string) string {
    // 1. Détection de la marque
    if strings.Contains(config, "version") && strings.Contains(config, "cisco") {
        return "<ul>" + AnalyserCisco(config) + "</ul>"
    } 
    if strings.Contains(config, "mikrotik") {
        return "<ul>" + AnalyserMikrotik(config) + "</ul>"
    }
    
    return "Équipement non supporté ou inconnu."
}

