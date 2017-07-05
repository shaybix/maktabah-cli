// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"net/http"

	"io/ioutil"

	"io"

	"github.com/spf13/cobra"
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Download a book by it's ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 {
			cmd.Println(cmd.UsageString())
			return
		} else if id != 0 && all == true {
			cmd.Println(cmd.UsageString())
			return
		}

		if err := getBook(id); err != nil {
			log.Println(err)
			return
		}

		return
	},
}

var (
	fileType string
	all      bool
	id       int
	dir      string
)

func init() {
	downloadCmd.AddCommand(bookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bookCmd.Flags().StringVar(&fileType, "type", "rar", "Give filetype to download: rar | epub | pdf")
	bookCmd.Flags().BoolVar(&all, "all", false, "Download all books")
	bookCmd.Flags().IntVar(&id, "id", 0, "Id of book to download")
	bookCmd.Flags().StringVar(&dir, "dir", "books", "directory name where in to store the book(s)")
}

// Book represents the book being downloaded
type Book struct {
	Name      string
	URL       []byte
	Size      int
	Type      string
	Directory string
}

// getBook gets a book by the ID provided and if book found it is saved otherwise
// returns an error value
func getBook(id int) error {

	pageURL := "http://shamela.ws/index.php/book/" + strconv.Itoa(id)

	pageBody, err := getPage(pageURL)
	if err != nil {
		return err
	}
	url, err := findRARLink(pageBody)
	if err != nil {
		return err
	}

	file, err := downloadBook(url)
	if err != nil {
		return err
	}

	if err = saveBook(string(id), file); err != nil {
		return err
	}

	return nil
}

// getPage gets the page of the url specified and returns the body as
// io.ReadCloser and error
func getPage(url string) ([]byte, error) {

	var (
		err error
		p   []byte
	)

	resp, err := http.Get(url)
	if err != nil {
		return p, err
	}

	p, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}
	resp.Body.Close()

	return p, err
}

// findRARLink looks up in ReadCloser whether a rar link is present and if so
// returns the link as slice of byte, with error set to nil, otherwise returns error
func findRARLink(p []byte) ([]byte, error) {

	var rarURL []byte

	// It contains the RAR url link and therefore must extract it and return it.
	re := regexp.MustCompile(`http://shamela.ws/books/\d+/\d+.rar`)
	rarURL = re.Find(p)
	if rarURL == nil {
		return rarURL, errors.New("no rar link found")
	}

	return rarURL, nil
}

// downloadBook downloads the book
func downloadBook(url []byte) ([]byte, error) {
	var f []byte

	return f, nil
}

// saveBook saves downloaded book to the given directory specified in the flag
func saveBook(name string, p []byte) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0700); err != nil {
			return err
		}
	}

	file, err := os.Create(filepath.Join(dir, name+fileType))
	if err != nil {
		return err
	}

	rdr := bytes.NewReader(p)

	_, err := io.Copy(file, rdr)
	if err != nil {
		return err
	}

	return nil

}
