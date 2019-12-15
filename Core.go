package main

import(
	"regexp"
)

func check(e error) {
    if e != nil{
        panic(e)
    }
}

func find(slice []int, val int)(bool){
    for _, item := range slice{
        if item == val{
            return true
        }
    }
    return false
}

func isNum(s string)(bool){
	return regexp.MustCompile("^[0-9]+$").MatchString(s)
}