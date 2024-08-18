package main

import (
	"log"
	"quantum/ui"
)

/*
 * This is the main entry point of the application.
 * @Author: AkiChan
 * @Version: 1.0.0
 */
func main() {
	log.Default().SetFlags(0)
	ui.Start()
}
