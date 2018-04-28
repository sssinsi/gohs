# gohs
gohs is horizontal split cli-tool for go.

# sample
You need file path and split count.
Default split string is blank (LF only).
```a.file
$cat a.file
1
2
3
4
5
```

```command
$gohs -p a.file -c 3 -s "************"
1
2
3
************
4
5
```

# option
gohs has some options.
```
  -c uint
        [Option] : Set split line count (default 5)
  -o string
        [Option] : Set output file path
  -p string
        [required] : Set source file path
  -s string
        [Option] : Set split string(ex: "************", //////////// )
```