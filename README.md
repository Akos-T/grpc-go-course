# gRPC [Golang] Master Class: Build Modern API & Microservices | Udemy

## Overview

This repo has been created during doing this <a href="https://www.udemy.com/course/grpc-golang/">gRPC course on Udemy</a>.

The `greet` and `blog` app I did together with the instructor. <br/>
The `calculator` app I did mostly on my own. (Of course at the end I checked the solution and fixed mine if it was needed). <br/>
In `Section 9 - Advanced features` there are errors, deadlines/timeouts, SSL and reflection that I did along with the instructor and affected `calculator` as well as `greet`.

After the course I added:

- [Example interceptor](https://github.com/Akos-T/grpc-go-course/blob/main/greet/server/main.go) to the `greet` server ()
- [Example tests](https://github.com/Akos-T/grpc-go-course/blob/main/greet/server/greet_test.go) to the `greet` app

## Changes compared to the course

- Instead of using `--go_opt=modules=github.com/Akos-T/grpc-go-course` I'm using `--go_opt=paths=source_relative`. Same for `--go-grpc_opt`.
- Instead of using a `makefile`, I'm using `Buf CLI` and `go generate`.

## How to build the apps (`greet` / `calculator`)

1. **This is only needed for greet:** In the ssl folder, run `ssl.sh` to generate your self-signed certificates for TLS in greet
2. Navigate to the project folder (greet or calculator)
3. Run `go generate`
4. Start the server:
   1. In case of greet: `./bin/greet/server`
   2. In case of calculator: `./bin/server`
5. Start the client:
   1. In case of greet: `./bin/greet/client`
   2. In case of calculator: `./bin/client`

## How to build the `blog` app

1. Navigate to the blog folder
2. Run `go generate`
3. Run `docker compose up`
4. In another terminal navigate to the blog folder and run `./bin/server`
5. In another terminal navigate to the blog folder and run `./bin/client`

## How to run the `greet` tests

1. Navigate to the greet folder
2. Run `go test ./server`

## Certification - 2023-10-22

- Udemy: https://www.udemy.com/certificate/UC-f25db073-5270-4204-98df-5e2fe9913c55/
- LinkedIn Post: https://www.linkedin.com/feed/update/urn:li:activity:7121857412495532032/

## Useful links

- https://www.udemy.com/course/grpc-golang/
- https://grpc.io/docs/languages/go/quickstart/
- https://buf.build/docs/installation
- https://github.com/Clement-Jean/grpc-go-course/blob/master/ssl/ssl.sh
- https://github.com/ktr0731/evans
