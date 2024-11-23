package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

type constructorData struct {
	StandingsTable struct {
		StandingsList struct {
			ConstructorStandings []struct {
				Points      string `xml:"points,attr"`
				Constructor struct {
					Name string `xml:"Name"`
				} `xml:"Constructor"`
			} `xml:"ConstructorStanding"`
		} `xml:"StandingsList"`
	} `xml:"StandingsTable"`
}

type driverStruct struct {
	DriverStandings []struct {
		Points string `xml:"points,attr"`
		Driver struct {
			GivenName  string `xml:"GivenName"`
			FamilyName string `xml:"FamilyName"`
		} `xml:"Driver"`
		Team struct {
			Name string `xml:"Name"`
		} `xml:"Constructor"`
	} `xml:"StandingsTable>StandingsList>DriverStanding"`
}

// Greet prints a greeting message.

func getDriverStandings(year int) [][]string {
	driverURL := fmt.Sprintf("http://ergast.com/api/f1/%d/driverstandings", year)

	resp, err := http.Get(driverURL)
	if err != nil {
		fmt.Println("error fetching: ", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading resp: ", err)
	}

	var data driverStruct

	err = xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("error parsing resp: ", err)
	}

	var standings [][]string

	for _, standing := range data.DriverStandings {
		row := []string{
			fmt.Sprintf(standing.Driver.GivenName + " " + standing.Driver.FamilyName),
			standing.Team.Name,
			standing.Points,
		}
		standings = append(standings, row)
	}

	return standings
}

func getConstructorStandings(year int) [][]string {
	constructorURL := fmt.Sprintf("http://ergast.com/api/f1/%d/constructorStandings", year)

	resp, err := http.Get(constructorURL)
	if err != nil {
		fmt.Println("error fetching: ", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading resp: ", err)
	}

	var data constructorData

	err = xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("error parsing resp: ", err)
	}

	var standings [][]string

	for _, standing := range data.StandingsTable.StandingsList.ConstructorStandings {
		row := []string{
			fmt.Sprintf(standing.Constructor.Name),
			standing.Points,
		}
		standings = append(standings, row)
	}

	return standings
}

// TODO: format returned into table: input:year
func DriverStandings(year int) {
	data := getDriverStandings(year)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for i, row := range data {
		fmt.Fprintln(w, fmt.Sprintf("%d.", i+1)+"\t"+row[0]+"\t"+row[1]+"\t"+row[2])
	}
	w.Flush()
}

// TODO: constructor standings: input:year
func ConstructorStandings(year int) {
	// return fmt.Sprintf("constructr standings")
	data := getConstructorStandings(year)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for i, row := range data {
		fmt.Fprintln(w, fmt.Sprintf("%d.", i+1)+"\t"+row[0]+"\t"+row[1])
	}
	fmt.Println("")

	w.Flush()
}
