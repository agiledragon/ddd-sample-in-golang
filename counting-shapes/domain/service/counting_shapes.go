package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/counting-shapes/domain/model"
   // "fmt"
    "fmt"
)

func CountingTriangles(points model.Points, lines []model.Line) []model.Points {
    sets := model.Subset(points, 3)
    matches := make([]model.Points, 0)
    for _, set := range sets {
        if model.IsTriangle(set, lines) {
            matches = append(matches, set)
        }
    }
    fmt.Println("matches:", matches)
    return matches
}

func CountingQuadrangles(points model.Points, lines []model.Line) []model.Points {
    sets := model.Subset(points, 4)
    matches := make([]model.Points, 0)
    for _, set := range sets {
        a := set[0]
        b := set[1]
        c := set[2]
        d := set[3]
        orderSets := []model.Points{
            model.Points([]model.Point{a, b, c, d}),
            model.Points([]model.Point{a, b, d, c}),
            model.Points([]model.Point{a, c, b, d}),
            model.Points([]model.Point{a, c, d, b}),
            model.Points([]model.Point{a, d, b, c}),
            model.Points([]model.Point{a, d, c, b}),
        }
        
        for _, orderSet := range orderSets {
            if model.IsQuadrangle(orderSet, lines) {
                matches = append(matches, orderSet)
                break
            }
        }
    
    }
    //fmt.Println("matches:", matches)
    return matches
}