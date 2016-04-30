package main

import(
	"fmt"
	"flag"
	"github.com/Daniel-M/preuler/id1"
    "github.com/Daniel-M/preuler/id2"
    "github.com/Daniel-M/preuler/id3"
    "github.com/Daniel-M/preuler/id4"
    "github.com/Daniel-M/preuler/id5"
    "github.com/Daniel-M/preuler/id6"
    "github.com/Daniel-M/preuler/id7"
    "github.com/Daniel-M/preuler/id8"
    "github.com/Daniel-M/preuler/id9"
//	"github.com/Daniel-M/preuler/id10"
)

func main(){

	fmt.Println("Project Euler solutions in Go")
	fmt.Println("(c) 2016 - Daniel Mejía R. (danielmejia55@gmail.com)")

	commandLineArguments := flag.Int("id",0,"Id of the problem to be solved")

	flag.Parse()

	var idToSolve int = *commandLineArguments

	switch idToSolve {
	
	default:
		fmt.Println("Problem",idToSolve,"Unsolved yet!")

	case 0:
		flag.PrintDefaults()

	case 1:
		id1.Id1()
	case 2:
		id2.Id2()
	case 3:
		id3.Id3()
	case 4:
		id4.Id4()
	case 5:
		id5.Id5()
	case 6:
		id6.Id6()
	case 7:
		id7.Id7()
	case 8:
		id8.Id8()
	case 9:
		id9.Id9()
/*	case 10:
		id10.Id10()*/
	}
	
}


