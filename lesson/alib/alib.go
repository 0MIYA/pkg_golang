package alib

func Average(s []int) int{ //アベレージ関数？整数のスライスを渡すと整数値を返す
	total := 0             //totalを０としてfor文で回してトータルに値を追加していく処理
	for _,i:=range s{
		total += i
	}
	return int(total / len(s))
	//トータルがスライスの要素数を割った/平均値となり、整数に戻して性数値を返す
}
//main.goの方で実行処理を記載