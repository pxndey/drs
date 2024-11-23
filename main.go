package main

import (
	"flag"
	"fmt"
	"time"

	"example.com/drs/internal/helpers"
)

func main() {
	yearFlag := flag.Int("year", time.Now().Year(), "race year")
	raceFlag := flag.Int("race", 0, "race data flag")
	driversChampionshipFlag := flag.Bool("drivers", false, "drvier's championship for given year")
	constructorsChampionshipFlag := flag.Bool("constructors", false, "constructor's championship for given year")
	helpFlag := flag.Bool("help", false, "show help")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Welcome to DRS!\nUsage:")
		helpers.Help() // Call Help function to show the usage
		return
	}

	if *helpFlag {
		helpers.Help()
		return
	}

	if *raceFlag != 0 {
		helpers.Race(*yearFlag, *raceFlag)
	}
	if *driversChampionshipFlag {
		helpers.DriverStandings(*yearFlag)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	if *constructorsChampionshipFlag {
		helpers.ConstructorStandings(*yearFlag)
		fmt.Printf("\n")
	}
}
