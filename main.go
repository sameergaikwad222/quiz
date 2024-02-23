package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv","problem.csv","csv file for question and answer")
	timeLimit := flag.Int("timelimit",30,"the time limit for quiz in seconds")
	flag.Parse()
	file,err := os.Open(*csvFileName) 
	handleError(err,"Error Occured while reading file...")
	r := csv.NewReader(file)
	lines, err := r.ReadAll() 
	handleError(err,"Error while reading file")
	problems := parseLines(lines)
	timers := time.NewTimer(time.Duration(*timeLimit)* time.Second)
	correctAns := 0
	problemLoop:
	for i,prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i + 1, prob.q)
		ansChan := make(chan string)
		go func (){
			var answer string
			fmt.Scanf("%s\n",&answer)
			ansChan <- answer			
		}()
		select {
			case <-timers.C:
				break problemLoop
				
			case ans:= <- ansChan:
				if ans == prob.a { 
					correctAns++
				}
		}
	}
	fmt.Printf("\n\nYou have answered %v successfull answers out of %v questions",correctAns,len(problems))	
}

func parseLines(lines [][]string) [] problem { 
	var ret []problem
	for _, val := range lines {
		ret = append(ret, problem{
			q:val[0],
			a:strings.TrimSpace(val[1]),
		}) 
	}
	return ret
}

type problem struct {
	q string
	a string
}


func handleError(err error,msg string){
	if err != nil {
		fmt.Printf(msg,err)
		os.Exit(1)
	}
}