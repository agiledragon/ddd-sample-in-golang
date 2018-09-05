package test

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
    "github.com/agiledragon/ddd-sample-in-golang/counting-shapes/app/service"
)

func TestCountingShapes(t *testing.T) {
    points := "abcdefghijk"
    lines := []string{"abh", "acgi", "adfj", "aek", "bcde", "hgfe", "hijk"}
    
    Convey("TestCountingShapes", t, func() {
        
        Convey("counting triangles", func() {
            num := service.CountingTriangles(points, lines)
            So(num, ShouldEqual, 24)
        })
    
        Convey("counting quadrangles", func() {
            num := service.CountingQuadrangles(points, lines)
            So(num, ShouldEqual, 18)
        })
    })
}

