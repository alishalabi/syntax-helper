// Package cmd provides CLI utility
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

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "A series of commands to get Python code syntax",
	Run: func(cmd *cobra.Command, args []string) {
		// c is a new instance of the colly collector
		c := colly.NewCollector(
		// Put allowed domain parameters here
		)
		// Case - Single Argument: Print Index Of All Package Methods
		if len(args) == 1 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://docs.python.org/3/library/%s.html", args[0])
			// Selector is the type of html we want to scrape
			selector := "code.descname"

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				fmt.Println(e.Text)
			})
			c.Visit(link)
		}

		// Case - Two Arguments: Print Documentation For Single Package Method
		if len(args) == 2 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://docs.python.org/3/library/%s.html#%s.%s", args[0], args[0], args[1])
			selector := fmt.Sprintf("#%s-constants > dl > dd > p", args[0])
			fmt.Println(selector)

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				fmt.Println("Function Syntax:")
				fmt.Println(e.Text)
				// fmt.Printf("For Synatx example: \n$ ./syntax-cli python %s %s example",
				// 	args[0], args[1])

			})
			c.Visit(link)
		}
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
