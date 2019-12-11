// Package CMD provides CLI utility
package cmd

import (
	// Package fmt allows formating options (ex: print line)
	"fmt"
	// Package strings allows string manipulation (ex: join)
	// "strings"

	// Package cobra facilitates CLI projects
	"github.com/spf13/cobra"
	// Package colly permits web scrapping
	"github.com/gocolly/colly"
)

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "Get Golang Syntax",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// All CLI functionality is nested under the RUN section
	// All submitted arguments can be found in args slice
	Run: func(cmd *cobra.Command, args []string) {
		// c is a new instance of the colly collector
		c := colly.NewCollector(
		// Put allowed domain parameters here
		)
		// Case - Single Argument: Print Index Of All Package Methods
		if len(args) == 1 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://golang.org/pkg/%s/", args[0])
			// Selector is the type of html we want to scrape
			selector := "#manual-nav > dl > dd > a"

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				if e.Text == "" {
					fmt.Println("No package in golang docs with that name, please check spelling")
				} else {
					fmt.Println(e.Text)
				}


			})

			c.Visit(link)
		}

		// Case - Two Arguments: Print Documentation For Single Package Method
		if len(args) == 2 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://golang.org/pkg/%s/#%s", args[0], args[1])
			// Selector is the type of html we want to scrape
			selector := fmt.Sprintf("h2[id=%s]+pre", args[1])

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				if e.Text == "" {
					fmt.Println("No package/function combo in golang docs with that name, please check spelling")
				} else {
					fmt.Println("Function syntax:")
					fmt.Println(e.Text)
					fmt.Printf("For usage example: \n$ ./syntax-cli golang %s %s example \n",
						args[0], args[1])
				}
			})
			c.Visit(link)
		}

		if len(args) == 3 && args[2] == "example" {
			link := fmt.Sprintf("https://golang.org/pkg/%s/#%s", args[0], args[1])

			selector := fmt.Sprintf("#example_%s > div.expanded > div > div.input > textarea", args[1])

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				if e.Text == "" {
					fmt.Println("No example in golang docs available")
				} else {
					fmt.Println("Usage example:")
					fmt.Println(e.Text)
				}
			})
			c.Visit(link)
		}

	},
}

func init() {
	rootCmd.AddCommand(golangCmd)
}
