package main

import (
	"strings"
)

// AnalyserMikrotik traite spécifiquement les fichiers de configuration Mikrotik
func AnalyserMikrotik(config string) string {
	var res []string

	// 1. Vérification du protocole Telnet (doit être désactivé)
	if strings.Contains(config, "telnet") && !strings.Contains(config, "disabled=yes") {
		res = append(res, "Faille : Telnet actif (non sécurisé).")
	}

	// 2. Vérification du Syslog (Traçabilité)
	if !strings.Contains(config, "/system logging") {
		res = append(res, "Faille : Système de logs non configuré.")
	}

	// 3. Vérification de la segmentation (VLANs)
	if !strings.Contains(config, "/interface vlan") {
		res = append(res, "Attention : Aucun VLAN détecté (réseau plat à risque).")
	}

	// 4. Vérification du Firewall (Protection périmétrique)
	if !strings.Contains(config, "/ip firewall filter") {
		res = append(res, "Faille : Aucune règle de pare-feu détectée.")
	}

	if len(res) == 0 {
		return "Audit Mikrotik : Configuration robuste."
	}

	return strings.Join(res, " | ")
}

