
# Go Todo

A minimal command-line todo list app that showcases my beginner-level experience with Go.


## Demo

```bash
$ go run main.go add "Learn Go basics"
Task added: Learn Go basics

$ go run main.go add "Build CLI tool"
Task added: Build CLI tool

$ go run main.go list
1. [ ] Learn Go basics
2. [ ] Build CLI tool

$ go run main.go done 1
Marked task 1 as done.

$ go run main.go list
1. [✔] Learn Go basics
2. [ ] Build CLI tool
