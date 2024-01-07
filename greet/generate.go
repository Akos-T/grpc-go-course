package main

//go:generate buf format ./proto -w
/// Uncomment the following line to run the linter (currently fails; hence commented out)
///go:generate buf lint ./proto
//go:generate buf generate ./proto
//go:generate go build -o bin/greet/server ./server
//go:generate go build -o bin/greet/client ./client
