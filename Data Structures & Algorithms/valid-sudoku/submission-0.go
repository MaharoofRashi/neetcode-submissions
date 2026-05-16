func isValidSudoku(board [][]byte) bool {

	// find if all the row are valid, if not return false
	for i := 0; i < 9; i++ {
		m := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, exists := m[board[i][j]]; exists {
					return false
				}
				m[board[i][j]] = struct{}{}
			}
		}
	}


	// find if all the coloums are valid, if not return false
	for i := 0; i < 9; i++ {
		m := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if _, exists := m[board[j][i]]; exists {
					return false
				}
				m[board[j][i]] = struct{}{}
			}
		}
	}

	// find if 3*3 sub-boxes of the grid is valid if not return false

	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			m := make(map[byte]struct{})
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					cell := board[boxRow*3+i][boxCol*3+j]
					if cell != '.' {
						if _, exists := m[cell]; exists {
							return false
						}
						m[cell] = struct{}{}
					}
				}
			}
		}
	}

	// if all the previous didn't get hit return true
	return true

}
