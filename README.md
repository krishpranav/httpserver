# httpserver
A simple http server built in golang

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

<img src=""></img>

# Installation
```
git clone https://github.com/krishpranav/httpserver
cd httpserver
go get
go build cmd/httpserver/httpserver.go
```

# Output
```
$ ./httpserver
2021/01/11 21:40:48 Serving . on http://0.0.0.0:8000/...
2021/01/11 21:41:15 [::1]:50181 "GET / HTTP/1.1" 200 383
2021/01/11 21:41:15 [::1]:50181 "GET /hello.html HTTP/1.1" 404 19
```
