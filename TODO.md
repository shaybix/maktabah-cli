Below are a list of features and tasks to be either implemented in the near future or in the process of being implemented.



##### Fetching Shamela resources from the website 

###### Crawling the website to gather urls
- [ ] Crawl through the website to gather urls for Author pages
- [ ] Crawl through the website to gather urls of books by Category ID
- [ ] Crawl through the website to gather urls of books by Author ID

###### Abillity to store urls in a cached file in a number of formats
- [ ] Store urls with author information in a json format
- [ ] Store urls only in a plain text file with url per newline


###### Download resources 
- [ ] Download BOK file in a directory
- [ ] Download PDF file in a directory
- [ ] Download Epub file in a directory



#### Linting issues to investigate
```
cmd/crawl.go
initWorker passes lock by value: cmd.Job contains sync.Mutex(53,1)
literal copies lock value from job: cmd.Job contains sync.Mutex(57,1)
assignment copies lock value to job: cmd.Job contains sync.Mutex(69,1)
```