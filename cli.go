package main

import (
	"fmt"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -csv Create excel csv ")
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {

	if len(os.Args) < 2 {
		fmt.Println("run default get Hacker news")
		hackernews.FetchHackerNews()
		os.Exit(1)
	} else {
		fmt.Println(os.Args[1])

		switch os.Args[1] {
		case "-csv":
			fmt.Println("run make csv")
			os.Exit(1)
		default:
			cli.printUsage()
			os.Exit(1)
		}
	}

}
