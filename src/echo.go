package main

/*
丸括弧で同じ種類の宣言をひとまとめに出来る
 */
import (
	"os"
	"flag" //コマンドラインオプションのパーサ
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space = " "
	Newline = "\n"
)

func main() {
	flag.Parse() //パラメータリストを調べてflagに設定
	/*
	var 変数名 型 = 初期値
	初期値から型がわかるので
	var s = ""
	や、もっと短く
	s := """
	なんかもできます
	*/
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
