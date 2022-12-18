package main

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	OpenAIApiKey string
}

func load() (*Config, error) {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	openAIApiKey, ok := viper.Get("OPENAI_API_KEY").(string)
	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		return nil, errors.New("failed loading OpenAI API key")
	}

	return &Config{
		OpenAIApiKey: openAIApiKey,
	}, nil
}
