package tool

import "github.com/joho/godotenv"

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
