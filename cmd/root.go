package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-2024",
	Short: "Command for processing Advent of Code 2024 Solutions",
}

var inputFile *string

func Execute() {
	inputFile = rootCmd.PersistentFlags().String("if", "input.txt", "Input file to be processed")
	rootCmd.MarkFlagRequired("if")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewStarCommand(name string, star func(string)) *cobra.Command {
	var starCmd = &cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			star(*inputFile)
		},
	}

	return starCmd
}
