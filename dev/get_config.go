package dev

import "github.com/joho/godotenv"

func SetConfig() error {
	err := godotenv.Load("./dev/local.env")
	if err != nil {
		return err
	}

	return nil
}
