# Galore

A todo app backend in Go, written in an attempt to demonstrate a minimalistic and idiomatic backend that can be used as an example to build web APIs. 

## Features
* Routing
* Middlewares
* Rich error response
* Ease of extension
* Logging

## Dependencies
* github.com/gorilla/context
* github.com/julienschmidt/httprouter
* gopkg.in/yaml.v2


## Installation

##### Import the repository    
`go get github.com/mantishK/galore`
##### Import dependencies    
`cd $GOPATH/github.com/mantishK/galore`    
`go get ./...`
##### Override configurations   

Configurations can be overridden either via file or command line arguments

File    
Copy `config/config.yaml` to any path and set the path name as environment variable "GALORE_CONFIG".

e.g -    
`cp $GOPATH/github.com/mantishK/galore/config/config.yaml /tmp/`    
`export GALORE_CONFIG=/tmp/config.yaml`

Flags    
Configurations can be set as arguments while running the application    

e.g -    
`galore -pg_uname=postgres -pg_pass="" -pg_name=todo -pg_ip=192.168.59.103 -pg_port=5432`


##### Create tables necessary
The sql dump required for Galore is available at `extra/galore.sql`

##### Postman link for checking APIs
Download the collection from this [link](https://www.getpostman.com/collections/984ed914e6aa12ca0270)


## Running the application
### Without installing        
`cd $GOPATH/github.com/mantishK/galore`    
`go run main.go`    

### Installing    
`go install github.com/mantishK/galore`    
`galore`

## Contributing
Contributions welcome via Github pull requests and issues.

## Credits
Most of my work is heavily inspired by [MatRyer's talk](http://go-talks.appspot.com/github.com/matryer/golanguk/building-apis.slide) at Golang UK Conference held at London.

## License
This project is licensed under the MIT License. Please refer the License.txt file.