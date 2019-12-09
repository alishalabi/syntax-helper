/*
Copyright © 2019 Ali Shalabi <ashalabi13@gmail.com>

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
				fmt.Println(e.Text)
				// formatedText := fmt.Sprintf("%s", e.Text)
				// content = append(content, formatedText)
				// fmt.Println("Got content")
				// fmt.Println(content)

				// content += formatedText

			})
			// message := strings.Join(content, "\n")
			// fmt.Println(message)
			c.Visit(link)
		}

		// Case - Two Arguments: Print Documentation For Single Package Method
		if len(args) == 2 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://golang.org/pkg/%s/#%s", args[0], args[1])
			// fmt.Println(link)
			// Selector is the type of html we want to scrape
			selector := fmt.Sprintf("h2[id=%s]+pre", args[1])

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				fmt.Println("Function Syntax:")
				fmt.Println(e.Text)
				fmt.Printf("For Synatx example: \n$ ./syntax-cli golang %s %s example",
					args[0], args[1])
				// formatedText := fmt.Sprintf("%s", e.Text)
				// content = append(content, formatedText)
				// fmt.Println("Got content")
				// fmt.Println(content)

				// content += formatedText

			})


			// message := strings.Join(content, "\n")
			// fmt.Println(message)
			c.Visit(link)
		}

		if len(args) == 3 && args[2] == "example" {
			link := fmt.Sprintf("https://golang.org/pkg/%s/#%s", args[0], args[1])

			selector := fmt.Sprintf("#example_%s > div.expanded > div > div.input > textarea", args[1])

			c.OnHTML(selector, func(e *colly.HTMLElement) {
				fmt.Println("Function Example:")
				fmt.Println(e.Text)
				// formatedText := fmt.Sprintf("%s", e.Text)
				// content = append(content, formatedText)
				// fmt.Println("Got content")
				// fmt.Println(content)

				// content += formatedText

			})
			c.Visit(link)
		}

	},
}

func init() {
	rootCmd.AddCommand(golangCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// golangCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
