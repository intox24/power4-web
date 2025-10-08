package jeu

func TableauJeu(rows, cols int) [][]string {
	tableau := make([][]string, rows)
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
	col := len(tableau[0])
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
	cols := len(tableau[0])
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
	cols := len(tableau[0])
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
	cols := len(tableau[0])
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
	cols := len(tableau[0])
	for i := 3; i <= rows; i++ {
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
