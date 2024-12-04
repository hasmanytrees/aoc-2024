package cmd

import "github.com/spf13/cobra"

type Day struct {
	name  string
	star1 func(string)
	star2 func(string)
}

func NewDay(name string, star1 func(string), star2 func(string)) *Day {
	return &Day{
		name,
		star1,
		star2,
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

func (d *Day) Init() {
	var dayCmd = &cobra.Command{
		Use: d.name,
	}

	rootCmd.AddCommand(dayCmd)

	dayCmd.AddCommand(NewStarCommand("star1", d.star1))
	dayCmd.AddCommand(NewStarCommand("star2", d.star2))
}
