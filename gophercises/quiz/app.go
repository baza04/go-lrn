package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	srcFile := flag.String("source", "problem.csv", "sourse file to quiz questions")
	tickTime := flag.Int("timer", 30, "time to each question solving")
	flag.Parse()

	ticker := time.NewTicker(*tickTime * time.Second)
	defer ticker.Stop()

	file, err := os.Open("./" + *srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answers := make(map[string]int)
	line := make([]string, 0, 2)
	ansChan := make(chan string)
	answer := ""

	for scanner.Scan() {
		line = strings.Split(scanner.Text(), ",")
		fmt.Print(line[0], "=")
		go func(ansChan chan string, answer string) {
			fmt.Scan(&answer)
			ansChan <- answer
		}(ansChan, answer)

		select {
		case <-ticker.C:
			continue

		case ans := <-ansChan:
			if ans == line[1] {
				answers["right"]++
			} else {
				answers["wrong"]++
			}
		}
	}
	fmt.Printf("answers:\nright -> %d\nwrong -> %d\n", answers["right"], answers["wrong"])
}
