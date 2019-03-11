// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var tag_opened string
var page_opened string

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if tag_opened == "" && page_opened == "" {
			fmt.Println("tag or page flag is required")
			os.Exit(1)
		}

		fmt.Println("open called")
		var file *os.File
		home, err := homedir.Dir()

		filepath := home + "/.ohp"

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		file, err = os.OpenFile(filepath, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()

		reader := bufio.NewReaderSize(file, 4096)

		var count_opened_page = 0

		for {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			data := strings.Split(string(line), ",")
			fmt.Println(data)
			page := data[0]
			tag := data[1]
			url := data[2]

			if page == page_opened || tag == tag_opened {
				browser.OpenURL(url)
				count_opened_page++
			}

			fmt.Printf("%d pages opened", count_opened_page)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(&tag_opened, "tag", "t", "", "page tag you want to see")
	openCmd.Flags().StringVarP(&page_opened, "page", "p", "", "page name you want to see")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
