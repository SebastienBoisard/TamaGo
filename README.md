# TamaGo

[![Build Status](https://travis-ci.org/SebastienBoisard/tamago.svg?branch=master)](https://travis-ci.org/SebastienBoisard/tamago)
[![Go Report Card](https://goreportcard.com/badge/github.com/SebastienBoisard/tamago)](https://goreportcard.com/report/github.com/SebastienBoisard/tamago)
[![codebeat badge](https://codebeat.co/badges/072b5178-c192-4b3f-99c0-bae916589180)](https://codebeat.co/projects/github-com-sebastienboisard-tamago)

TamaGo is a simple note taking system developed on Go.  

The goal of this experimental project is to discover and use different tools with Go:
  - [MongoDB](https://www.mongodb.com/), a NoSQL document database
  - [Mgo](http://labix.org/mgo), a MongoDB driver for Go   
  - [Gin](https://github.com/gin-gonic/gin), a HTTP web framework written in Go
  - [Dep](https://github.com/golang/dep), Go dependency management tool 
  - [Cobra](https://github.com/spf13/cobra), a Commander for modern Go CLI interactions
  - [Jsonrpc](https://golang.org/pkg/net/rpc/jsonrpc/), a light weight remote procedure call protocol library implementation in Go 


## Architecture

TamaGo is structured in 2 servers: 
  - a front server: a simple webserver based on Gin, which calls remote procedures through a websocket with a 
  Javascript library called "[Simple JsonRPC](https://github.com/jershell/simple-jsonrpc-js)"
  - a back server: a small RPC server, which stores an retrieves notes from a MongoDB database.
  
## Build

```
go build -o tamago main/tamago.go
```


## Usage

To run the front-end part of the project:
```
tamago front
```

To run the back-end part of the project: 
```
tamago back
```

## Todo

  - Create a web design
  - Add several actions (delete and update a note for example)
  - Add unit tests

  
  
