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

// CMD is a package that provides CLI utility
package cmd

import (
	// Fmt is a package that allows formating options
	"fmt"
	// Strings is a package that allows string manipulation
	"strings"

	// Cobra is a package that facilitates CLI projects
	"github.com/spf13/cobra"
	// Colly is a web scrapper
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
		// Content is the scrapped data we will be printing the command line 
		content := []string{}
		c := colly.NewCollector(
			// Put allowed domain parameters here
		)
		if len(args) == 1 {
			// Link is the website we will be scrapping
			link := fmt.Sprintf("https://golang.org/pkg/%s/", args[0])
			Selector is the
			selector := "#manual-nav > dl > dd > a"
			fmt.Println(selector)
			c.OnHTML(selector, func(e *colly.HTMLElement) {
				fmt.Println(e.Text)
				formatedText := fmt.Sprintf("%s", e.Text)
				content = append(content, formatedText)
				// fmt.Println("Got content")
				// fmt.Println(content)
				message := strings.Join(content, "")
				fmt.Println(message)
				// content += formatedText

			})

			// c.OnRequest(func(r *colly.Request) {
			// 	fmt.Printf("Visiting package: \"%s\"", args[0])
			// })

			c.Visit(link)
			//

		}

		// if len(args) == 2 {
		//
		// 	fmt.Println("Length is 2")
		// 	selector := fmt.Sprintf("h2[id=%s]+pre", args[1])
		//
		// 	link := fmt.Sprintf("From: https://golang.org/pkg/%s/#%s", args[0], args[1])
		// 	// fmt.Println(link)
		//
		// 	c.OnHTML(selector, func(e *colly.HTMLElement) {
		// 		fmt.Println(e.Text)
		// 		formatedText := fmt.Sprintf("`%s`", e.Text)
		// 		// c.Visit(e.Request.AbsoluteURL(link))
		// 		// fmt.Println(e.ChildText(selector))
		// 		content = append(content, formatedText)
		// 	})
		//
		// 	c.OnRequest(func(r *colly.Request) {
		// 		fmt.Printf("Visiting package: \"%s\" function:\"%s\" ", args[0], args[1])
		// 	})
		//
		// 	c.Visit(selector)
		// }

		// fmt.Println(content)

		// fmt.Printf("golang called, %+v number of arguments", numArgs)

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
