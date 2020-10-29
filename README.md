# recheck

**WARNING: β quality code**

recheck is a program for validating regular expressions in go programs.

## Why?

Currently validity of regular expressions is checked at run time, either by
examining returned error of `regexp.Compile` or having `regexp.MustCompile`
panic (usually in `init`).

`recheck` allows you to check your regular expressions at test/lint time.
`receck` examines called to `regexp.*Compile*` and if the regular expression is
a literal string - will check that it's valid.

`recheck` will also examine lines that have a comment in the following format:

    r.HandleFunc("/articles/{category}/{id:[0-9]+}", handler) // recheck:0

Then number after "recheck:" is the argument to check (first one in this case)

## Install

    go get github.com/tebeka/recheck/cmd/recheck
