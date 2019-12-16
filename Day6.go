package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
)

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}

func main(){
    file, err := os.Open("Day6.txt")
    check(err)
    scanner := bufio.NewScanner(file)
    f := regexp.MustCompile("[0-9]+")
    maxX := 0
    maxY := 0
    minX := 1000
    minY := 1000
    maxdist := 10000
    var points [][]int
    for scanner.Scan() {
        s := f.FindAllString(scanner.Text(), -1)
        v := []int{}
        a, _ := strconv.Atoi(s[0])
        v=append(v, a)
        a, _ = strconv.Atoi(s[1])
        v=append(v, a)
        points = append(points, v)
        if v[0] > maxX{
            maxX = v[0]
        }
        if v[1] > maxY{
            maxY = v[1]
        }
        if v[0] < minX{
            minX = v[0]
        }
        if v[1] < minY{
            minY = v[1]
        }
    }
    var counts[]int
    var inf[]int
    curRangeTot := 0
    for _, _ = range points{
        counts = append(counts, 0)
        inf = append(inf, 0)
    }
    fmt.Println(points, maxX, minX, maxY, minY)
    for i := minX; i <= maxX; i++{
        for j:=minY; j<=maxY; j++{
            mindist := 1000
            totdist := 0
            var minpoint int
            curatpoint := 0
            for x, point := range points{
                curdist := abs(point[0] - i) + abs(point[1] - j)
                totdist += curdist
                if curdist < mindist{
                    mindist = curdist
                    minpoint = x
                    curatpoint = 0
                } 
                if curdist == mindist{
                    curatpoint += 1
                }
            }
            if curatpoint == 1{
                counts[minpoint] += 1
            }
            if i == minX || i == maxX || j == minY || j == maxY{
                inf[minpoint] += 1
            }
            if totdist < maxdist{
                curRangeTot += 1
            }
        }
    }
    maxcount := 0
    for x, point := range counts{
        if inf[x] == 0{
            if maxcount < point{
                maxcount = point
            }
        } else {
            //counts[x] = 0
            continue
        }
    }
    fmt.Println(counts, inf, maxcount, curRangeTot)
}
