/*importされない件：go get github.com/mattn/go-sqlite3
このコマンドを実行すると、go-sqlite3 パッケージが 
GitHub リポジトリから取得され、Go モジュールキャッシュにインストールされます*/
//「sqlite3 example.sql」コマンドでexamplesqlの中身に入る
//.tables・.tables　テーブル一覧表示

package main
import (
	"database/sql"
	"log"
	//"fmt"

	_ "github.com/mattn/go-sqlite3"
)
/* コード内で使用しない為アンダーバー_でインポートしています。
ビルドして一緒にコンパイルしないとSQLにアクセスできない。*/

//databse + sqlite3
//テーブル作成

var Db *sql.DB//DBを指すポインタ変数としてdbを宣言し初期化、main関数に入る

type Person struct{
	Name string
	Age int
}

func main(){
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データベースの初期設定
	/*
	cmd := `CREATE TABLE IF NOT EXISTS persons( 
				name STRING,
				age INT)` 
				//option+ろ=``で囲むと改行できる
				//personsテーブルがなければpersonsテーブルを作成する

	_, err := Db.Exec(cmd)//Execでcmd内を実行、使わないので_アンダーバー
	
	if err != nil {
		log.Fatalln(err)//errを受け取ったら出力する指定
	}
	*/

	//データ追加　CRUDO処理
	/*
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	//?,?はレコードを作成する時にカラムに該当する値を置換してくれる
	_, err := Db.Exec(cmd,"Fneko",10)//_, err := Db.Exec(cmd,"MIYA",23)に続きデータをもう１つ追加
	//エラーハンドリング
	if err != nil{
		log.Fatalln(err)
	}
	//SELECT * FROM persons;
	*/


	//特定のデータの取得
	/*
	cmd := "SELECT * FROM persons where age = ?"
	//whereで取得するデータを指定
	row := Db.QueryRow(cmd, 23)
	//QueryRow　1レコードを取できる
	var p Person //上記でpersonというstructを作りデータを入れて取得している
	err := row.Scan(&p.Name, &p.Age)//ポインタpにNameフィールド、Ageフィールドで取得したデータを格納が可能になる
	if err != nil {
			if err == sql.ErrNoRows{
				log.Println("No row")
				} else {
					log.Println(err)//それ以外のエラー
			}
		}
		fmt.Println(p.Name, p.Age)//変数pにrowで取得したデータがScanでフィールドに代入されているので、データを確認
	*/

	//複数のデータの取得
	/*
	cmd := "SELECT * FROM persons"
	//Query は条件に合うものを全て取得
	rows, _:= Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next(){
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil{
			log.Println(err)
	}
	pp = append(pp, p)
}

//まとめてエラーハンドリング
err = roes.Err()
if err != nil{
	log.Fatalln(err)
}

//表示
for _, p := range pp{
	fmt.Println(p.Name, p.Age)
	}
*/

	//データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd,"tarou")
	if err != nil{
		log.Fatalln(err)
	}
	//SELECT * FROM persons;
}


/*SQL文サンプル
CREATE TABLE文


【SELECT文とは】
SELECT文とは先ほども伝えた通り「データベースから、データを取得する命令」です。このSELECT文をしようすることで、データベースからデータを取得することが出来ます。

【SELECT文の基本構文】
SELECT文の基本的な構文は以下になります。

SELECT [取得したい要素] FROM [使用テーブル];
欲しい情報と、どこから取ってくるかを指定すれば取ってくることが出来ます。


【特定要素だけ表示する】
基本的な構文は以下になります。

実行命令:SELECT name FROM users;
意味:「usersテーブル」から全データの「name」だけを取得


【「,」で複数指定】
「,」で区切ることで、複数の要素を取得可能です。次の例では「名前」と「レベル」のみを取得・表示しています。

実行命令:SELECT name,email FROM users;
意味:「usersテーブル」から全データの「name, email」を取得



【* で全部指定】
取得したい要素の項目に*を入れることで、全ての要素を取得することが可能です。

実行命令:SELECT * FROM users
意味:「usersテーブル」から全データの「全要素」を取得。


【SELECTの基本構文(まとめ)】
SELECT [取得したい要素] FROM [使用テーブル];
この一文だけでデータを取ってくることが可能です。まずはこれを覚えておきましょう。


【「WHERE」で条件を絞る】
WHEREで条件設定を行うことができる。使い方は簡単で、先ほどのSELECT文の後ろWHERE [条件文]を追加するだけです。

SELECT [取得したい要素] FROM [使用テーブル] WHERE [条件文];
例えば先ほどあげた例「level30以上のキャラクターだけ取得したり」を実行してみましょう。

実行命令:SELECT * FROM users WHERE name = "mike";
意味「usersテーブル」の「nameがmikeのデータ」の「全要素」を取得


【INSERT文で新規追加】
データの追加はINSERT文で行います。

INSERT INTO [データを追加したいテーブル名] (要素名A,要素名B) VALUES(要素Aに入れるデータ,要素Bに入れるデータ);
例えば今回userを増やすためにINSERTを行うなら、以下のようになる。

実行命令:INSERT INTO users (id, name) VALUES(1,"joe");
usersテーブルにid=1, name="joe"のデータを追加。


実行命令:SELECT * FROM users;
を実行すると、データが追加されているのが確認できる。


【UPDATE文で変更】
データの更新はUPDATE文を使う。

実行命令:UPDATE users SET name = "nancy" WHERE id = 1;
意味:「usersテーブル」の「id = 1のデータ」の「nameを"nancy"」に更新


【DELETE文で削除】
データの削除はDELETE文で使う。

実行命令:DELETE FROM users WHERE id = 1;
意味:「usersテーブル」の「idが１のデータ」を削除する

*/