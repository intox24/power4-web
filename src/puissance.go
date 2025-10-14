package power4

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tableau [][]string
var JoueurActuel string = ""
var gagnant string = ""
var joueur1 string = ""
var joueur2 string = ""

func TableauJeu(rows, cols int) [][]string {
	tableau = make([][]string, rows)
	for i := range tableau {
		tableau[i] = make([]string, cols)
		for j := range tableau[i] {
			tableau[i][j] = "."
		}
	}
	return tableau
}

func JetonPlacer(tableau [][]string, col int, jeton string) bool {
	rows := len(tableau)
	for i := rows - 1; i >= 0; i-- {
		if tableau[i][col] == "." {
			tableau[i][col] = jeton
			return true
		}
	}
	return false
}

func VictoireHorizontale(tableau [][]string, jeton string) bool {
	rows := len(tableau)
	if rows == 0 {
		return false
	}
	cols := len(tableau[0])
	if cols < 4 {
		return false
	}
	for i := 0; i < rows; i++ {
		count := 0
		for j := 0; j < cols; j++ {
			if tableau[i][j] == jeton {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func VictoireVerticale(tableau [][]string, jeton string) bool {
	rows := len(tableau)
	if rows < 4 {
		return false
	}
	cols := len(tableau[0])
	if cols == 0 {
		return false
	}
	for j := 0; j < cols; j++ {
		count := 0
		for i := 0; i < rows; i++ {
			if tableau[i][j] == jeton {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func VictoireDiagonaleD(tableau [][]string, jeton string) bool {
	rows := len(tableau)
	if rows < 4 {
		return false
	}
	cols := len(tableau[0])
	if cols < 4 {
		return false
	}
	for i := 0; i <= rows-4; i++ {
		for j := 0; j <= cols-4; j++ {
			if tableau[i][j] == jeton && tableau[i+1][j+1] == jeton && tableau[i+2][j+2] == jeton && tableau[i+3][j+3] == jeton {
				return true
			}
		}
	}
	return false
}

func VictoireDiagonaleA(tableau [][]string, jeton string) bool {
	rows := len(tableau)
	if rows < 4 {
		return false
	}
	cols := len(tableau[0])
	if cols < 4 {
		return false
	}
	for i := 3; i < rows; i++ {
		for j := 0; j <= cols-4; j++ {
			if tableau[i][j] == jeton && tableau[i-1][j+1] == jeton && tableau[i-2][j+2] == jeton && tableau[i-3][j+3] == jeton {
				return true
			}
		}
	}
	return false
}

func Victoire(tableau [][]string, jeton string) bool {
	return VictoireVerticale(tableau, jeton) || VictoireHorizontale(tableau, jeton) || VictoireDiagonaleD(tableau, jeton) || VictoireDiagonaleA(tableau, jeton)
}

func TableauPlein(tableau [][]string) bool {
	if len(tableau) == 0 {
		return false
	}
	for j := 0; j < len(tableau[0]); j++ {
		if tableau[0][j] == "." {
			return false
		}
	}
	return true
}

var difficulteActuelle string = "difficulte1"

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		j1 := r.URL.Query().Get("joueur1")
		j2 := r.URL.Query().Get("joueur2")
		difficulte := r.URL.Query().Get("difficulte")

		if j1 != "" && j2 != "" {
			joueur1 = j1
			joueur2 = j2
			JoueurActuel = "X"
			gagnant = ""

			if difficulte != "" {
				difficulteActuelle = difficulte
			}

			var rows, cols int
			switch difficulteActuelle {
			case "difficulte1":
				rows, cols = 6, 7
			case "difficulte2":
				rows, cols = 6, 9
			case "difficulte3":
				rows, cols = 7, 8
			default:
				rows, cols = 6, 7
			}
			TableauJeu(rows, cols)
		}
	}

	if r.Method == http.MethodPost && gagnant == "" {
		r.ParseForm()
		col, _ := strconv.Atoi(r.FormValue("colonne"))

		JetonPlacer(tableau, col, JoueurActuel)

		if Victoire(tableau, JoueurActuel) {
			if JoueurActuel == "X" {
				gagnant = joueur1
			} else {
				gagnant = joueur2
			}
		} else if TableauPlein(tableau) {
			gagnant = "egalite"
		} else {
			if JoueurActuel == "X" {
				JoueurActuel = "O"
			} else {
				JoueurActuel = "X"
			}
		}
	}

	tmpl, err := template.New("play.html").Funcs(template.FuncMap{
		"eq": func(a, b string) bool { return a == b },
	}).ParseFiles("./pages/play.html")
	if err != nil {
		log.Fatal(err)
	}
	var joueurAffiche string
	if JoueurActuel == "X" {
		joueurAffiche = joueur1
	} else {
		joueurAffiche = joueur2
	}

	data := struct {
		Tableau [][]string
		Joueur  string
		Gagnant string
		Joueur1 string
		Joueur2 string
	}{Tableau: tableau, Joueur: joueurAffiche, Gagnant: gagnant, Joueur1: joueur1, Joueur2: joueur2}

	tmpl.Execute(w, data)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	gagnant = ""
	JoueurActuel = joueur1

	var rows, cols int

	switch difficulteActuelle {
	case "difficulte1":
		rows, cols = 6, 7
	case "difficulte2":
		rows, cols = 6, 9
	case "difficulte3":
		rows, cols = 7, 8
	default:
		rows, cols = 6, 7
	}

	TableauJeu(rows, cols)

	http.Redirect(w, r, "/play?joueur1="+joueur1+"&joueur2="+joueur2, http.StatusSeeOther)
}
