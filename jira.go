package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func requireEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Missing required environment variable %s", key))
	}

	return value
}

func unique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}

	return uniqSlice
}

func main() {
	baseUrl := requireEnv("JIRA_BASE_URL")
	prefixes := strings.Split(requireEnv("JIRA_PREFIXES"), ",")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		prefixesRegex := strings.Join(prefixes, "|")
		regex := regexp.MustCompile(`(` + prefixesRegex + `)-?(\d+)`)
		matches := unique(regex.FindAllString(input, -1))

		for _, match := range matches {
			fmt.Println(baseUrl + "/browse/" + regex.ReplaceAllString(match, "$1-$2"))
		}
	}
}
