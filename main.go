package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

const numWorkers = 4

type Result struct {
	Username string
	Count    int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.RemoveAll("intermediate")
	os.MkdirAll("intermediate", 0755)

	// Read all lines into memory
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Split lines into chunks for each worker
	chunkSize := (len(lines) + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup
	var mu sync.Mutex // protects concurrent file writes

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		if start >= len(lines) {
			break
		}

		wg.Add(1)
		go func(chunk []string) {
			defer wg.Done()
			for _, line := range chunk {
				parts := strings.Split(line, "|")
				username := strings.TrimSpace(parts[0])

				mu.Lock()
				outFile, err := os.OpenFile("intermediate/"+username+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					mu.Unlock()
					log.Fatal(err)
				}
				outFile.WriteString(username + ",1\n")
				outFile.Close()
				mu.Unlock()
			}
		}(lines[start:end])
	}

	wg.Wait()

	// reading frm /intermediate
	results := []Result{}
	entries,err := os.ReadDir("intermediate")
	if(err!=nil){
		log.Fatal(err)
	}
	for _,entry := range entries{
		filepath := "intermediate/"+entry.Name()
		f,err:= os.Open(filepath)
		if(err!=nil){
			log.Fatal(err)
		}
		count:=0
		s:=bufio.NewScanner(f)
		for s.Scan(){
			count++
		}
		f.Close()

		username := strings.Replace(entry.Name(), ".txt", "", 1)
		

		results = append(results, Result{Username: username, Count: count})
	}

	sort.Slice(results,func(i,j int)bool{
		return results[i].Count > results[j].Count
	})

	for i,r := range results{
		if i >= 10{
			break
		}
		fmt.Println(r.Username,":",r.Count)
	}


	
}