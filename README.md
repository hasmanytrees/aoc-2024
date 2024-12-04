# Advent of Code 2024

## Structure
This repo is a CLI app for solving advent of code problems to earn stars.  Days and stars are added as subcommands, allowing you to easily add new star solvers to the CLI without touching the core command logic.

To create a new day, execute `newday.sh` passing in the name of the day you wish to create:

```
❯ ./newday.sh day1
```

## Execution
To execute your new star solver build and run the CLI specifying the day name, star name and input file location:

```
❯ go build
❯ aoc-2024 day1 star1 --if cmd/day1/input.txt
```