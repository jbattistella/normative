/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/jbattistella/normative/engine"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	showCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func todayEvents() {

	ep := engine.NewEventParams()

	ep.Region = Region
	ep.Impact = append(ep.Impact, Impact)

	ev, err := ep.GetEvents()
	if err != nil {
		log.Println(err)
	}

	for _, v := range ev {
		date := v.Date.String()[0:10]
		time := v.Time.String()[12:19]
		fmt.Printf("%s | %s | %s | %s \n", date, time, v.Impact, v.Name)
	}

}
