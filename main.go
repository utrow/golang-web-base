package main

import (
	"github.com/utrow/golang-web-base/interface/server"
	"log"
)

func main() {
	err := server.Launch()
	if err != nil {
		log.Fatalln("[ERR] サーバの起動に失敗しました")
	}
}
