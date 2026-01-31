// =============================================================================
// @file main.go
// @description Point d'entrée principal du Scanner de Ports Ultra-Rapide
//
// Ce programme scanne les ports d'une cible (hôte ou IP) en utilisant
// la concurrence massive via les Goroutines Go pour une performance optimale.
//
// @author Alexandre Albert Ndour
// @version 1.0.0
// @date Janvier 2026
// =============================================================================

package main

// Importation des packages nécessaires
import (
	"flag"                              // Parsing des arguments en ligne de commande
	"fmt"                                // Formatage des entrées/sorties
	"time"                               // Mesure du temps d'exécution
	"ultra-fast-port-scanner/pkg/report" // Module de génération de rapports
	"ultra-fast-port-scanner/pkg/scan"   // Module de scan de ports
)

// -----------------------------------------------------------------------------
// Fonction principale
// -----------------------------------------------------------------------------

func main() {
	// Enregistrement du temps de début pour mesurer la durée du scan
	debut := time.Now()

	// -----------------------------------------------------------------------------
	// Configuration des arguments en ligne de commande
	// -----------------------------------------------------------------------------

	// -host : Hôte cible (nom de domaine ou adresse IP)
	hostPtr := flag.String("host", "localhost", "Hostname ou IP cible")

	// -start : Port de début de la plage à scanner
	startPortPtr := flag.Int("start", 1, "Port de début")

	// -end : Port de fin de la plage à scanner
	endPortPtr := flag.Int("end", 1024, "Port de fin")

	// Parsing des arguments
	flag.Parse()

	// -----------------------------------------------------------------------------
	// Exécution du scan de ports
	// -----------------------------------------------------------------------------

	// Affichage du démarrage du scan
	fmt.Printf("Démarrage du Scanner de Ports Ultra-Rapide pour %s (%d-%d)...\n",
		*hostPtr, *startPortPtr, *endPortPtr)

	// Scan des ports ouverts
	portsOuverts := scan.ScanPorts(*hostPtr, *startPortPtr, *endPortPtr)

	// -----------------------------------------------------------------------------
	// Affichage des résultats
	// -----------------------------------------------------------------------------

	// Calcul de la durée du scan
	duree := time.Since(debut)

	// Affichage du résumé
	fmt.Printf("\nScan terminé en %v. %d ports ouverts trouvés.\n",
		duree, len(portsOuverts))

	// Affichage du rapport détaillé
	report.ImprimerRapport(*hostPtr, portsOuverts)
}
