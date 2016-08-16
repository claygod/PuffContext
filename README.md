# PuffContext
Library PuffContext differs from the basic library Context so that it contexts exist in the form of layers, and this layer can be removed if necessary variables.

[![API documentation](https://godoc.org/github.com/claygod/PuffContext?status.svg)](https://godoc.org/github.com/claygod/PuffContext)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/PuffContext)](https://goreportcard.com/report/github.com/claygod/PuffContext)

# Usage

An example of using the PuffContext Library:

```go
package main

import (
	"fmt"
	"github.com/claygod/PuffContext"
)

func main() {
    fmt.Print("\nTesting PuffContext library")

    fmt.Print("\n1) create and fill the root context:")
    c1 := PuffContext.New()
    c1.Set("a", 5)
    c1.Set("b", 6)
    fmt.Print("\n Read variable in c1: 'a' -> ", c1.Get("a")) // --> 5
    fmt.Print("\n Read variable in c1: 'b' -> ", c1.Get("a")) // --> 6

    fmt.Print("\n2) create a context branch and fill the variables with the same key but different valuest:")
    c2 := c1.Fix()
    c2.Set("a", 15)
    c2.Set("b", 16)
    fmt.Print("\n Read variable in c2: 'a' -> ", c2.Get("a")) // --> 15
    fmt.Print("\n Read variable in c2: 'b' -> ", c2.Get("a")) // --> 16

    fmt.Print("\n3) creating more context branch and add one variable with the same key but different value:")
    c3 := c1.Fix()
    c3.Set("a", 25)
    fmt.Print("\n Read variable in c3: 'a' -> ", c3.Get("a")) // --> 25
    fmt.Print("\n Read variable in c3: 'b' -> ", c3.Get("b")) // --> 6 !!!
	c3.Del("a")
	fmt.Print("\n Read variable in c3: 'a' -> ", c3.Get("a")) // --> 5 !!!
}

```

# API

Methods:
-  *New* - create a new PuffContext
-  *Set* - add the variable in PuffContext.
-  *Get* - get variable (from the context).
-  *Fix* - fix the state to create a branch.
-  *Del* - remove variable (can only be from the current branch!).

# Perfomance

Library 11 shows a slightly higher speed

- BenchmarkPuffContext-4	2000000000	         0.05 ns/op
- BenchmarkContext-4    	2000000000	         0.09 ns/op
