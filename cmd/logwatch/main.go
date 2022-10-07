package main

import (
	"fmt"

	"github.com/paveldanilin/logwatch/service"
)

func main() {
	w := service.NewTailFileWatcher("./test.log")
	w.Watch(func(line string) {
		fmt.Println(line)
	})
}
