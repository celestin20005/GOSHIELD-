package main

import "strings"

// AnalyserMikrotik contient les règles spécifiques à Mikrotik
func AnalyserMikrotik(config string) string {
    var result []string

    // Règle 1 : Vérifier si le service Telnet est actif
    if strings.Contains(config, "/ip service") && strings.Contains(config, "telnet") && !strings.Contains(config, "disabled=yes") {
        result = append(result, "Faille : Service Telnet actif (non sécurisé).")
    }

    // Règle 2 : Vérifier si les mots de passe par défaut ou faibles sont gérés (exemple simple)
    if strings.Contains(config, "name=admin group=full") && !strings.Contains(config, "password=") {
        result = append(result, "Attention : Utilisateur 'admin' sans mot de passe défini.")
    }

    // Règle 3 : Vérifier le pare-feu (Firewall)
    if !strings.Contains(config, "/ip firewall filter") {
        result = append(result, "Alerte : Aucune règle de filtrage firewall détectée.")
    }

    if len(result) == 0 {
        return "Audit Mikrotik : Configuration sécurisée."
    }
    return "Audit Mikrotik : " + strings.Join(result, " | ")
}

