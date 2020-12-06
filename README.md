# Advent of Code 2020 Solutions

My set of solutions for [Advent of Code 2020](https://adventofcode.com/2020), written in Golang as while my instinct would be to solve these style of questions in Python, wanted to solidify my Golang as most of my golang experience is from developing web microservices interfacing with cloud services, rather than nitty-gritty coding challenge stylings.

## Running

My inputs from the website are all stored in the `inputs` directory, and at the time of writing these files are effectively hardcoded into the running.

To run a particular day (i.e. _dayone_, _daytwo_, ...)
```
go run main.go -challenge dayone
```

To run all days
```
go run main.go -all
```