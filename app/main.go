package main

import (
	"sns-sample/src/infrastructure"
)

func main() {
	config := infrastructure.NewConfig()
	db := infrastructure.NewDB()
	google := infrastructure.NewGoogle(config)

	r := infrastructure.NewRouting(db, google)

	r.Run(r.Port)
}
