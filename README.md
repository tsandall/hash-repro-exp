# Example

```
(master)torin:~/go/src/github.com/tsandall/hash-repro-exp$ docker run -it -v $PWD:/go/src/github.com/tsandall/hash-repro-exp golang:1.8.7 go run /go/src/github.com/tsandall/hash-repro-exp/main.go
(master)torin:~/go/src/github.com/tsandall/hash-repro-exp$ docker run -it -v $PWD:/go/src/github.com/tsandall/hash-repro-exp golang:1.9.0 go run /go/src/github.com/tsandall/hash-repro-exp/main.go
panic: hashes differ: tgkvnvn1r 5561414296527422833 1767073917199631498

goroutine 1 [running]:
main.F.func1(0x49ff60, 0xc42000e1f0, 0x4d2e1cd6494bc571)
	/go/src/github.com/tsandall/hash-repro-exp/main.go:37 +0x18c
main.F(0x49ff60, 0xc42000e1f0)
	/go/src/github.com/tsandall/hash-repro-exp/main.go:40 +0x74
main.main()
	/go/src/github.com/tsandall/hash-repro-exp/main.go:14 +0xb1
exit status 2
```
