### Go run

Work only on a .go file with package main, implementing a main() func

What if we go run on a .go file without a main() func ?
Error message:
go run: cannot run non-main package

### Go build

on an executable:
Produce a binary executable file, in the same directory as the .go file

on a package:
nothing happens

### Go install

on an executable:
Ask to set the variable GOBIN if not set.
Produce an binary executable in the ~/go/bin/ directory

on a package:
can't see anything happening (nothing in ~/go/pkg nor in ~/go/bin)
