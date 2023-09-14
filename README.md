# Golang-Demoapp

Interesting in learning some Golang. Using this repo as a way to hack together some code and integrate it with our CICD processes. 

## Issues

After switching my runtime stage from Alpine to a Debian Slim base image, my original endpoint was returning this error: 

> failed to get releases: Get "https://api.coinbase.com/v2/prices/BTC-USD/sell": tls: failed to verify certificate: x509: certificate signed by unknown authority: bad response from API

The response error is related to SSL/TLS certificate verification when the application tries to make an HTTPS request to the Coinbase API. The error message "certificate signed by unknown authority" indicates that the root certificate authority (CA) that signed the server's certificate is not trusted by your application/container.

Alpine uses a different set of root certificates than Debian, the Debian image you're using doesn't come with the necessary root certificates installed to trust the Coinbase API's certificate.

Adding this to the `runtime` stage solved the problem: 

```
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
```

## Golang Base

This repo is used as a scaffolding for new Golang projects within a Docker container. 

## Docker

Use the dockerfile to build a development environment locally. You can update the version of go you install in your container by updating the version of goland used in the base images of the Dockerfile. 

Once built, simply run `docker run -it -v ${PWD}:/work go sh`. The use of the volume will allow you to use your IDE outside of your running docker container. 

### Build Local Development Container

```
docker build --build-arg UID=$(id -u) --build-arg GID=$(id -g) -t golang:dev --target dev .
```

The `--build-arg`s map your local user to the default used on the dockerfile's dev target. 

The `--target dev` will only build the dev base defined in your Dockerfile. 

### Run the Dev Target

```
docker run -it -v ${PWD}:/work -p 8000:8000 golang:dev bash 
```
> port 8000 is used for the godoc command below..

## Go Module

```
go mod init github.com/andygodish/golang-base
```

## Godoc

The dockerfile dev target will install the godoc module automatically. Make sure you are exposing port 8000 in your `docker run` command (`-p 8000:8000`)

```
godoc -http :8000
```

## Testify

https://github.com/stretchr/testify

```
go get github.com/stretchr/testify
```

## Benchmarking

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

```
go test -bench=.
```
## Built-in Testing Coverage

Go's built-in testing toolkit features a coverage tool.

```
go test -cover
```

## Variadic Functions

https://gobyexample.com/variadic-functions




