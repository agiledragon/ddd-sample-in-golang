package model

import (
    "strings"
    //"fmt"
)

type Point = byte
type Line = string
type Points = string


func belong(points Points, lines []Line) bool {
    flag := false
    for _, line := range lines {
        flag = true
        for _, point := range points {
            if !strings.ContainsRune(line, point) {
                flag = false
                break
            }
        }
        if flag {
            return true
        }
    }
    return false
}

func Subset(points Points, n int) []Points {
    l := len(points)
    if l < n {
        return nil
    }
    if l == n {
        return []string{points}
    }
    
    results := make([]Points, 0)
    if n == 1 {
        for i := range points {
            results = append(results, Points([]byte{points[i]}))
        }
        return results
    }
    
    firsts := Subset(points[1:], n - 1)
    for _, first := range firsts {
        results = append(results, Points([]byte{points[0]}) + first)
    }
    
    lasts := Subset(points[1:], n)
    results = append(results, lasts...)
    
    return results
}

func hasConnected(lines []Line) func(x, y Point) bool {
    return func(x, y Point) bool {
        points := make([]Point, 2)
        points[0] = x
        points[1] = y
        return belong(Points(points), lines)
    }
}

func inSameLine(lines []Line) func(x, y, z Point) bool {
    return func(x, y, z Point) bool {
        points := make([]Point, 3)
        points[0] = x
        points[1] = y
        points[2] = z
        return belong(Points(points), lines)
    }
}

func not(flag bool) bool {
    return !flag
}

func ringOrderConnected(lines []Line) func(points ...Point) bool {
    connected := hasConnected(lines)
    
    return func(points ...Point) bool {
        //  fmt.Println("ring:", points)
        prev := points[0]
        for i := 1; i < len(points); i++ {
            next := points[i]
            if !connected(prev, next) {
                return false
            }
            prev = next
        }
        first := points[0]
        last := points[len(points) - 1]
        if connected(last, first) {
            return true
        }
        return false
    }
}

func hasCrossConnected(lines []Line) func(ls, rs Points) bool {
    return func(ls, rs Points) bool {
        flag := false
        lfound := false
        rfound := false
        ll := ""
        rl := ""
        for _, line := range lines {
            if !lfound {
                flag = true
                for _, l := range ls {
                    if !strings.ContainsRune(line, l) {
                        flag = false
                        break
                    }
                }
                if flag {
                    ll = line
                    lfound = true
                }
            }
    
            if !rfound {
                flag = true
                for _, r := range rs {
                    if !strings.ContainsRune(line, r) {
                        flag = false
                        break
                    }
                }
                if flag {
                    rl = line
                    rfound = true
                }
            }
            
            if lfound && rfound {
                break
            }
            
        }
        
        //fmt.Println("ll:", ll)
        //fmt.Println("rl:", rl)
        
        lstart := 0
        lend := 0
        for i := range ll {
            if ll[i] == ls[0] {
                lstart = i
            } else if ll[i] == ls[1] {
                lend = i
            }
        }
        if lstart > lend {
            lstart, lend = lend, lstart
        }
        //fmt.Println("lstart, lend:", lstart, lend)
        
    
        rstart := 0
        rend := 0
        for i := range rl {
            if rl[i] == rs[0] {
                rstart = i
            } else if rl[i] == rs[1] {
                rend = i
            }
        }
        if rstart > rend {
            rstart, rend = rend, rstart
        }
        //fmt.Println("rstart, rend:", rstart, rend)
        for i := 1; i < lend - lstart; i++ {
            for j := 1; j < rend - rstart; j++ {
                //fmt.Println("ll[i]:", ll[lstart + i])
                //fmt.Println("rl[j]:", rl[rstart + j])
                if ll[lstart + i] == rl[rstart + j] {
                    return true
                }
            }
        }
        return false
    }
}