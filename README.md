# Introduction  

Webgin adalah skeleton web golang using framwork [gin](https://github.com/gin-gonic/gin).  

## Getting Started  

## Prerequisites  

1. Install latest [golang](https://golang.org/doc/install)  

## Installation process  

1. Clone Repository  

```sh
$ git clone git@github.com:rulzi/webgin.git
```  

2. Get Depedency repository  

```sh  
$ go get ./...
```  

3. Open config.yml in project root dan setting your environment  

```sh  
$ cp config.yml.example config.yml  
```  

## Usage  

1. use golang run

```sh  
$ go run cmd/main.go 
```  

## Makefile  

- build application golang to make sure application runing well  

```sh  
$ make build  
```  

- build application and generate binary image  

```sh  
$ make build-generate  
```  

## Key Directory

* `cmd`: Main Golang
* `assets`: Assey file for css, golang, image
* `internal`: All file untuk internal
* `internal/config`: Config File
* `internal/constant`: Constant File
* `internal/helpers`: Helper File
* `internal/middleware`: Middleware File
* `internal/migration`: File for setting Migration Database
* `internal/pkg/controller`: All Controllrt Layer (This layer will act as the presenter. Decide how the data will presented. Could be as REST API, or HTML File.
This layer also will accept the input from user. Sanitize the input and sent it to Usecase layer.)
* `internal/pkg/form_validation`: Struct File Validation
* `internal/pkg/models`: All Models layer (Same as Entities, will used in all layer. This layer, will store any Object’s Struct and its method)
* `internal/pkg/repository`: Repository layer (Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.)
* `internal/pkg/usecase`: All Usecase Layer (This layer will act as the business process handler. Any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here.)
* `internal/provider`: Provide function to communicate with gin
* `internal/server`: Run Server
* `internal/services`: Service for db
* `internal/template`: Template file
* `internal/views`: Views File
