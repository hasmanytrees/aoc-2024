# Advent of Code 2024

## Structure
This repo is a CLI app for solving advent of code problems to earn stars.  Days and stars are added as subcommands, allowing you to easily add new star solvers to the CLI without touching the core command logic.

To create a new day, execute `newday.sh` passing in the name of the day you wish to create:

```
❯ ./newday.sh day1
```

The subcommand for the newly created day will be located under the `cmd` folder.

## Execution
To execute your new star solver, build and run the CLI specifying the day name subcommand, star name subcommand, and input file param (`--if` which will be sourced from the day subcommand folder):

```
❯ go build
❯ ./aoc-2024 day1 star1 --if input.txt
```