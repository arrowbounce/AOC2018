package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

func main(){
	file, err := os.Open("Day5.txt")
    check(err)
    reader := bufio.NewReader(file)
    s := ""
    for{
    	r, _, err := reader.ReadRune()
    	if err == io.EOF{
    		break
    	}
    	if len(s) != 0{
	    	hold := string(s[len(s)-1])
	    	hold = strings.ToUpper(hold)
	    	n := strings.ToUpper(string(r))
	    	if hold == n && string(s[len(s)-1]) != string(r){
	    		s = s[:len(s)-1]
	    	}else{
	    		s = s + string(r)
	    	}
	    } else {
	    	s = s + string(r)
	    }
    }
    fmt.Println(len(s))
    alpha := "abcdefghijklmnopqrstuvwxyz"
    maxlen := 20000
    for _, l := range(alpha){
    	file.Seek(0,0)
    	reader = bufio.NewReader(file)
    	s = ""
    	for{
	    	r, _, err := reader.ReadRune()
	    	if err == io.EOF{
	    		break
	    	}
	    	if strings.ToUpper(string(r)) != strings.ToUpper(string(l)){
		    	if len(s) != 0{
			    	hold := string(s[len(s)-1])
			    	hold = strings.ToUpper(hold)
			    	n := strings.ToUpper(string(r))
			    	if hold == n && string(s[len(s)-1]) != string(r){
			    		s = s[:len(s)-1]
			    	}else{
			    		s = s + string(r)
			    	}
			    } else {
			    	s = s + string(r)
			    }
			}
	    }
	    if len(s) < maxlen{
	    	maxlen = len(s)
	    }
    }
    fmt.Println(maxlen)
    file.Close()
}