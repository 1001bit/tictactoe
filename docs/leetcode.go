package docs

func Tictactoe(moves [][]int) string {
	rows := [3][2]int{{0, 0}, {0, 0}, {0, 0}}
	cols := [3][2]int{{0, 0}, {0, 0}, {0, 0}}
	diags := [2][2]int{{0, 0}, {0, 0}}

	who := 0
	for i, m := range moves {
		rows[m[0]][who] += 1
		cols[m[1]][who] += 1
		if m[0] == m[1] {
			diags[0][who] += 1
		}
		if m[0] == 2-m[1] {
			diags[1][who] += 1
		}

		if rows[m[0]][who] == 3 || cols[m[1]][who] == 3 || diags[0][who] == 3 || diags[1][who] == 3 {
			if who == 0 {
				return "A"
			} else {
				return "B"
			}
		}

		if who == 0 {
			who = 1
		} else {
			who = 0
		}

		if i == 8 {
			return "Draw"
		}
	}
	return "Pending"
}
