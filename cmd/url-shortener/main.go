package main

import (
	"GoProjects/massive/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)


	// TODO: init config: cleanenv

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi

	// TODO: run server
}