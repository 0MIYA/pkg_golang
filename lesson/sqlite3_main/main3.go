//?マークを＄に変更する　例＄1,$2 sqllite3とほぼ同じ内容
package main

import (
	"database/sql"
	"log"
	//"fmt"//R

	_ "github.com/lib/pq"
)

var Db *sql.DB//DBを指すポインタ変数としてdbを宣言し初期化、main関数に入る

var err error

type Person struct{
	Name string
	Age int
}

func main(){
	Db, err = sql.Open("postgres","user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil{
		log.Panicln(err)
	}
	defer Db.Close()

/*
	//データ追加　CRUDO処理
//C
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	//?,?はレコードを作成する時にカラムに該当する値を置換してくれる
	_, err := Db.Exec(cmd, "Fneko", 10)//_, err := Db.Exec(cmd,"MIYA",23)に続きデータをもう１つ追加
	//エラーハンドリング
	if err != nil{
		log.Fatalln(err)
	}
}
	
*/
	//特定のデータの取得
//R
/*
	cmd := "SELECT * FROM persons where age = $1"
	//whereで取得するデータを指定
	//QueryRow　1レコードを取できる
	row := Db.QueryRow(cmd, 10)
	var p Person //上記でpersonというstructを作りデータを入れて取得している
	err = row.Scan(&p.Name, &p.Age)//ポインタpにNameフィールド、Ageフィールドで取得したデータを格納が可能になる
	if err != nil {
		//データがなかったら
			if err == sql.ErrNoRows {
				log.Println("No row")
				//それ以外のエラー
			} else {
				log.Println(err)//それ以外のエラー
			}
		}
		fmt.Println(p.Name, p.Age)//変数pにrowで取得したデータがScanでフィールドに代入されているので、データを確認

	cmd = "SELECT * FROM persons"
	//Query は条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	//struxt作成
	var pp []Person
	//取得したデータをループでスライスに追加
	for rows.Next() {
		var p Person
		//scan データ追加
		err := rows.Scan(&p.Name, &p.Age)
		//１つずつエラーハンドリング
		if err != nil {
			log.Println(err)
	}
	pp = append(pp, p)
}
//まとめてエラーハンドリング
err = rows.Err()
if err != nil {
	log.Fatalln(err)
}
//表示
for _, p := range pp {
	fmt.Println(p.Name, p.Age)
}
}
*/

/*
//U
	//データアップデート
	cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	_, err := Db.Exec (cmd, 25, "Fneko")//25に更新
	if err != nil {
		log.Fatalln(err)
	}
}
*/


//D
	//データの削除
	cmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(cmd,"Fneko")
	if err != nil{
		log.Fatalln(err)
	}
}
	