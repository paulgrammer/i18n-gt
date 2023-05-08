package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"github.com/joho/godotenv"
	"golang.org/x/text/language"
)

func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)

	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)

	if err != nil {
		return "", err
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)

	if err != nil {
		return "", fmt.Errorf("translate: %v", err)
	}

	if len(resp) == 0 {
		return "", fmt.Errorf("translate returned empty response to text: %s", text)
	}

	return resp[0].Text, nil
}

func main() {
	godotenv.Load()

	if text, err := translateText("pl", "Hello world"); err == nil {
		fmt.Println(text)
	}
}
