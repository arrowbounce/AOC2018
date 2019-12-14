package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil{
        panic(e)
    }
}

func Find(slice []int, val int) (bool){
    for _, item := range slice {
        if item == val{
            return true
        }
    }
    return false
}

func main() {
    // dat, err := ioutil.ReadFile("Day1a.txt")
    file, err := os.Open("Day1.txt")
    check(err)
    x := 0
    var y string
    var s []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        y = scanner.Text()
        v, err := strconv.Atoi(y)
        check(err)
        x += v
        s = append(s, v)
    }
    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }
    fmt.Println(x)
    var l []int
    a := true
    x = 0
    for a {
        for _,v := range s{
            x += v
            if Find(l, x){
                fmt.Println(x)
                a = false
                break
            }
            l = append(l, x)
        }
    }
}