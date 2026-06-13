package main

import (
	"strings"
)

// AnalyserCisco traite spécifiquement les fichiers de configuration Cisco
func AnalyserCisco(config string) string {
	var res []string

	// 1. Vérification du chiffrement des mots de passe
	if !strings.Contains(config, "service password-encryption") {
		res = append(res, "Faille : Mots de passe non chiffrés (service password-encryption absent).")
	}

	// 2. Vérification du Syslog (Traçabilité)
	if !strings.Contains(config, "logging host") {
		res = append(res, "Faille : Syslog non configuré, aucune traçabilité.")
	}

	// 3. Vérification de la segmentation (VLANs)
	if !strings.Contains(config, "vlan") {
		res = append(res, "Attention : Aucun VLAN détecté (réseau plat à risque).")
	}

	// 4. Vérification SSH (Accès sécurisé)
	if !strings.Contains(config, "transport input ssh") {
		res = append(res, "Faille : SSH non forcé sur les lignes VTY.")
	}

	if len(res) == 0 {
		return "Audit Cisco : Configuration robuste."
	}
	
	return strings.Join(res, " | ")
}

