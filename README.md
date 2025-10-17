# Power4-Web

Un jeu de Puissance 4 moderne et élégant développé en Go avec une interface web responsive.

## Caractéristiques

- **Jeu complet** : Toutes les règles du Puissance 4 implémentées
- **Détection de victoire** : Vérification horizontale, verticale et diagonale
- **Détection d'égalité** : Gestion des matchs nuls
- **3 niveaux de difficulté** :
  - Facile : Grille 6x7 (classique)
  - Normal : Grille 6x9
  - Difficile : Grille 7x8
- 📱 **Responsive** : Adapté mobile, tablette et desktop
-  **Reset instantané** : Relancer une partie en un clic

##  Installation

### Prérequis

- Go 1.22.2+

### Étapes

1. **Cloner le repository**
```bash
git clone https://github.com/intox24/power4-web.git
cd power4-web
```

2. **Lancer le serveur**
```bash
go run main.go
```

3. **Ouvrir dans le navigateur**
```
http://localhost:8080
```

## Structure du projet

```
power4-web/
├── main.go                 # Point d'entrée de l'application
├── go.mod                  # Configuration du module Go
├── index.html              # Page d'accueil
├── README.md               # Documentation
├── src/
│   ├── server.go          # Configuration du serveur HTTP
│   └── puissance.go       # Logique du jeu
├── pages/
│   └── play.html          # Interface de jeu
└── static/
    ├── styles.css         # Styles page d'accueil
    ├── game.css           # Styles page de jeu
    └── home.css           # (optionnel) Styles alternatifs
```

## Comment jouer

1. **Démarrer une partie**
   - Entrez les noms des deux joueurs
   - Choisissez une difficulté (optionnel, par défaut : Facile)
   - Cliquez sur "Commencer la partie"

2. **Jouer**
   - Le Joueur 1 commence (jetons rouges 🔴)
   - Cliquez sur un bouton "Col 1/2/3..." pour placer votre jeton
   - Les jetons tombent automatiquement à la position la plus basse
   - Le Joueur 2 joue ensuite (jetons jaunes 🟡)

3. **Gagner**
   - Alignez 4 jetons de votre couleur horizontalement, verticalement ou en diagonale
   - Le jeu détecte automatiquement la victoire
   - Un message s'affiche avec le nom du gagnant 

4. **Match nul**
   - Si toutes les cases sont remplies sans gagnant
   - Un message "Match nul" s'affiche 

5. **Rejouer**
   - Cliquez sur "Nouvelle partie" pour recommencer
   - Cliquez sur "Accueil" pour changer de joueurs ou de difficulté

## Technologies utilisées

- **Backend** : Go 1.22.2
  - `net/http`
  - `html/template`
  
- **Frontend** : HTML, CSS
  - Design moderne
  - Animations CSS
  - Dégradés et effets 3D

## Auteur

Julien (intox24)
