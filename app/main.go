package main

import (
	"sns-sample/src/infrastructure"
)

func main() {
	db := infrastructure.NewDB()

	r := infrastructure.NewRouting(db)

	r.Run(r.Port)
}
