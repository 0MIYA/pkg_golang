package alib
import "testing"

var Debug bool = false

//テストの内容
func TestAverage(t *testing.T){
	if Debug{//Debugがtruseであればスキップ
		t.Skip("スキップします")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v !=3 {
		t.Errorf("%T != %T", v, 3)
	}
}
//cd ..で１つディレクトリ戻る
// go test ./alibで相対パス実行

