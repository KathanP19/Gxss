# Gxss
A Light Weight Tool for checking reflecting Parameters in a URL. Inspired by [kxss](https://github.com/tomnomnom/hacks/tree/master/kxss) by [@tomnomnom](https://twitter.com/TomNomNom).

# Installation

`go get -u github.com/KathanP19/Gxss`

# Usage

```
                  
 _____ __ __ _____ _____ 
|   __|  |  |   __|   __|
|  |  |-   -|__   |__   |
|_____|__|__|_____|_____|
                         
        0.1 - @KathanP19

  -c int
        Set the Concurrency  (default 50)
  -p string
        Payload you want to Send to Check Refelection
  -v    Verbose mode
```

* Checking Single Url

    `echo "https://target.com/some.php?first=hello&last=world | Gxss -c 100 -p Payload`
    
* Checking List of Urls

    `cat urls.txt | Gxss -c 100 -p XssReflected`

# Use Case

`echo "testphp.vulnweb.com" | waybackurls | httpx -silent | Gxss -c 100 -p Xss | grep "URL" | cut -d '"' -f2 | sort -u | dalfox pipe` 

* [Dalfox](https://github.com/hahwul/dalfox) is Xss Scanner by [@hahwul](https://twitter.com/hahwul)

# Thanks To

* [Zoid](https://twitter.com/z0idsec) for helping me out with code.

# To Support Me 

* You Can Buy Me A Coffee

    <a href="https://www.buymeacoffee.com/kathanp19" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
