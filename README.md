# recheck

**PRE ALPHA: No Code Yet**

recheck is a program for validating regular expressions in go programs.


## Why?

Currently validity of regular expressions is checked at run time, either by
examining returned error of `regexp.Compile` or having `regexp.MustCompile`
panic (usually in `init`).

recheck allows you to check your regular expressions at test/lint time.
