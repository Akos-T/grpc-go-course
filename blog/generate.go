package main

//go:generate buf format ./proto -w
//go:generate buf lint ./proto
//go:generate buf generate
//go:generate go build -o ./bin/ ./server
