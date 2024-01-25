# Gophercises - Exercise 4: HTML Link Parser
This is my solution for the excercise 4 of Gophercises [gophercises/link](https://github.com/gophercises/link).

For this exercise I had to create a simple parser that gets the anchor tags from an html document and returns them in a slice of Link instances.
```go
type Link struct {
  Href string
  Text string
}
```
To achive this I made use of the [x/net/html](https://pkg.go.dev/golang.org/x/net/html) package, and I followed a TDD strategy.
