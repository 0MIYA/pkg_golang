//uuidとはソフトウェア上のオブジェクトを一意に識別するための識別子
//例　いろんなバージョンがあるらしい
   // 6aa3c114-d911-11ee-9b10-7831c1bb6ac6
  //7ee21fda-0f4b-4973-a779-5b88e7068eb6
//uuid pkg インストール[go get "github.com/google/uuid"]

package main
import (
	"fmt"

	"github.com/google/uuid"
)

//UUID

func main(){
	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println(" ", uuidObj2.String())
}