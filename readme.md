# Learning go in few steps like suggested by [TravisMedia](https://www.youtube.com/watch?v=_6Kvp03srKc&list=PL22QEtzSG4EURqpYaWUJJF64Vbck7WQfA&index=3)

1. [Go - The Complete Developer's Guide By Stephen Grider](https://geni.us/fDTKmK)

2. [Go By Example](https://gobyexample.com/)

3. [The Go Programming Language](https://amzn.to/3At4apj)

4. [Gophercises](https://gophercises.com/)

5. [Building Web Applications With Go by Trevor Sawler](https://geni.us/40uDc6x)

## Some special notes

**This error:**

```error
go pls was not able to find modules in your workspace
```

It happens when you create a new module with a new folder inside of an existing project directory.

To fix it:

1. cd into the new dir
2. run `go mod init module.go`
3. run `go mod tidy`

**The \_ sign: [blank identifier](https://go.dev/doc/effective_go#blank)**
You use it when you have declared but unused variables to let Go know that everything is ok. Especially useful in range loops.

**The nil value**
`nil` in Go is used to validate errors, so far I have been encountering it on these idiomatic expressions:

```go
func (d deck) read_from_file() error {
	content, err := os.ReadFile("../cards/my_cards")
	if err != nil {
		log.Fatal(err)
	}
```

Here we are saying: if `nil` is not == 0 (because nil is effectively 0 in Golang), return an error.

## Heap and Stack in Go

Source: https://medium.com/@quicktechlearn/stack-and-heap-memory-in-golang-eec3fb7ec113

**Sum up**

The language does not speak about stacks and heaps, formally.

In Go, you don’t have direct control over whether a variable is stored in the stack or heap. The Go runtime and compiler make these decisions based on the variable’s scope, lifetime, and size. Stack memory is generally more efficient for small, short-lived variables, while heap memory is used for larger, dynamically allocated data structures.

Go’s automatic memory management, including garbage collection for heap-allocated objects, simplifies memory management for developers and helps prevent common memory-related bugs like memory leaks and use-after-free errors.
