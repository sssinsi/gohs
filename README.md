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
$gohs a.file 3
1
2
3

4
5
```

# option
gohs has some options.

`-o` : set destination file path.

`-s` : set split string.(ex '==========', '---------')