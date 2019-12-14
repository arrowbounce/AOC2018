package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
    "strings"
    "strconv"
)

type col struct {
    place int
    row []int
}

func interp(s []string) (int, int, int, int) {
    a, _ := strconv.Atoi(s[2])
    b, _ := strconv.Atoi(s[3])
    c, _ := strconv.Atoi(s[4])
    d, _ := strconv.Atoi(s[5])
    return a, b, c, d
}

func main(){
    file, err := os.Open("Day3.txt")
    check(err)
    scanner := bufio.NewScanner(file)
    f := regexp.MustCompile("[^#@,:x ]*")
    cols := make(map[int]*col)
    overlaps := make(map[int]*col)
    for scanner.Scan(){
        sub := strings.Replace(scanner.Text(), " ", "", -1)
        s := f.FindAllString(sub, -1)
        x, y, xdiff, ydiff := interp(s)
        // fmt.Printf("%d %d %d %d\n", x, y, xdiff, ydiff)
        for i := y; i < y + ydiff; i++{
            if _, ok := cols[i]; !ok{
                var a []int
                cols[i] = &col{i, a}
            }
            curcol := cols[i]
            for j := x; j < x + xdiff; j++{
                if !find(curcol.row, j){
                    curcol.row = append(curcol.row, j)                   
                } else {
                    if _, ok := overlaps[i]; !ok{
                        var b []int
                        overlaps[i] = &col{i, b}
                    }
                    loc := overlaps[i]
                    if !find(loc.row, j){
                        loc.row = append(loc.row, j)
                    }
                }
            }
        }
    }
    s := 0
    for _, v := range overlaps{
        s += len(v.row)
    }
    fmt.Println(s)
    file.Seek(0,0)
    scanner = bufio.NewScanner(file)
    for scanner.Scan(){
        sub := strings.Replace(scanner.Text(), " ", "", -1)
        s := f.FindAllString(sub, -1)
        x, y, xdiff, ydiff := interp(s)
        // fmt.Printf("%d %d %d %d\n", x, y, xdiff, ydiff)
        laps := false
        for i := y; i < y + ydiff; i++{
            if _,ok := overlaps[i]; ok{
                curcol := overlaps[i]
                for j := x; j < x + xdiff; j++{
                    if find(curcol.row, j){
                        laps = true
                    }
                }
            }
        }
        if !laps{
            fmt.Println(s[1])            
        }
    }
}