# Learning go in few steps like suggested by [TravisMedia](https://www.youtube.com/watch?v=_6Kvp03srKc&list=PL22QEtzSG4EURqpYaWUJJF64Vbck7WQfA&index=3)

1. [Go - The Complete Developer's Guide By Stephen Grider](https://geni.us/fDTKmK)

2. [Go By Example](https://gobyexample.com/)

3. [The Go Programming Language](https://amzn.to/3At4apj)

4. [Gophercises](https://gophercises.com/)

5. [Building Web Applications With Go by Trevor Sawler](https://geni.us/40uDc6x)

## Some special notes

This error:

```error
go pls was not able to find modules in your workspace
```

It happens when you create a new module with a new folder inside of an existing project directory.

To fix it:

1. cd into the new dir
2. run `go mod init module.go`
3. run `go mod tidy`
