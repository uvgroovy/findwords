package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"github.com/uvgroovy/findwords/wordmap"
	"github.com/uvgroovy/findwords/server"
	"net/http"
	"flag"
)

const (
	DEFAULT_SOURCE = "/usr/share/dict/words"
	SOURCE_ENVVAR_NAME = "FINDWORDS_SOURCE"
)

func getWords(source string) io.ReadCloser {
	
	if source == "" {
		source = os.Getenv(SOURCE_ENVVAR_NAME)
	}
	
	if source == "" {
		source = DEFAULT_SOURCE
	}
	
	if (strings.HasPrefix(source, "http")) {
		resp, err := http.Get(source)
		if err != nil {
			panic("Can't load url " + source)
		}
		return resp.Body

	} else {
		file, err := os.Open(source)
		if err != nil {
			panic("no dict file")
		}
		
		return file	
	}
	
	panic("No source for words found")
	
}

func Serve(words wordmap.WordsMap) {
	
	wh := &server.WordsHandler{words}
	
	mux := http.NewServeMux()
	
	mux.Handle("/words", wh)
	mux.Handle("/",  http.FileServer(http.Dir("./html")))
	
	log.Fatal(http.ListenAndServe(":8080", mux))

}


func CreateWordsMap(source string) wordmap.WordsMap {
	words  := wordmap.WordsMap{}
	reader := getWords(source)
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			words.AddWord(line)
		}
	}

	return words
}

func commandLine(words wordmap.WordsMap) {
	
    fmt.Println("Type q to quit.")
	for {
		fmt.Print("Enter letters: ")
		var letters string
		if _, err := fmt.Scanln(&letters); err != nil {
			if err != io.EOF {
				log.Fatalln(err)	
			}
			return
		}

		if letters == "q" {
			return
		}


		letters = strings.TrimSpace(letters)

		for word := range words.GetWords(letters) {
			
				fmt.Println(word)
		}
	}
}

func main() {
	
	server := flag.Bool("server", false, "if set to true, runs in server mode. otherwise runs in commandline mode")
	source := flag.String("source", "", "Where to get the word list from. can also be set using the FINDWORDS_SOURCE env var.")
	
	flag.Parse()
	
	words := CreateWordsMap(*source)
	if *server {
		Serve(words)	
	} else {
		commandLine(words)
	}
	
}
