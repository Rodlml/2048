package functions

import(
    "math/rand"
    "time"
)

func Random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

func CopyMatrix(Matrix [4][4]int) [4][4]int{
	var Copy [4][4]int
	for i:=0; i<4; i++{
		for j:=0; j<4; j++{
			Copy[i][j] = Matrix[i][j]
		}
	}
	return Copy
}

func CheckIfWon(Matrix [4][4]int, n int) bool{
	b := false
	for i:=0; i<4; i++{
		for j:=0; j<4; j++{
			if Matrix[i][j] >= n{
				b = true
			}
		}
	}
	return b
}
