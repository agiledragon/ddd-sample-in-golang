package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/counting-shapes/domain/service"
)

func CountingTriangles(points string, lines []string) int {
    return len(service.CountingTriangles(points, lines))
}
