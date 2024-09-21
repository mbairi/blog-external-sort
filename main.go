package main

import (
	"blog-external-sort/src"
	"fmt"
	"strings"
	"time"
)

func main() {
	// faker.GenerateMockFile(5_000_000)
	start := time.Now()
	frame, _ := src.NewChunkyFrame("mock.jsonl", 20_000, "temp")
	frame.Sort(func(a map[string]interface{}, b map[string]interface{}) int {
		return strings.Compare(a["Name"].(string), b["Name"].(string))
	})
	fmt.Println(time.Since(start))
}
