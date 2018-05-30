// Copyright © 2017 Abdisamad Hashi <shaybix@tuta.io>
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
	"fmt"

	"github.com/spf13/cobra"
)

// authorCmd represents the author command
var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Download all books specific to an author.",
	Run: func(cmd *cobra.Command, args []string) {
		author := Author{}

		if err := author.run(); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	downloadCmd.AddCommand(authorCmd)
}

// Author represents the author of books on shamela
type Author struct {
	name  string
	died  string
	books []Book
}

// run will execute the process of scraping the website
// for authors
func (a *Author) run() error {
	return nil
}
