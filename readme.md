# brainfreeze

a tiny brainf\*ck interpreter written in golang

## usage

the package is not yet published to https://pkg.go.dev so if you want to try it out, you have to build it from source

```
git clone https://github.com/0xmukesh/brainfreeze
cd brainfreeze
go build -o brainfreeze cmd/interpreter/main.go
```

create a new file ending with .bf (test.bf)

```
>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+
+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-
]<+.
```

```
./brainfreeze -file test.brt
```

here is the output for the above code

```
H
e
l
l
o
,

W
o
r
l
d
!
```
