# gRPC [Golang] Master Class: Build Modern API & Microservices | Udemy

## Overview
This repo has been created during doing this <a href="https://www.udemy.com/course/grpc-golang/">gRPC course on Udemy</a>.

The `greet` app I did together with the instructor. 
The `calculator` app I did on my own. (Of course at the end I checked the solution and fixed mine if it was needed)

## Changes compared to the course
- Instead of using `--go_opt=modules=github.com/Akos-T/grpc-go-course` I'm using `--go_opt=paths=source_relative`. Same for `--go-grpc_opt`.
- Instead of using a `makefile`, I'm using `Buf CLI` and `go generate`.

## How to build the apps
1. Navigate to the project folder (greet or calculator)
2. run `go generate`
3. Start the server:
   1. In case of greet: ./bin/greet/server
   2. In case of calculator: ./bin/server
4. Start the client:
   1. In case of greet: ./bin/greet/client
   2. In case of calculator: ./bin/client
   

## Useful links
- https://www.udemy.com/course/grpc-golang/
- https://grpc.io/docs/languages/go/quickstart/
- https://buf.build/docs/installation
