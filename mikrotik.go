func AnalyserMikrotik(config string) string {
    var result []string

    // Audit existant
    if strings.Contains(config, "telnet") && !strings.Contains(config, "disabled=yes") {
        result = append(result, "Faille : Telnet actif.")
    }

    // NOUVEAU : Module Syslog
    if !strings.Contains(config, "/system logging") {
        result = append(result, "Faille : Système de logs désactivé.")
    }

    // NOUVEAU : Module VLANs
    if !strings.Contains(config, "/interface vlan") {
        result = append(result, "Attention : Pas de segmentation VLAN détectée.")
    }

    if len(result) == 0 { return "Audit Mikrotik : Configuration robuste." }
    return "Audit Mikrotik : " + strings.Join(result, " | ")
}

