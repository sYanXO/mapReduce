package main
import (
	"bufio"
	"os"
	"fmt"
	"log"
	"strings"
	"sort"
)
type Result struct {
    Username string
    Count    int
}
func main(){
	file,err := os.Open("input.txt")
	
	if err!=nil{
		log.Fatal(err)
	}
	
	os.RemoveAll("intermediate")
	os.MkdirAll("intermediate", 0755)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		parts:= strings.Split(scanner.Text(),"|")
		username:= strings.TrimSpace(parts[0])
		
		outFile, err := os.OpenFile("intermediate/"+username+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
    		log.Fatal(err)
		}
		outFile.WriteString(username + ",1\n")
		outFile.Close()
	}

	if err:= scanner.Err();err!=nil{
		log.Fatal(err)
	}

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