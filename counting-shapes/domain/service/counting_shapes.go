package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/counting-shapes/domain/model"
)

func CountingTriangles(points model.Points, lines []model.Line) []model.Points {
    sets := model.Subset(points, 3)
    matches := make([]model.Points, 0)
    for _, set := range sets {
        if model.IsTriangle(set, lines) {
            matches = append(matches, set)
        }
    }
    return matches
}
