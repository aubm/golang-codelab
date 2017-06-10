# Basics

## How it works

You will have to get through all the exercises in the right order, so:

- 01-basic-data-types
- 02-composite-types
- 03-functions/01-simple
- 03-functions/02-recursion
- etc...

An exercise is considered to be done when all the associated tests successfully pass.
To do that, each exercise package contains an `src.go` file on which you will have to apply
the necessary changes.
In the exercise folder, use this command to run all the tests:

```bash
$ go test
```

For some exercises, first the code will not even compile. That being said, it is recommended
to run them anyway because output errors might give useful hints about how to solve the
exercise. For example, the first exercise `01-basic-data-types` will output:

```bash
$ go test
# github.com/aubm/golang-codelab/exercises/basics/01-basic-data-types
./src.go:4: missing return at end of function
./src.go:7: missing return at end of function
./src.go:10: missing return at end of function
./src.go:13: missing return at end of function
./src.go:16: missing return at end of function
./src.go:19: missing return at end of function
./src.go:22: missing return at end of function
./src.go:25: undefined: who
FAIL	github.com/aubm/golang-codelab/exercises/basics/01-basic-data-types [build failed]
```

At some point the code will successfully compile, then you will have to apply the
appropriate changes to the code, so that the tests exit without errors.

Here is an example output when the code successfully compiles, but the tests are not ok:

```bash
$ go test
--- FAIL: TestMustReturnTheSum (0.00s)
	src_test.go:27: when a=0 and b=0, expected 0, got 3
	src_test.go:27: when a=1 and b=0, expected 1, got 3
	src_test.go:27: when a=3 and b=3, expected 6, got 3
	src_test.go:27: when a=3 and b=4, expected 7, got 3
FAIL
exit status 1
FAIL	github.com/aubm/golang-codelab/solutions/basics/01-basic-data-types	0.012s
```

When the job is done, you should see something like that:

```bash
$ go test
PASS
ok  	github.com/aubm/golang-codelab/solutions/basics/01-basic-data-types	0.009s
```

The tests output will give you useful hints, but reading the yet-to-be completed-exercise
source code in the `src.go` file should give you enough information to get started.

For each exercise, some recommended readings are given, you might want to take a look at them if you
want to get a better understanding on the aspects of the language that the exercise addresses.

### Solutions

If you really need to, you can take a look at possible solutions for each exercise.
Solutions are given in `$GOPATH/src/github.com/aubm/golang-codelab/solutions/basics`.

## Readings

### 01: Basic data types

- https://gobyexample.com/variables
- https://gobyexample.com/if-else

### 02: Composite types

- https://gobyexample.com/arrays
- https://gobyexample.com/slices
- https://gobyexample.com/sorting
- https://gobyexample.com/structs

### 03: Functions

- https://gobyexample.com/functions
- https://gobyexample.com/multiple-return-values
- https://gobyexample.com/variadic-functions
- https://gobyexample.com/defer

### 04: Methods

- https://gobyexample.com/methods

### 05: Interfaces

- https://gobyexample.com/interfaces

### 06: Concurrency

- https://gobyexample.com/goroutines
- https://gobyexample.com/channels
- https://gobyexample.com/select

## What after it?

You completed the basics when this command successfully passes:

```bash
$ go test github.com/aubm/golang-codelab/exercises/basics/...
ok  	github.com/aubm/golang-codelab/exercises/basics/01-basic-data-types	0.015s
ok  	github.com/aubm/golang-codelab/exercises/basics/02-composite-types	0.017s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/01-simple	0.015s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/02-recursion	0.017s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/03-multiple-return-values	0.013s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/04-error-handling	0.030s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/05-variadic-functions	0.018s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/06-deferred-function-calls	0.015s
ok  	github.com/aubm/golang-codelab/exercises/basics/03-functions/07-pointers	0.014s
ok  	github.com/aubm/golang-codelab/exercises/basics/04-methods/01-declarations	0.013s
ok  	github.com/aubm/golang-codelab/exercises/basics/04-methods/02-pointer-receiver	0.012s
ok  	github.com/aubm/golang-codelab/exercises/basics/04-methods/03-encapsulation	0.019s
ok  	github.com/aubm/golang-codelab/exercises/basics/05-interfaces/01-interface-types	0.018s
ok  	github.com/aubm/golang-codelab/exercises/basics/05-interfaces/02-anonymous-interfaces	0.020s
ok  	github.com/aubm/golang-codelab/exercises/basics/05-interfaces/03-discriminating-errors	0.015s
ok  	github.com/aubm/golang-codelab/exercises/basics/06-concurrency/01-goroutines	1.013s
ok  	github.com/aubm/golang-codelab/exercises/basics/06-concurrency/02-multiplexing-with-select	4.022s
ok  	github.com/aubm/golang-codelab/exercises/basics/06-concurrency/03-cancellation	9.024s
ok  	github.com/aubm/golang-codelab/exercises/basics/06-concurrency/04-races	2.523s
```

Go back to the [home page of this repository](https://github.com/aubm/golang-codelab#now-lets-actually-get-started) and follow the link to get started with writing your firsts
real Go programs.
