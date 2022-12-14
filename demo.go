package main

import (
	"flag"
	"os"
	"server/internal/logs"
)

func demo() {
	exe, _ := os.Executable()

	logFlag := flag.String("log", "stdFlags", "log flags")

	flag.Parse()

	logs.Init(os.Stdout, *logFlag)

	logs.Info.Println(exe, "started")
}
