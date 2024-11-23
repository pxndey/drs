package helpers

import (
	"fmt"
)

func Help() {
	fmt.Printf("  -year <year>: Specify the race year (default is current year).")
	fmt.Printf("  -race <race number>: Specify the race data flag.")
	fmt.Printf("  --drivers: Display driver's championship standings for the given year.")
	fmt.Printf("  --constructors: Display constructor's championship standings for the given year.")
}
