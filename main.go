package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

var (
	concurrency int
	verbose     bool
	outputFile  string
	payload     string
)

func banner() {
	fmt.Println(`                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
	2.0 - @KathanP19
	`)
}

func main() {
	flag.IntVar(&concurrency, "c", 50, "Set the Concurrency")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.StringVar(&payload, "p", "Gxss", "Payload you want to Send to Check Refelection")
	flag.StringVar(&outputFile, "o", "", "Save Result to OuputFile")
	flag.Parse()

	if verbose == true {
		banner()
	}

	if payload != "" {

		if outputFile != "" {
			emptyFile, err := os.Create(outputFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(emptyFile)
			emptyFile.Close()

			var wg sync.WaitGroup
			for i := 0; i < concurrency; i++ {
				wg.Add(1)
				go func() {
					testref(payload, verbose, outputFile)
					wg.Done()
				}()
				wg.Wait()
			}

		} else {

			var wg sync.WaitGroup
			for i := 0; i < concurrency; i++ {
				wg.Add(1)
				go func() {
					testref(payload, verbose, outputFile)
					wg.Done()
				}()
				wg.Wait()
			}
		}
	} else {
		flag.PrintDefaults()
	}
	if verbose == true {
		fmt.Println("\nFinished Checking, Thank you for using Gxss.")
	}
}

func testref(payload string, verbose bool, outputFile string) {
	time.Sleep(500 * time.Microsecond)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		link := scanner.Text()
		checkreflection(link)
	}

}

func checkreflection(link string) {
	decoded, _ := url.QueryUnescape(link)
	u, err := url.Parse(decoded)
	if err != nil {
		decoded := url.QueryEscape(link)
		v, err := url.Parse(decoded)
		if err != nil {
			fmt.Printf("Error is %e", err)
		}
		u = v
	}

	if verbose == true {
		fmt.Println("[+] Testing URL : " + link)
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Printf("Error is %e", err)
	}
	for key, value := range q {
		var tm string = value[0]
		q.Set(key, payload)
		u.RawQuery = q.Encode()
		_, body, _ := requestfunc(u.String())

		re := regexp.MustCompile(payload)
		match := re.FindStringSubmatch(body)
		if match != nil {
			if verbose == true {
				fmt.Printf("Url : %q\n", u)
				fmt.Printf("Reflected Param : %q\n", key)
			} else {
				fmt.Fprint(os.Stderr, u.String()+"\n")
			}
			if outputFile != "" {
				f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				if _, err := f.WriteString(u.String() + "\n"); err != nil {
					log.Fatal(err)
				}
				f.Close()
			}
		}
		q.Set(key, tm)
	}
}

func requestfunc(u string) (resp gorequest.Response, body string, errs []error) {

	resp, body, errs = gorequest.New().Get(u).TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error { return http.ErrUseLastResponse }).
		Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36").
		End()

	return resp, body, errs
}
