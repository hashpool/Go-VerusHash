# Getting Started

```
$ git clone git@github.com:hashpool/Go-VerusHash.git
$ cd Go-VerusHash
$ git checkout hashpool
$ cd ./verushash
$ cmake .  
$ make  
$ cp libverushash.a ./build/libverushash.a 
$ swig -c++ -go -intgosize 64 -cgo verushash.i
$ cd ..
$ go run main.go
```
