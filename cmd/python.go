/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
