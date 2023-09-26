# gRPC [Golang] Master Class: Build Modern API & Microservices | Udemy

## Overview
This repo has been created during doing this <a href="https://www.udemy.com/course/grpc-golang/">gRPC course on Udemy</a>.

The `greet` app I did together with the instructor. <br/>
The `calculator` app I did mostly on my own. (Of course at the end I checked the solution and fixed mine if it was needed). <br/>
In `Section 9 - Advanced features` there are errors, deadlines/timeouts, SSL and reflection that I did along with the instructor and affected `calculator` as well as `greet`.

## Changes compared to the course
- Instead of using `--go_opt=modules=github.com/Akos-T/grpc-go-course` I'm using `--go_opt=paths=source_relative`. Same for `--go-grpc_opt`.
- Instead of using a `makefile`, I'm using `Buf CLI` and `go generate`.

## How to build the apps
1. **This is only needed for `greet`:** In the ssl folder, run ssl.sh to generate your self-signed certificates for TLS in greet
2. Navigate to the project folder (greet or calculator)
3. run `go generate`
4. Start the server:
   1. In case of greet: ./bin/greet/server
   2. In case of calculator: ./bin/server
5. Start the client:
   1. In case of greet: ./bin/greet/client
   2. In case of calculator: ./bin/client
   

## Useful links
- https://www.udemy.com/course/grpc-golang/
- https://grpc.io/docs/languages/go/quickstart/
- https://buf.build/docs/installation
- https://github.com/Clement-Jean/grpc-go-course/blob/master/ssl/ssl.sh
- https://github.com/ktr0731/evans