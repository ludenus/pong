# Pong

A simple http echo server written in golang based on https://github.com/kataras/iris


## Build
default settings
```
$ go build pong.go
```

small binary
```
$ GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" pong.go
$ upx --ultra-brute ./pong
```

## Run locally

run pong server on default port 80
```
$ ./pong
```

specify custom port number via env var
```
$ PONG_LISTENING_ADDRESS="0.0.0.0:7654" ./pong
```

or via command line arg (overrides env var)
```
$ ./pong :8181
```


## Build docker

```
$ docker build -t ludenus/pong:0.0.1 .
```


## Run docker

directly
```
$ docker run -ti --rm --name pong ludenus/pong:0.0.1
```

or via docker-compose
```
$ docker-compose up
```


## Usage

```
$ ./pong :1234
$ curl -XGET http://localhost:1234/ping/blablabla
{"pong":"blablabla"}
```


## Troubleshooting
```
$ ./pong 1234
[ERRO] 2018/09/23 19:56 listen tcp: address 1234: missing port in address
```

Forgot column, should be `:1234` instead of `1234`

## Versions
works on `Ubuntu 18.04.1 LTS`
```
$ uname -a
Linux hostname 4.15.0-34-generic #37-Ubuntu SMP Mon Aug 27 15:21:48 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```

```
$ go version
go version go1.10.1 linux/amd64
```
