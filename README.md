# Ultra-Fast Port Scanner

Un scanner de ports haute performance écrit en Go utilisant une concurrence massive (Goroutines).

## Fonctionnalités

- **Concurrence Massive** : Scanne des milliers de ports en quelques secondes en utilisant des Goroutines légères.
- **Pool de Workers** : Gère efficacement les ressources avec un pool de workers basé sur un sémaphore (par défaut 1000 workers).
- **Sortie Claire** : Affiche les résultats dans un tableau formaté.

## Installation

1.  Clonez le dépôt :
    ```bash
    git clone <url-du-depot>
    cd ultra-fast-port-scanner
    ```
2.  Compilez le projet (nécessite Go 1.21+) :
    ```bash
    go build -o scanner cmd/scanner/main.go
    ```

## Utilisation

Exécutez le binaire compilé :

```bash
./scanner -host example.com -start 1 -end 1024
```

### Options

- `-host` : Hôte cible ou adresse IP (par défaut : `localhost`).
- `-start` : Port de début (par défaut : `1`).
- `-end` : Port de fin (par défaut : `1024`).

## Performance

Capable de scanner 1000 ports en moins de 3 secondes dans des conditions réseau standard.

## Structure du Projet

```
ultra-fast-port-scanner/
├── cmd/
│   └── scanner/
│       └── main.go     # Point d'entrée
├── pkg/
│   ├── scan/
│   │   ├── scan.go        # Module de scan des ports
│   │   └── scan_test.go   # Tests unitaires
│   └── report/
│       ├── report.go        # Module de rapport
│       └── report_test.go   # Tests unitaires
├── go.mod
├── README.md
└── LICENSE
```

## Licence

Copyright (c) 27 Janvier 2026 - Antigravity
Voir [LICENSE](LICENSE) pour plus de détails.
