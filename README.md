# SimpleWebScraper
v.0.1 Alpha

## What is it for

Simple web scraper is exactly what the name assumes: it scrapes web sites without complicatd stuff. It does not change the links, just saves the whole website (or part of it, which you specify) to the local disk location, you point to. For example, you may want to use it to make periodic snapshot of your own website.

Caveat: this is not an ftp copy of the site. For example, you won't get php or py files,
you will not get the database, you will only get the HTML and images that the user of the site sees.

## How to run it

This is a command line tool so you run it from terminal / shell / however you call it. 

1. Copy the binary to your path location (or whevere you want, but then it's up you how to start it).

1. Open terminal and run:

In Linux & MacOS:

> `screenscraper <your-site> <local-disk-location`

In Windows:

> `screenscraper.exe <your-site> <local-disk-location`

Your site may be a part of it, e.g. `http://site.com/my/local/folder` For example:

> `screenscraper http://site.com/my/local/folder/ ~/Temp/mysite/`

##Subtle things if you need more

1. Some CMS systems randomly put or omisss index.php from the path. For example,
   in Dripal http://site.com/index.php/node/100 and http://site.com/node/100
   are the same. So scraper will remove "index.php/" from the path to avoid duplication of files.
   
1. Google SItes has a bug with links that look like "//blogspot.com/...". For such links "http:"
   prefix will be automatically added.
   
1. Normally, url args are ignored, however for some CMS some of the need to be processed.
   Use the command line flag --allow_args="comma-separated list" to set it explicitely. 
   The defualt is --allow_args="page", since page happens in some CMS systems without which
   some pages may be missing.
     
   
## How to get the binary:

1. Option 1: Get sources from `https://github.com/eldarm/SimpleWebScraper` and build it
using [Go dev sdk](https://go.dev/doc/install).

1. Option 2: get the binary from `https://github.com/eldarm/SimpleWebScraper/bin`.
 1. scraper.exe for Windows.
 1. scraper.app for new Macs (ARM based)
 1. scraper for Linux

## That's it for now.

Let me know if you need more info.
