package main

import (
    "bufio"
    "fmt"
    "flag"
    "strings"
    // "reflect"
    // "io"
    // "io/ioutil"
    "os"
)

var (
	in = flag.String("in", "", "Specifies the input filename.")
)

type repeat struct{
	n string //name
	l int //length
	m string //match
	o string //orientation
}

func main() {
	
	// Executes command-line parsing
	flag.Parse()
	if *in == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	in := *in + "_Censor.fasta.map"

	repeats := createRepeatList(in)	

	fmt.Println(repeats)

}



func check(e error) {
    if e != nil {
        panic(e)
    }
}


func createRepeatList(in string) []string{
	// open input file in read-only mode
	f, err := os.Open(in)
	check(err)
	defer f.Close()
	
	// define repe
	var repeats []string

	//scanning
	sc:=bufio.NewScanner(f)
	for sc.Scan(){
		
		// Convert input into a slice
		fields := strings.Fields(sc.Text())

		//repeats
		repeats = append(repeats, fields[0])
	}
	//remove duplicate repeats
	repeats = removeDuplicates(repeats)

	return repeats
}


//adapted from <https://groups.google.com/forum/#!topic/golang-nuts/-pqkICuokio>
func removeDuplicates(a []string) []string { 
        result := []string{} 
        seen := map[string]string{} 
        for _, val := range a { 
                if _, ok := seen[val]; !ok { 
                        result = append(result, val) 
                        seen[val] = val 
                } 
        } 
        return result 
} 