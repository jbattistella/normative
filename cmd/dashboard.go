/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"fmt"
	"log"
	"time"

	db "github.com/jbattistella/normative/db/sqlc"

	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dashboard()
	},
}

var Region string
var Impact1 string
var Impact2 string

func init() {
	rootCmd.AddCommand(dashboardCmd)

	flag.StringVar(&Region, "region", "United States", "a string var")

	flag.StringVar(&Impact1, "impact1", "high", "a string var")

	flag.StringVar(&Impact2, "impact2", "medium", "a string var")

	flag.Parse()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashboardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dashboard() {
	// ranges
	current, yr, mn, wk, yrC, mnC, wkC, err := db.GetPriceData()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Current Price: $%.01f\n", current)
	fmt.Printf("Period        | price change    | percent change  |\n")
	fmt.Printf("year-to-date  | %.01f           | %.01f%%          |\n", yr, yrC*100)
	fmt.Printf("month         | %.01f           | %.01f%%          |\n", mn, mnC*100)
	fmt.Printf("week          | %.01f            | %.01f%%           |\n", wk, wkC*100)
	fmt.Printf("\n")
	fmt.Printf("----------------------------------------------------------------\n")

	//events
	ep := db.NewEventParams()

	ep.Region = Region
	ep.Impact = append(ep.Impact, Impact1)
	ep.Impact = append(ep.Impact, Impact2)

	ev, err := ep.GetEvents()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("************ECONOMIC EVENTS****************")
	fmt.Printf("  date       |   time       |   name                     |  impact  |\n ")
	for _, v := range ev {
		date := v.Date.String()[0:10]
		time := v.Time.Add(time.Hour * -6).String()[12:19]

		fmt.Printf("%s  |  %s CT  |   %s   |   %s   |\n", date, time, v.Name, v.Impact)

	}

}
