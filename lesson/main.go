package main

import (
	"fmt"
	"golang_udemy/lesson/alib"
)

/*
func main(){ //エントリーポイントmain関数が実行される場所
	fmt.Println("HelloWorld")
}
*/

//cd　コマンドを使って移動しまくる
//lsでディレクトリ内を見る

/*私の場合＞Desktop＞golangUdemrフォルダ＞
[lsで実行したいテキストファイル確認]main.goを発見＞ターミナルでgo run main.go実行
HelloWorldが出力されたらOK*/

//テスト
func IsOne(i int)bool{
	if i == 1 {//iに１代入
		return true//1であった場合
	}else{
		return false//それ以外の場合
	}
}

func main(){
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s :=[]int{1,2,3,4,5}
	fmt.Println(alib.Average(s))
}
//ターミナル上でmain_test.goを作成