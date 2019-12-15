package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    sc "strconv"
)

func comptimes(a []int, b[]int)(bool){
    for n, t := range a{
        if t < b[n]{
            return false
        }
        if t > b[n]{
            return true
        }
    }
    return true
}

func sortimes(a[][]int)([][]int){
    l := len(a)
    if l <= 1{
        return a
    }
    b := sortimes(a[:l/2])
    c := sortimes(a[l/2:])
    i := 0
    j := 0
    var d[][]int
    for {
        if !comptimes(b[i], c[j]) {
            d = append(d, b[i])
            i += 1
        } else{
            d = append(d, c[j])
            j += 1
        }
        if i == len(b){
            d = append(d, c[j:]...)
            break
        }
        if j == len(c){
            d = append(d, b[i:]...)
            break
        }
    }
    return d
}

func main(){
    file, err := os.Open("Day4.txt")
    check(err)
    scanner := bufio.NewScanner(file)
    f := regexp.MustCompile("[^\\[\\-\\]: #]+")
    var checkins [][]int
    var wakes [][]int
    var sleeps [][]int
    var cur[]int
    for scanner.Scan(){
        s := f.FindAllString(scanner.Text(), -1)
        if isNum(s[6]){
            cur = nil
            a, _ := sc.Atoi(s[1])
            b, _ := sc.Atoi(s[2])
            c, _ := sc.Atoi(s[3])
            d, _ := sc.Atoi(s[4])
            e, _ := sc.Atoi(s[6])
            cur = []int{a, b, c, d, e}
            checkins = append(checkins, cur)
        } else {
            a, _ := sc.Atoi(s[1])
            b, _ := sc.Atoi(s[2])
            c, _ := sc.Atoi(s[3])
            d, _ := sc.Atoi(s[4])
            cur = []int{a, b, c, d}
            if s[6] == "up"{
                wakes = append(wakes, cur)
            } else{
                sleeps = append(sleeps, cur)
            }
        }
    }
    sleeps = sortimes(sleeps)
    wakes = sortimes(wakes)
    checkins = sortimes(checkins)
    drivers := make(map[int]int)
    j := 0
    curdriver := checkins[j][4]
    drivers[curdriver] = 0
    for i := 0; i < len(sleeps); i++{
        for {
            if j != len(checkins) - 1{
                if comptimes(sleeps[i], checkins[j+1]){
                    j++
                } else{
                    break
                }
            } else{
                break
            }
        }
        curdriver = checkins[j][4]
        drivers[curdriver] += wakes[i][3] - sleeps[i][3]
    }
    maxdriver := 0
    maxmins := 0
    for k, v := range drivers {
        if v > maxmins{
            maxdriver = k
            maxmins = v
        }
    }
    mins := make([]int, 60)
    j = 0
    for i := 0; i < len(sleeps); i++{
        for {
            if j != len(checkins) - 1{
                if comptimes(sleeps[i], checkins[j+1]){
                    j++
                } else{
                    break
                }
            } else{
                break
            }
        }
        curdriver = checkins[j][4]
        if curdriver != maxdriver {
            continue
        }
        for k := sleeps[i][3]; k < wakes[i][3]; k++{
            mins[k] += 1
        }
    }
    maxmins = 0
    max := 0
    for k, v := range mins {
        if v > maxmins{
            maxmins = v
            max = k
        }
    }
    fmt.Println(max, maxdriver, max * maxdriver)
    // fmt.Println(drivers, maxdriver, mins)
    maxmins = 0
    max = 0
    min := 0
    for v, _ := range drivers{
        mins := make([]int, 60)
        j = 0
        for i := 0; i < len(sleeps); i++{
            for {
                if j != len(checkins) - 1{
                    if comptimes(sleeps[i], checkins[j+1]){
                        j++
                    } else{
                        break
                    }
                } else{
                    break
                }
            }
            curdriver = checkins[j][4]
            if curdriver != v {
                continue
            }
            for k := sleeps[i][3]; k < wakes[i][3]; k++{
                mins[k] += 1
            }
        }
        for k, m := range mins {
            if m > maxmins{
                maxmins = m
                max = v
                min = k
            }
        }
    }
    fmt.Println(max, maxmins, min, max * min)
}