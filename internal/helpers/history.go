package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

type RaceResult struct {
	RaceName string `xml:"RaceTable>Race>RaceName"`
	Result   []struct {
		Driver struct {
			GivenName  string `xml:"GivenName"`
			FamilyName string `xml:"FamilyName"`
		} `xml:"Driver"`
		Constructor struct {
			Name string `xml:"Name"`
		} `xml:"Constructor"`
	} `xml:"RaceTable>Race>ResultsList>Result"`
}

func getRace(year int, race int) [][]string {
	raceURL := fmt.Sprintf("https://ergast.com/api/f1/%d/%d/results", year, race)

	resp, err := http.Get(raceURL)
	if err != nil {
		fmt.Println("error fetching: ", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading resp: ", err)
	}

	var data RaceResult

	err = xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("error parsing: ", err)
	}

	fmt.Printf("%d %s\n", year, data.RaceName)
	var standings [][]string
	for _, standing := range data.Result {
		row := []string{
			fmt.Sprintf(standing.Driver.GivenName + " " + standing.Driver.FamilyName),
			standing.Constructor.Name,
		}
		standings = append(standings, row)
	}

	return standings
}

// Greet prints a greeting message.
func Race(year int, race int) {
	data := getRace(year, race)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for i, row := range data {
		fmt.Fprintln(w, fmt.Sprintf("%d.", i+1)+"\t"+row[0]+"\t"+row[1])
	}
	w.Flush()

}
