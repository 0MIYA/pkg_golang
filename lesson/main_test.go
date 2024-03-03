package main
import "testing"//テスト用のpkg

var Debug bool = false//変数デバック　初期値をfalse 
//テスト用の処理 
//★テストのルールとしてTest IsOneのテストと始まる関数がテスト用の関数と表す
func TestIsOne(t *testing.T){//tをテスティングのTと渡す
	i :=1//iを1として変数を作り

	if Debug {
		//Debug変数がtruseであればテストを行わないが上記でfolseとしているので
		t.Skip("スキップさせる")//trueの場合こちらが実行される
	}
	v :=IsOne(i)//ISOne関数でfを渡してbool値を受け取る
	//bool値がtureであれば分岐せず、falseであればエラーが出力される

	if !v{
		t.Errorf("%v != %v", i,1)
	}
}
//cd lesson
//cd main_test.go
//go test　で実行
//go test -v でテストの詳細を表示

//No2_integer_type.go:4:1: expected 'package', found 'func'>解決方法No2と被っていた為別のフォルダを作成