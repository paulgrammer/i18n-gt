package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func translateText(sourceLanguage, targetLanguage string, text []string) ([]string, error) {
	ctx := context.Background()
	var translated []string

	lang, err := language.Parse(targetLanguage)
	target, err2 := language.Parse(sourceLanguage)

	if err != nil {
		return translated, fmt.Errorf("language.Parse: %v", err)
	}

	if err2 != nil {
		return translated, fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)

	if err != nil {
		return translated, err
	}

	defer client.Close()

	resp, err := client.Translate(ctx, text, lang, &translate.Options{
		Source: target,
	})

	if err != nil {
		return translated, fmt.Errorf("translate: %v", err)
	}

	if len(resp) == 0 {
		return translated, fmt.Errorf("translate returned empty response to text: %s", text)
	}

	for _, translation := range resp {
		translated = append(translated, translation.Text)
	}

	return translated, nil
}

type content map[string]interface{}

func readJson(path string) content {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	data := make(content)
	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		panic(err)
	}

	return data
}

// get returns the value of a nested key in the map using dot notation.
func (c content) get(query string) interface{} {
	keys := strings.Split(query, ".")
	for _, key := range keys {
		value, ok := c[key]
		if !ok {
			return nil
		}
		c, ok = value.(map[string]interface{})
		if !ok {
			return value
		}
	}

	return c
}

// set sets the value of a nested key in the map using dot notation.
func (c content) set(query string, value interface{}) {
	keys := strings.Split(query, ".")
	lastKey := keys[len(keys)-1]
	for _, key := range keys[:len(keys)-1] {
		subData, ok := c[key].(map[string]interface{})
		if !ok {
			subData = make(map[string]interface{})
			c[key] = subData
		}
		c = subData
	}

	c[lastKey] = value
}

// keys returns a slice of all keys in the map, with nested keys represented using dot notation.
func (c content) keys() []string {
	return c.getSubKeys("", c)
}

// getSubKeys recursively gets all keys in a nested map, with nested keys represented using dot notation.
func (c content) getSubKeys(prefix string, subMap map[string]interface{}) []string {
	var keys []string

	for key, value := range subMap {
		fullKey := fmt.Sprintf("%s%s", prefix, key)
		if subSubMap, ok := value.(map[string]interface{}); ok {
			subKeys := c.getSubKeys(fmt.Sprintf("%s%s", fullKey, "."), subSubMap)
			keys = append(keys, subKeys...)
		} else {
			keys = append(keys, fullKey)
		}
	}

	return keys
}

// save to json file
func (c content) save(path string) {
	// Convert the map to JSON
	rawJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}

	// Replace the Unicode escape sequences with the corresponding characters
	jsonString := string(rawJSON)
	re := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	jsonString = re.ReplaceAllStringFunc(jsonString, func(m string) string {
		u, _ := strconv.ParseInt(m[2:], 16, 32)
		return string(rune(u))
	})

	// Write the JSON to a file
	err = os.WriteFile(path, []byte(jsonString), 0644)
	if err != nil {
		panic(err)
	}
}

// splitIntoArrays returns 2D Array given maxlength for each item
func splitIntoArrays(strs []string, maxLen int) [][]string {
	var result [][]string
	for i := 0; i < len(strs); i += maxLen {
		end := i + maxLen
		if end > len(strs) {
			end = len(strs)
		}
		result = append(result, strs[i:end])
	}
	return result
}

func main() {
	// Parse command-line arguments
	inputFile := flag.String("input", "", "Path to the input JSON file")
	sourceLanguage := flag.String("source", "en", "Source language code")
	targetLanguage := flag.String("target", "", "Target language code")
	outputFile := flag.String("output", "output.json", "Path to the output JSON file")
	flag.Parse()

	// Check that all required arguments are provided
	if *inputFile == "" || *targetLanguage == "" {
		fmt.Println("Input file or target language missing.")
		os.Exit(1)
	}

	input := readJson(*inputFile)
	output := make(content)

	// type.googleapis.com/google.rpc.BadRequest
	// Number of text queries >  Maximum: 128
	splitted := splitIntoArrays(input.keys(), 128)

	for i, keys := range splitted {
		var values []string

		for _, key := range keys {
			if str, ok := input.get(key).(string); ok {
				values = append(values, str)
			}
		}

		text, err := translateText(*sourceLanguage, *targetLanguage, values)

		if err != nil {
			panic(err)
		}

		for i, key := range splitted[i] {
			output.set(key, text[i])
		}
	}

	// save output
	output.save(*outputFile)
}
