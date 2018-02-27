# TamaGo

[![Build Status](https://travis-ci.org/SebastienBoisard/TamaGo.svg?branch=master)](https://travis-ci.org/SebastienBoisard/TamaGo)
[![codecov](https://codecov.io/gh/SebastienBoisard/TamaGo/branch/master/graph/badge.svg)](https://codecov.io/gh/SebastienBoisard/TamaGo)
[![Coverage Status](https://coveralls.io/repos/github/SebastienBoisard/TamaGo/badge.svg)](https://coveralls.io/github/SebastienBoisard/TamaGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/SebastienBoisard/TamaGo)](https://goreportcard.com/report/github.com/SebastienBoisard/TamaGo)
[![codebeat badge](https://codebeat.co/badges/072b5178-c192-4b3f-99c0-bae916589180)](https://codebeat.co/projects/github-com-sebastienboisard-tamago)

TamaGo is a simple note taking system developed on Go.  

The goal of this experimental project is to discover and use different tools with Go:
  - [MongoDB](https://www.mongodb.com/), a NoSQL document database
  - [Mgo](http://labix.org/mgo),  MongoDB driver for Go   
  - [Fin](https://github.com/gin-gonic/gin), a HTTP web framework written in Go
  - [Dep](https://github.com/golang/dep), Go dependency management tool 
  - [Cobra](https://github.com/spf13/cobra), a Commander for modern Go CLI interactions
  - [Jsonrpc](https://golang.org/pkg/net/rpc/jsonrpc/), a light weight remote procedure call protocol library implementation in Go 


## Architecture

TamaGo is structured in 2 servers: 
  - a front server: a simple webserver based on Gin, which calls remote procedures through a websocket with a 
  Javascript library called "[Simple JsonRPC](https://github.com/jershell/simple-jsonrpc-js)"
  - a back server: a small RPC server, which stores an retrieves notes in a MongoDB database.
  
## Build

```
go build -o tamago main/tamago.go
```


## Usage

To run the front-end part of the project:
```
tamago web
```

To run the back-end part of the project: 
```
tamago rpc
```

## Todo

  - Create a web design
  - Add several actions (delete and update a note for example)
  - Add unit tests

  
  
http://denis.papathanasiou.org/posts/2012.10.14.post.html
http://spf13.com/presentation/mongodb-and-go/
http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
