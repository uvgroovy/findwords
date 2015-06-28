# Find Words
This utility helps you to find words that are made from a certain set of letters.

Useful when you are stuck while playing "words with friends"

This code was written as a go exercise. feedback is welcome.

# How to run:
## command line interface (go run)
```shell
go run findwords.go
```
for this to work you need to have a words file in
/usr/share/dict/words. if you don't have one, use the *-source* command line flag to specify a different words file.

## web interface (Docker)
```shell
docker build -t findwords .
docker run -p 8080:8080 findwords
```
then just set your broswer to http://localhost:8080/
## With a different words file
use the FINDWORDS_SOURCE environment variable to configure a different source for words.
```shell
docker run -e "FINDWORDS_SOURCE=http://..." -p 8080:8080 findwords
```
Or
```shell
docker run -p 8080:8080 findwords -source "http://..."
```

The file format is a list of words, one word per line.
