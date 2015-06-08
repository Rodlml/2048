// Rodrigo Marroquín
// l,,l,

package main

import (
	"terminal"
	"bufio"
	"os"
	"fmt"
	"functions"
)

 var WIN int
var Matrix [4][4]int

func main() {
	WIN := 2048
	reader := bufio.NewReader(os.Stdin)

	//initialize matrix
	for i := 0; i < 4; i++{
		for j:= 0; j < 4; j++{
			Matrix[i][j] = 0
		}
	}

	insertRandom()
	for{ // while true
		somethingMoved := false
		terminal.PrintBoard(Matrix)
		fmt.Printf("\nEnter your move: ")
		move, _ := reader.ReadString('\n')

		if move == "a\n" || move == "A\n"{
			somethingMoved = moveLeft()
		} else if move == "s\n" || move == "S\n"{
			somethingMoved = moveDown()
		} else if move == "d\n" || move == "D\n"{
			somethingMoved = moveRight()
		} else if move == "w\n" || move == "W\n"{
			somethingMoved = moveUp()
		} else {
			fmt.Printf("invalid")
		}
		// CHEK IF WON
		if functions.CheckIfWon(Matrix, WIN){
			won(true)
		}

		//CHECK IF LOSE
		backup := functions.CopyMatrix(Matrix)
		bl := moveLeft()
		Matrix = functions.CopyMatrix(backup)
		br := moveRight()
		Matrix = functions.CopyMatrix(backup)
		bu := moveUp()
		Matrix = functions.CopyMatrix(backup)
		bd := moveDown()
		Matrix = functions.CopyMatrix(backup)

		if !bl && !br && !bu && !bd{
			won(false)
		} 

		//Wheter to insert or not
		if somethingMoved{
			insertRandom()
		}
	}
}

func won(w bool){
	reader := bufio.NewReader(os.Stdin)
	if w {
		terminal.Won()
	}else{
		terminal.Lose()
	}
	fmt.Printf("\nPlay Again? (Y/N) ")
	choice, _ := reader.ReadString('\n')
	if choice == "Y\n" || choice == "y\n" {
		terminal.Clear()
		main()
	}else if choice == "N\n" || choice == "n\n" {
		terminal.Clear()
		os.Exit(0)
	}else{
		won(w)
	}
}

func moveLeft() bool{
	b := false
	for i:=0; i<4; i++{
		//Move Everything to the Left
		var aux [4]int
		var aux2 [4]int
		for j,n := 0,0; j<4; j++{
			if Matrix[i][j] != 0{
				aux[n] = Matrix[i][j]
				if j != n{
					b = true
				}
				n++
			}
		}
		//Sum where possible
		for j,k := 0,0; j<3; j++{
			if aux[j] == aux[j+1]{
				aux2[k] = aux[j]*2
				j++
				if aux[j] != 0{
					b = true
				}
			}else{
				aux2[k] = aux[j]
			}
			k++
			if j<3{
				aux2[k] = aux[3]
			}
		}
		//Copy to board
		for j:=0; j<4; j++{
			Matrix[i][j] = aux2[j]
		}
	}
	return b
}


func moveRight() bool{
	b := false
	for i:=0; i<4; i++{
		//Move Everything to the Left
		var aux [4]int
		var aux2 [4]int
		for j,n := 3,3; j>=0; j--{
			if Matrix[i][j] != 0{
				aux[n] = Matrix[i][j]
				if j != n{
					b = true
				}
				n--
			}
		}
		//Sum where possible
		for j,k := 3,3; j>0; j--{
			if aux[j] == aux[j-1]{
				aux2[k] = aux[j]*2
				j--
				if aux[j] != 0{
					b = true
				}
			}else{
				aux2[k] = aux[j]
			}
			k--
			if j>0{
				aux2[k] = aux[0]
			}
		}
		//Copy to board
		for j:=0; j<4; j++{
			Matrix[i][j] = aux2[j]
		}
	}
	return b
}

func moveUp() bool{
	b := false
	for i:=0; i<4; i++{
		//Move Everything to the Left
		var aux [4]int
		var aux2 [4]int
		for j,n := 0,0; j<4; j++{
			if Matrix[j][i] != 0{
				aux[n] = Matrix[j][i]
				if j != n{
					b = true
				}
				n++
			}
		}
		//Sum where possible
		for j,k := 0,0; j<3; j++{
			if aux[j] == aux[j+1]{
				aux2[k] = aux[j]*2
				j++
				if aux[j] != 0{
					b = true
				}
			}else{
				aux2[k] = aux[j]
			}
			k++
			if j<3{
				aux2[k] = aux[3]
			}
		}
		//Copy to board
		for j:=0; j<4; j++{
			Matrix[j][i] = aux2[j]
		}
	}
	return b
}

func moveDown() bool{
	b := false
	for i:=0; i<4; i++{
		//Move Everything to the Left
		var aux [4]int
		var aux2 [4]int
		for j,n := 3,3; j>=0; j--{
			if Matrix[j][i] != 0{
				aux[n] = Matrix[j][i]
				if j != n{
					b = true
				}
				n--
			}
		}
		//Sum where possible
		for j,k := 3,3; j>0; j--{
			if aux[j] == aux[j-1]{
				aux2[k] = aux[j]*2
				j--
				if aux[j] != 0{
					b = true
				}
			}else{
				aux2[k] = aux[j]
			}
			k--
			if j>0{
				aux2[k] = aux[0]
			}
		}
		//Copy to board
		for j:=0; j<4; j++{
			Matrix[j][i] = aux2[j]
		}
	}
	return b
}

func insertRandom(){
	rand1 := functions.Random(0, 4)
	rand2 := functions.Random(0, 4)
	for Matrix[rand1][rand2] != 0{
		rand1 = functions.Random(0, 4)
		rand2 = functions.Random(0, 4)
	}
	randn := functions.Random(1,3)*2
	Matrix[rand1][rand2] = randn
}