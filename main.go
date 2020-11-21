package main

import "github.com/tl-dream-team/todo-backend/router"

func main() {
	r := router.Router()

	r.Run()
}
