package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var hecServer string
var hecToken string
var sourceType string
var hecURL string
var port string
var index string
var source string
var skip int
var verbose string

func main() {
	flag.StringVar(&hecServer, "h", "127.0.0.1", `host IP, Example -h="192.168.0.33"`)
	flag.StringVar(&hecToken, "t", "lilbigdata", `HEC Token, Example -t="7129b26a-c177-4705-aa5d-0eavf3b09cdf"`)
	flag.StringVar(&sourceType, "st", "mahdata", `Manually sets sourcetype, Example -c="hax"`)
	flag.StringVar(&port, "p", "8088", `Sets port, Example -p="443"`)
	flag.StringVar(&index, "i", "main", `Sets index, Example -p="main"`)
	flag.StringVar(&source, "s", "lilbigdata", `Sets source, Example -p="hax"`)
	flag.IntVar(&skip, "skip", 0, `Skips specified lines, Example -skip=4`)
	flag.StringVar(&verbose, "v", "false", `Turns on verbose mode, Example -v="true"`)
	flag.Parse()
	hecURL = fmt.Sprintf("https://%s:%s/services/collector", hecServer, port)

	/* //unable to get this to work with multiple lines. ping google.com for example triggers this.
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'mirp' | datas")
		return
	}
	*/

	reader := bufio.NewReader(os.Stdin)
	var output []rune
	var r = '\r'
	var n = '\n'
	var text string
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			text = string(output)
			hecSend(text)
			break
		}
		if input == r || input == n {
			text = string(output)
			if skip != 0 {
				if verbose != "false" {
					log.Println("skipped line ", skip)
				}
				skip = skip - 1
				input = ' '
				output = nil
			} else {
				hecSend(text)
				input = ' '
				output = nil
			}
		}
		output = append(output, input)
	}
}

func hecSend(text string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if text == " " {
		return
	}
	format := fmt.Sprintf(`{"sourcetype" : "%s", "source" : "%s", "index" : "%s", "event" : "%s"}`, sourceType, source, index, text)
	if verbose != "false" {
		log.Println("sending ", format)
	}
	body := strings.NewReader(format)
	req, err := http.NewRequest("POST", hecURL, body)
	_ = checkErr(err)
	token := fmt.Sprintf("Splunk %s", hecToken)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if verbose != "false" {
		log.Println(string(bodyBytes))
	}
}

func checkErr(err error) string {
	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
	return ""
}
