package main

import (
    "fmt"
    "bufio"
    "os"
)

func check(e error) {
    if e != nil{
        panic(e)
    }
}

func counts(s string) (bool, bool){
	m := make(map[rune]int)
	for _, char := range s {
		if _, ok := m[char]; ok{
			m[char] += 1
		} else {
			m[char] = 1
		}
	}
	two := false
	three := false
	for _, v := range m {
		if v == 2{
			two = true
		}
		if v == 3{
			three = true
		}
	}
	return two, three
}

func comp(s string, t string) (bool, string){
	var hold string
	errs := 0
	for i, v := range s {
		if s[i] == t[i] {
			hold += string(v)
		} else {
			errs += 1
		}
		if errs == 2{
			return false, hold
		}
	}
	return true, hold
}

func main(){
	file, err := os.Open("Day2.txt")
    check(err)
    scanner := bufio.NewScanner(file)
    var y string
    twos := 0
    threes := 0
    var c []string
    for scanner.Scan(){
        y = scanner.Text()
        c = append(c, y)
        c1, c2 := counts(y)
        if c1{
        	twos += 1
        }
        if c2{
        	threes += 1
        }
    }
    fmt.Println(twos * threes)
    for _, a := range c {
    	for _, b := range c{
    		if a != b {
    			check, res := comp(a, b)
    			if check{
    				fmt.Println(res)
    			}
    		}
    	}
    }
}