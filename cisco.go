func AnalyserCisco(config string) string {
    var result []string

    // Audit existant
    if !strings.Contains(config, "service password-encryption") {
        result = append(result, "Faille : Mots de passe en clair.")
    }
    
    // NOUVEAU : Module Syslog
    if !strings.Contains(config, "logging host") {
        result = append(result, "Faille : Syslog non configuré (pas de traçabilité).")
    }

    // NOUVEAU : Module VLANs
    if !strings.Contains(config, "vlan") {
        result = append(result, "Attention : Aucun VLAN détecté (réseau plat à risque).")
    }

    if len(result) == 0 { return "Audit Cisco : Configuration robuste." }
    return "Audit Cisco : " + strings.Join(result, " | ")
}

