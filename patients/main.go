package main

import (
	"github.com/newtoallofthis123/patients/api"
	"github.com/newtoallofthis123/patients/utils"
)

func main() {
	env := utils.ReadEnv()
	api, err := api.NewApiServer(&env)
	if err != nil {
		panic(err)
	}

	api.Run()
}
