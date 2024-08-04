# Project is now Deprecated as Dalfox has discovery option which can be used to identify reflecting params `dalfox file urls.txt --skip-xss-scanning -o reflecting.txt`

# Gxss v4.0

A Light Weight Tool for checking reflecting Parameters in a URL. Inspired by [kxss](https://github.com/tomnomnom/hacks/tree/master/kxss) by [@tomnomnom](https://twitter.com/TomNomNom).

# Installation

`go install github.com/KathanP19/Gxss@latest`

* If the above step doesn't work then you can try pre-built binary file from here
  https://github.com/KathanP19/Gxss/releases

# Usage

```
                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
        4.0 - @KathanP19

Usage of Gxss:
  -c int
        Set the Concurrency (default 50)
  -d string
        Request data for POST based reflection testing
  -h value
        Set Custom Header.
  -o string
        Save Result to OutputFile
  -p string
        Payload you want to Send to Check Reflection (default "Gxss")
  -u string
        Set Custom User agent. Default is Mozilla (default "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
  -v    Verbose mode
  -x string
        Proxy URL. Example: http://127.0.0.1:8080
```

* Checking Single Url

    `echo "https://target.com/some.php?first=hello&last=world" | Gxss -c 100 `
    
* Checking List of Urls

    `cat urls.txt | Gxss -c 100 -p XssReflected`

* Save Urls Which have Reflecting Params in a file for further analysis

    `cat urls.txt | Gxss -c 100 -o Result.txt`

* For verbose mode `-v`

    `cat urls.txt | Gxss -c 100 -o Result.txt -v `
    
* Send Custom Header `-h`
    
    `cat urls.txt | Gxss -c 100 -p Xss -h "Cookie: Value"`
    
* Send Custom User-Agent `-u`
    
    `cat urls.txt | Gxss -c 100 -p Xss -h "Cookie: Value" -u "Google Bot"`


# How It Works
1. It takes Urls from STDIN
2. It check for the reflected value on params one by one. (There are some tool like qsreplace which replace all params value but gxss checks payload one by one which makes it different from all those tools.)
```
For Example- 
Url is https://example.com/?p=first&q=second

First it will check if p param reflects
https://example.com/?p=Gxss&q=second

Then it will check if q param reflects
https://example.com/?p=first&q=Gxss
```
3. If reflection for any param is found it tells which param reflected in response.

[![asciicast](https://asciinema.org/a/84mXOOcDrxzZ3eyW16Ap3eHwX.svg)](https://asciinema.org/a/84mXOOcDrxzZ3eyW16Ap3eHwX)

# Use Case or How to add to your workflow

`echo "testphp.vulnweb.com" | waybackurls | httpx -silent | Gxss -c 100 -p Xss | sort -u | dalfox pipe` 

* [Dalfox](https://github.com/hahwul/dalfox) is Xss Scanner by [@hahwul](https://twitter.com/hahwul)

# TODO

- [ ] TimeOut Option. 
- [x] Add Post Method Support.
- [x] Add Proxy Support.
- [x] Add an option for user to add there own headers
- [x] Add an option for User-Agent

# Thanks To

* [Zoid](https://twitter.com/z0idsec) for helping me out with code.
* [Parth Parmar](https://twitter.com/Parth97531) for adding Custom Header and User-Agent Support.
* [Luska](https://github.com/LuskaBol) for adding proxy support and custom post data support

# To Support Me 

* You Can Buy Me A Coffee

    <a href="https://www.buymeacoffee.com/kathanp19" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
