# Advent of Code 2024

This repo builds a CLI app for solving advent of code problems to earn stars.  Days and stars are added as subcommands, allowing you to easily add new star solvers to the CLI without touching the core command logic.

Follow these steps to add a new day (view `day1` as an example):

1. Add a new package for your day under the `cmd` directory.
2. Define a new `*cmd.Day` object and initialize it by using the `cmd.NewDay` helper function and then calling the `Init` method of the day object.
3. Create functions that match the signature `func xyz(inputFile string)` for solving each star.
4. Import your new day package in `main.go` using a blank import.

With that complete, you should be able to execute your new star solver with the following commands (substitue the day name, which star you are wanting to solve for, and the input file location):

```
> go build .
> aoc-2024 day1 star1 --if cmd/day1/input.txt
```