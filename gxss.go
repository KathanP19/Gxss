package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"
)

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}

var httpClient = &http.Client{
	Transport: transport,
}

func main() {

	var c int
	var p string
	flag.IntVar(&c, "c", 50, "Set the Concurrency (Default 50)")
	flag.StringVar(&p, "p", "gxss", "Payload you want to Send to Check Refelection")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			testxss(p)
			wg.Done()
		}()
		wg.Wait()
	}
}

func testxss(p string) {

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	time.Sleep(500 * time.Microsecond)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		link := scanner.Text()
		u, _ := url.Parse(link)

		q, _ := url.ParseQuery(u.RawQuery)
		for key, value := range q {
			var tm string = value[0]
			q.Set(key, p)
			u.RawQuery = q.Encode()
			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				return
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				return
			}

			bodyBuffer, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			bodyString := string(bodyBuffer)

			re := regexp.MustCompile(p)
			match := re.FindStringSubmatch(bodyString)
			if match != nil {
				fmt.Printf("URL: %q\n", u)
				fmt.Printf("Reflected Param : %q\n", key)
			}
			q.Set(key, tm)
		}
	}
}
