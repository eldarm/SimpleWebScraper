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

## How to get the binary:

1. Option 1: Get sources from `https://github.com/eldarm/SimpleWebScraper` and build it
using [Go dev sdk](https://go.dev/doc/install).

1. Option 2: get the binary from `https://github.com/eldarm/SimpleWebScraper` when
I'll get to upload them.

## That's it for now.

Let me know if you need more info.
