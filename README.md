# Golang tutorial

## memo
### WindowsでのGolangの環境変数

```
set GOROOT=C:\Go
set Path=%Path%;C:\Go\bin
set GOPATH=C:\Users\tenten0213\goplayground
```

`GOPATH`は適当な値で良いっぽい。IntelliJ IDEAでGoのプラグインを入れた際に設定していないと怒られたので設定した。

### Go commands

* run

```
$ go run echo.go hoge
hoge
```

* build

```
$ go build echo.go

$ ls
echo.exe  echo.go

$ echo.exe hoge
hoge
```

## Link

* [Goプログラミング言語のチュートリアル - golang.jp](http://golang.jp/go_tutorial)
* [Effective Go — プログラミング言語 Go ドキュメント v0.1 documentation](http://go.shibu.jp/effective_go.html)
