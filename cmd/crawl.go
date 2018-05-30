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
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "crawls through the shamela website for links",
	Long: `Crawl takes a subcommand that specifies what links to gather.
	You can specify books, authors, or categories. The shamela website then
	will be concurrently crawled for those links.`,
}

func init() {
	RootCmd.AddCommand(crawlCmd)
}

var (
	quit           chan int
	shamelaBaseURL = "http://www.shamela.ws"
	shamelaMainURL = shamelaBaseURL + "/rep.php/main"
)

// Worker is dispatched in its own goroutine
type Worker struct {
	wg    sync.WaitGroup
	queue chan Job
	job   Job
}

func initWorker(size int, job Job) chan Job {
	queue := make(chan Job, 10)

	for i := 1; i <= size; i++ {
		w := Worker{queue: queue, job: job}
		w.run()
	}

	return queue
}

// run runs the worker
func (w *Worker) run() {
	w.queue = make(chan Job, 10)
	for {
		select {
		case job := <-w.queue:
			w.wg.Add(1)
			go job.scrape(&w.wg)
		case <-time.After(3 * time.Second):
			quit <- 1
		}
	}
}

// Job represents a job that gets
type Job struct {
	mu sync.Mutex
	// targetURL is the url to be fetched and scraped
	targetURL string
	// urls are the links found in the targetURL after
	// having been scraped.
	urls []string
}

// scrape goes through a webpage and collects every link to be found in the webpage.
func (j *Job) scrape(wg *sync.WaitGroup) []string {

	wg.Add(1)

	// instantiate the slice of string holding the urls
	var urls []string

	resp, err := http.Get(j.targetURL)
	if err != nil {
		return urls
	}

	// instantiate a new Tokenizer with the response body
	z := html.NewTokenizer(resp.Body)

ForLoop:
	for {
		// we are looping over the the characters in the response body
		textToken := z.Next()

		switch {
		// We are checking if the textToken is an ErrorToken
		// This checks if we have hit the end, and if so we break out
		// of the named for loop 'ForLoop'
		case textToken == html.ErrorToken:
			break ForLoop
		// we check if we have the start of a tagToken '<a>'
		case textToken == html.StartTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, a := range token.Attr {
					if a.Key == "href" {
						if strings.Contains(a.Val, "category") {
							urls = append(urls, shamelaBaseURL+a.Val)
						}
						break
					}
				}

			}
		}
	}

	resp.Body.Close()

	wg.Done()

	return urls
}
