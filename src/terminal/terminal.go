package terminal

import(
	"fmt"
	"os"
    "os/exec"
    "runtime"
)

func Won(){
	Clear()
	fmt.Print(" **    **  ********  **    **     **            **  **  **    **\n")
	fmt.Print("  **  **   ********  **    **      **          **   **  ****  **\n")
	fmt.Print("   ****    **    **  **    **       **        **    **  ** ** **\n")
	fmt.Print("    **     **    **  **    **        **  **  **     **  **  ****\n")
	fmt.Print("    **     ********  ********         ********      **  **   ***\n")
	fmt.Print("    **     ********  ********          **  **       **  **    **\n")
}

func Lose(){
	Clear()
	fmt.Print(" **    **  ********  **    **     **        ********    ******  ********\n")
	fmt.Print("  **  **   ********  **    **     **        ********  ********  ********\n")
	fmt.Print("   ****    **    **  **    **     **        **    **  ****      ** **   \n")
	fmt.Print("    **     **    **  **    **     **        **    **      ****  ** **   \n")
	fmt.Print("    **     ********  ********     ********  ********  ********  ********\n")
	fmt.Print("    **     ********  ********     ********  ********  ******    ********\n")
}

func PrintBoard(Matrix [4][4]int) {
	Clear()
	for i := 0; i < 4; i++{
		fmt.Printf("-----------------------------\n")
		for j:= 0; j < 4; j++{
			fmt.Print("|")
			n := Matrix[i][j]
			if n>0 && n<10 {
				fmt.Printf("  %d   ",n)
			} else if n>10 && n<100 {
				fmt.Printf("  %d  ",n)
			} else if n>100 && n<1000 {
				fmt.Printf(" %d  ",n)
			} else if n > 1000{
				fmt.Printf(" %d ",n)
			}else{
				fmt.Printf("      ")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Printf("-----------------------------\n")
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cls") //Windows example it is untested, but I think its working 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func Clear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}