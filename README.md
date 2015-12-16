loops-20151217-tp
=================

`loops-20151217-tp` holds the code and exercizes for
[docker@LoOPS](http://talks.godoc.org/github.com/sbinet/talks/2015/20151217-loops-docker/talk.slide)

The programming language used is [Go](https://golang.org).
You won't be programming in [Go](https://golang.org) though, just modifying some
strings being printed out.

This repository holds 2 applications:

- `hello`, a command-line application that prints a greeting on the screen and tries to launch a sub-process (`pkg-config`, so nothing dangerous)
- `web-app`, a simple `http` server that serves a single page at `0.0.0.0:8080`

To build and install the `hello` command:

```sh
$> cd loops-20151217-tp
$> go install ./hello
$> hello
hello loOPS-20151217
running pkg-config --cflags python2 now...
-I/usr/include/python2.7
```

Similarly, for the `web-app` server, this time using the full repository name:

```sh
$> go install github.com/sbinet/loops-20151217-tp/web-app
$> web-app
2015/12/16 16:43:58 listening on: http://localhost:8080
```

Another possibility is to use `go run` that compiles on the fly the executable and runs it:

```sh
$> go run ./web-app/main.go
2015/12/16 16:43:58 listening on: http://localhost:8080
```
