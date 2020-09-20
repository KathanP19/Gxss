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

	"github.com/briandowns/spinner"
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
	var v bool
	flag.IntVar(&c, "c", 50, "Set the Concurrency (Default 50)")
	flag.StringVar(&p, "p", "", "Payload you want to Send to Check Refelection")
	flag.BoolVar(&v, "v", false, "Verbose mode")
	flag.Parse()

	if v == true {
		fmt.Println(`                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
	0.1 - @KathanP19
	`)
	}

	if p != "" {

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
		s.Suffix = " Testing XSS "
		s.Start() // Start the spinner
		time.Sleep(5 * time.Second)
		if v == true {
			s.Stop()
		}

		var wg sync.WaitGroup
		for i := 0; i < c; i++ {
			wg.Add(1)
			go func() {
				testxss(p, v)
				wg.Done()
			}()
			wg.Wait()
		}

	} else {
		flag.PrintDefaults()
	}
}

func testxss(p string, v bool) {

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	time.Sleep(500 * time.Microsecond)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		link := scanner.Text()
		u, _ := url.Parse(link)
		if v == true {
			fmt.Println("[+] Testing URL : " + link)
		}
		 q, err := url.ParseQuery(u.RawQuery)
       		 if err != nil {
            		fmt.Println(err)
            		return
        	}
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
