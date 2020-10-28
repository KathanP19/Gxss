# Gxss v2.0

A Light Weight Tool for checking reflecting Parameters in a URL. Inspired by [kxss](https://github.com/tomnomnom/hacks/tree/master/kxss) by [@tomnomnom](https://twitter.com/TomNomNom).

# Installation

`go get -u github.com/KathanP19/Gxss`

# Usage

```
                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
        2.0 - @KathanP19

  Usage of ./Gxss:
  -c int
        Set the Concurrency (default 50)
  -o string
        Save Result to OuputFile
  -p string
        Payload you want to Send to Check Refelection (default "Gxss")
  -v    Verbose mode

```

* Checking Single Url

    `echo "https://target.com/some.php?first=hello&last=world | Gxss -c 100 `
    
* Checking List of Urls

    `cat urls.txt | Gxss -c 100 -p XssReflected`

* Save Urls Which have Reflecting Params in a file for further analysis

    `cat urls.txt | Gxss -c 100 -o Result.txt`


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
# Use Case

`echo "testphp.vulnweb.com" | waybackurls | httpx -silent | Gxss -c 100 -p Xss | sort -u | dalfox pipe` 

* [Dalfox](https://github.com/hahwul/dalfox) is Xss Scanner by [@hahwul](https://twitter.com/hahwul)

# TODO

- [ ] Add an option for user to add there own headers
- [ ] Add an option for User-Agent

# Thanks To

* [Zoid](https://twitter.com/z0idsec) for helping me out with code.

# To Support Me 

* You Can Buy Me A Coffee

    <a href="https://www.buymeacoffee.com/kathanp19" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
