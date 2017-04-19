# Go Game of Life (GoGoL)

Simulate Game of Life in Go using looping over arrays and using goroutines+channels

![Look at their pathetic lives](http://www.giphy.com/gifs/xUA7aPDIIviSXVFm12)

## To run

```
go run src/golArrays/gameOfLifeArrays.go
go run src/golChannels/gameOfLifeChannels.go
```

## Time comparison

(This is after adding a return in the first lineof the `PrintOutput` function)

```
$ time go run src/golArrays/gameOfLifeArrays.go 

real	0m0.375s
user	0m0.277s
sys	0m0.112s

$ time go run src/golChannels/gameOfLifeChannels.go

real	0m0.470s
user	0m0.488s
sys	0m0.167s
```

The channels option is consistently slower