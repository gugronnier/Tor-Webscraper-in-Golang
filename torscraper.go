package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"io"
	"bufio"
	"strings"

	// import external libraries
	"golang.org/x/net/proxy"
	"github.com/PuerkitoBio/goquery"
)

// ## GLOBAL VARIABLES ##
// Declare global variable
var outputFile *os.File
var torAddr string
var outputPrefix string

// ## USAGE ##
// Usage function explain to user how program is working
func usage() {
	fmt.Println("Usage : ./tor_in_go http://exempleosgrodnsoeifjgojfr.onion/ outputFilenamePrefix ")
}


// ## PROCESSING
// if there is an error, exit of the program and print the error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Take file where write as parameter and string to write
// And write in the file without overwriting what is written in before
func writeInFile(input string) {
        outputFile.Sync()
        w := bufio.NewWriter(outputFile)
        _, err := w.WriteString(input + "\n")
        check(err)
        w.Flush()
}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
        // see if the href attribut exists on the element
        href, exists := element.Attr("href")
        if exists {
		if strings.HasPrefix(href, "http") {
                	writeInFile(href)
		}else {
			if !strings.HasPrefix(href, "bitcoin") {
				writeInFile(torAddr + strings.TrimPrefix(href, "/"))
			}
		}
        }
}

// Processing that allow you to extract all links available in a defined page
func getLinks(fname string) {
	// Open HTTP output file to use it
	httpFile, err := os.Open(fname)
	check(err)
	defer httpFile.Close()

        // Create a goquery document from HTTP response
        document, err := goquery.NewDocumentFromReader(httpFile)
        if err != nil {
                log.Fatal("Error loading HTTP response body", err)
        }

        // Create Links output file with crafted filename
        filename := outputPrefix + "_links.txt"
	outputFile, err = os.Create(filename) // write in global var
        check(err)
        defer outputFile.Close()

        _, err = outputFile.WriteString("All links for " + torAddr + "\n")
        check(err)

        // Find all links and process them with the function defined earlier
        document.Find("a").Each(processElement)

	// Close file after use
	outputFile.Close()
}

// [TODO] name file like `outputPrefix_id-[id].html` or `outputPrefix_pid-[pid].html` 
// and move them into directory named with `outputPrefix`


// [TODO] check if now created page already exist and is different
// if so we can overwrite it or store it under a subdirectory named
// like `outputPrefix_id-[id]` or `outputPrefix_pid-[pid]`
// and each new version end by `(n)` where 'n' is the version number


// ## MAIN FUNC ##
func main() {
	// targeted Onion Site must be enter as argument when calling the program
	args := os.Args[1:]
	if len(args) != 2 {
		usage()
		os.Exit(0)
	}
	torAddr = args[0]
	outputPrefix = args[1]

	// Create a socks5 dialer
	dialer, err := proxy.SOCKS5 ("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	check(err)

	// setup HTTP Transport
	tr := &http.Transport{
		Dial: dialer.Dial,
	}
	client := &http.Client{Transport: tr}

	// Create HTTP request and modify User-Agent
	request, err := http.NewRequest("GET", torAddr, nil)
	check(err)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; rv:59.0) Gecko/59.0 Firefox/59.0")

	// Make request
	response, err := client.Do(request)
        check(err)

	// Create HTML output file with crafted filename
	foutput := outputPrefix + "_output.html"
	outFile, err := os.Create(foutput)
	check(err)
	defer outFile.Close()

	// Copy data from HTTP response to ouput file
	_, err = io.Copy(outFile, response.Body)
	check(err)
	outFile.Close()

	// Call function for get all links available in a defined page
	getLinks(foutput)

	// Output Message for user
	fmt.Println("#####################################################################")
	fmt.Println("#    HTML document is available in `",outputPrefix,"_ouput.html`    #")
	fmt.Println("#    Links is available in `",outputPrefix,"_links.txt`             #")
	fmt.Println("#####################################################################")
}
