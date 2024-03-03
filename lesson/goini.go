//config.iniの接続
package main 

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {//読み込んだ値をプログラム上で使用する
	Port      int
	DbName    string
	SQLDriver string
}

//main関数の前に実行するコンフィグリストを作成
var Config ConfigList//stractをグローバルに宣言してコンフィグ

func init(){
	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080),
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

//main関数でコンフィグリストのフィールドを読み込み表示する
func main(){
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)
}
