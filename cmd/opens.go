/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	db "github.com/jbattistella/normative/db/sqlc"
	"github.com/spf13/cobra"
)

// opensCmd represents the opens command
var opensCmd = &cobra.Command{
	Use:   "opens",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateOpenPrices()
	},
}

func init() {
	updateCmd.AddCommand(opensCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opensCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opensCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func updateOpenPrices() {
	err := db.UpdateOpenPrices()
	if err != nil {
		log.Println(err)
	}
}
