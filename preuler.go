package main

import (
	"flag"
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/id1"
	"github.com/Daniel-M/ProjectEuler_Solutions/id10"
	"github.com/Daniel-M/ProjectEuler_Solutions/id11"
	"github.com/Daniel-M/ProjectEuler_Solutions/id12"
	"github.com/Daniel-M/ProjectEuler_Solutions/id13"
	"github.com/Daniel-M/ProjectEuler_Solutions/id14"
	"github.com/Daniel-M/ProjectEuler_Solutions/id2"
	"github.com/Daniel-M/ProjectEuler_Solutions/id3"
	"github.com/Daniel-M/ProjectEuler_Solutions/id4"
	"github.com/Daniel-M/ProjectEuler_Solutions/id5"
	"github.com/Daniel-M/ProjectEuler_Solutions/id6"
	"github.com/Daniel-M/ProjectEuler_Solutions/id7"
	"github.com/Daniel-M/ProjectEuler_Solutions/id8"
	"github.com/Daniel-M/ProjectEuler_Solutions/id9"
	"os"
)

func main() {

	fmt.Println("Project Euler solutions in Go")
	fmt.Println("(c) 2016 - Daniel Mejía R. (danielmejia55@gmail.com)")

	commandLineArgumentID := flag.Int("id", 0, "Id of the projecteuler problem to be solved")
	commandLineArgumentShow := flag.Bool("solved", false, "Show solved problems so far")

	flag.Parse()

	var idToSolve int = *commandLineArgumentID
	var iSolved int = 14

	if *commandLineArgumentShow == true && idToSolve > iSolved {
		fmt.Println("Problems from 1 to id", iSolved, " solved.")
		os.Exit(0)
	} else if *commandLineArgumentShow == true && idToSolve < iSolved {
		fmt.Println("Problems from 1 to id", iSolved, " solved.")
	}

	switch idToSolve {

	default:
		if idToSolve < 0 {
			fmt.Println("Invalid problem ID", idToSolve)
			fmt.Println("Call --solved option to show the solved problems so far.")
		} else {
			fmt.Println("Problem", idToSolve, "Unsolved yet!")
			fmt.Println("Call --solved option to show the solved problems so far.")
		}

	case 0:
		flag.PrintDefaults()
	case 1:
		id1.Solution()
	case 2:
		id2.Solution()
	case 3:
		id3.Solution()
	case 4:
		id4.Solution()
	case 5:
		id5.Solution()
	case 6:
		id6.Solution()
	case 7:
		id7.Solution()
	case 8:
		id8.Solution()
	case 9:
		id9.Solution()
	case 10:
		id10.Solution()
	case 11:
		id11.Solution()
	case 12:
		id12.Solution()
	case 13:
		id13.Solution()
	case 14:
		id14.Solution()
	}

	os.Exit(0)
}
