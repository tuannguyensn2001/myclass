package main

import (
	"log"
	"myclass/src/cmd"
	"myclass/src/config"
	logpkg "myclass/src/packages/log"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	err = logpkg.Init()
	defer func() {
		logpkg.SyncLog()
	}()
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd := cmd.Root()
	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

}
