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
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"

	"github.com/spf13/cobra"
)

// booksCmd represents the books command
var booksCmd = &cobra.Command{
	Use:   "books",
	Short: "To crawl for links specific to books",
	Run: func(cmd *cobra.Command, args []string) {
		books := Books{}
		if err := books.run(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	crawlCmd.AddCommand(booksCmd)

	booksCmd.Flags().StringVar(&dir, "dir", "books", "directory name where in to store the book(s)")
}

// Books ...
type Books struct {
	Books []Book
}

func newWorker(n int) <-chan string {

	c := make(chan string, n)

	return c

}

func (b *Books) run() error {

	Chan := newWorker(1)

ForLoop:
	for {
		select {
		case url := <-Chan:
			fmt.Println(url)
			continue
		case <-time.After(25 * time.Second):
			fmt.Println("oops!")
			break ForLoop
		}
	}
	return nil
}

// job is the actual job that scrapes a web page and sends
// relevant urls through the given channel c.
// job is also a recursive function that calls itself with each
// url that it finds.
func job(url string, wg *sync.WaitGroup, c chan string) {

	// scrape the url given and return the urls found.
	urls := scrape(url)

	// range over the urls given and for each url found
	// send off a job on a goroutine and send the url over
	// the given channel c.
	for _, url := range urls {
		wg.Add(1)
		go job(url, wg, c)
		fmt.Printf("Found URL:\t%s\n", url)
		c <- url
	}
	wg.Done()
	return
}

// scrape scrapes the reader given, looking for anchor tags
func scrape(url string) []string {
	var urls []string

	return urls

}

func getBooks(url string, wg *sync.WaitGroup, bookChan chan string) {

	wg.Add(1)

	resp, _ := http.Get(url)

	z := html.NewTokenizer(resp.Body)

ForLoop:
	for {
		textToken := z.Next()

		switch {
		case textToken == html.ErrorToken:
			break ForLoop
		case textToken == html.StartTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, a := range token.Attr {
					if a.Key == "href" {
						if strings.Contains(a.Val, "book") {
							bookChan <- shamelaBaseURL + a.Val
							// urls = append(urls, shamelaBaseURL+a.Val)
						}
						break
					}
				}

			}
		}
	}

	wg.Done()
	return
}
