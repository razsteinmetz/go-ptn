# go-ptn - Parse Torrent File Name

[![GoDoc](https://godoc.org/github.com/razsteinmetz/go-ptn?status.svg)](https://godoc.org/github.com/razsteinmetz/go-ptn)
[![License](https://img.shields.io/github/license/razsteinmetz/go-ptn.svg)](https://github.com/razsteinmetz/go-ptn/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/razsteinmetz/go-ptn)](https://goreportcard.com/report/github.com/razsteinmetz/go-ptn)

> Extract media information from movie and tv series filename

A port of  `middelink/go-parse-torrent-name`  awesome
[library](https://github.com/middelink/go-parse-torrent-name).

Extract all possible media information present in filenames. Multiple regex
rules are applied on filename string each of which extracts corresponding
information from the filename. If a regex rule matches, the corresponding part
is removed from the filename. In the end, the remaining part is taken as the
title of the content.

## Why?

Online APIs by providers like
[TMDb](https://www.themoviedb.org/documentation/api),
[TVDb](http://thetvdb.com/wiki/index.php?title=Programmers_API) and
[OMDb](http://www.omdbapi.com/) don't react to well to search
queries which include any kind of extra information. To get proper results from
these APIs, only the title of the content should be provided as the search
query where this library comes into play. The accuracy of the results can be
improved by passing in the year which can also be extracted using this library.

This port assumes the files folllow the current (2022) formats, and puts less focus on older naming styles.
This port fixed some bugs with file extentions (containter) and various title issues.
It includes a test file which can easily be updated as a json file.

## Information extracted

below is the resulting struct that the function returns.

```
Title      string `json:"title,omitempty"`
Season     int    `json:"season,omitempty"`
Episode    int    `json:"episode,omitempty"`
Year       int    `json:"year,omitempty"`
Resolution string `json:"resolution,omitempty"` //1080p etc
Quality    string `json:"quality,omitempty"`
Codec      string `json:"codec,omitempty"`
Audio      string `json:"audio,omitempty"`
Service    string `json:"service,omitempty"` // NF etc
Group      string `json:"group,omitempty"`
Region     string `json:"region,omitempty"`
Extended   bool   `json:"extended,omitempty"`
Hardcoded  bool   `json:"hardcoded,omitempty"`
Proper     bool   `json:"proper,omitempty"`
Repack     bool   `json:"repack,omitempty"`
Container  string `json:"container,omitempty"`
Widescreen bool   `json:"widescreen,omitempty"`
Website    string `json:"website,omitempty"`
Language   string `json:"language,omitempty"`
Sbs        string `json:"sbs,omitempty"`
Unrated    bool   `json:"unrated,omitempty"`
Size       string `json:"size,omitempty"`
Threed     bool   `json:"3d,omitempty"`
Country    string `json:"country,omitempty"`   // two letters uppercase at the end of the title US or UK only for now
IsMovie    bool   `json:"ismovie"` // true if this is a movie, false if tv show
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/razsteinmetz/go-ptn"
)

func main() {
	filename := "series.title.s01e03.720p.hdtv-GROUP.avi"
	info,_ := ptn.Parse(filename)
	fmt.Println(info)
}

// info has all the data you need
```

PTN works well for both movies and TV episodes. All meaningful information is
extracted and returned together in a dictionary. 


## Install

### Automatic

go-ptn can be installed using `go get`.

```sh
$ go get github.com/razsteinmetz/go-ptn
```

### Manual

First clone the repository.

```sh
$ git clone https://github.com/razsteinmetz/go-ptn go-ptn && cd go-ptn
```

And run the command for installing the package.

```sh
$ go install .
```


