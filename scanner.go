
// scanner.go
package main

// Analyseur définit le contrat que tout équipement doit respecter
type Analyseur interface {
	Analyser(contenu string) string
}

