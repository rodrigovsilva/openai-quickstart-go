package main

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	cfg, err := load()
	if err != nil {
		panic("failed loading configuration")
	}

	c := gogpt.NewClient(cfg.OpenAIApiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-002",
		Prompt:      generatePrompt("Elephant"),
		Temperature: 0.6,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}

	fmt.Println(resp.Choices[0].Text)
}

func generatePrompt(text string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	capitalized := caser.String(text)

	return fmt.Sprintf(`Suggest three names for an animal that is a superhero.
Animal: Cat
Names: Captain Sharpclaw, Agent Fluffball, The Incredible Feline
Animal: Dog
Names: Ruff the Protector, Wonder Canine, Sir Barks-a-Lot
Animal: %s
Names:`, capitalized)
}
