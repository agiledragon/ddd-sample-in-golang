package model

import (
    "strings"
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