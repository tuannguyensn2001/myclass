package main

import (
	"context"
	"log"
	"myclass/src/cmd"
	"myclass/src/config"
	"myclass/src/models"
	logpkg "myclass/src/packages/log"
	"myclass/src/services/user/userrepository"
	"time"
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

	repo := userrepository.New()
	user := models.User{
		Username:  "tuannguyensn2001",
		Password:  "java2001",
		Email:     "tuannguyensn2001a@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repo.Insert(context.Background(), &user)

	rootCmd := cmd.Root()
	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

}
