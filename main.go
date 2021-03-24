package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	concurrency   int
	verbose       bool
	outputFile    string
	payload       string
	useragent     string
	customHeaders string
)

type customh []string

func (m *customh) String() string {
	return "This is custom flag for getting custom headers."
}

func (m *customh) Set(value string) error {
	*m = append(*m, value)
	return nil
}

var custhead customh

func banner() {
	fmt.Println(`                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
	3.0 - @KathanP19
	`)
}

func main() {
	flag.IntVar(&concurrency, "c", 50, "Set the Concurrency")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.StringVar(&payload, "p", "Gxss", "Payload you want to Send to Check Reflection")
	flag.StringVar(&outputFile, "o", "", "Save Result to OutputFile")
	flag.StringVar(&useragent, "u", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36", "Set Custom User agent. Default is Mozilla")
	flag.Var(&custhead, "h", "Set Custom Header.")

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
				fmt.Println(u.String() + "\n")
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

//removed gorequest for more granular access to setting headers.

func requestfunc(u string) (resp *http.Response, body string, errs []error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", u, nil)
	req.Header.Add("User-Agent", useragent)

	//splitting headers and values by using : as separator
	for _, v := range custhead {
		s := strings.SplitN(v, ":", 2)
		req.Header.Add(s[0], s[1])
	}

	//Converting request dump to string for verbose mode
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	if verbose == true {
		fmt.Println(string(requestDump))
	}
	resp, err = client.Do(req)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return resp, bodyString, errs
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
