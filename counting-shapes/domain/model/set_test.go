package model

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestRingConnected(t *testing.T) {
    lines := []string{"abh", "acgi", "adfj", "aek", "bcde", "hgfe", "hijk"}
    ring_connected := ringOrderConnected(lines)
    Convey("TestRingConnected", t, func() {
        
        Convey("succ", func() {
            So(ring_connected([]Point("abc")...), ShouldBeTrue)
            So(ring_connected([]Point("bcgh")...), ShouldBeTrue)
            So(ring_connected([]Point("cdfg")...), ShouldBeTrue)
            So(ring_connected([]Point("abeg")...), ShouldBeTrue)
            So(ring_connected([]Point("abhg")...), ShouldBeTrue)
            So(ring_connected([]Point("defj")...), ShouldBeTrue)
        })
    
        Convey("failed", func() {
            So(ring_connected([]Point("acf")...), ShouldBeFalse)
            So(ring_connected([]Point("bcg")...), ShouldBeFalse)
            So(ring_connected([]Point("defk")...), ShouldBeFalse)
        })
        
    })
}

func TestHasCrossConnected(t *testing.T) {
    lines := []string{"abh", "acgi", "adfj", "aek", "bcde", "hgfe", "hijk"}
    cross_connected := hasCrossConnected(lines)
    Convey("TestHasCrossConnected", t, func() {
        
        Convey("succ", func() {
            So(cross_connected("ag", "bd"), ShouldBeTrue)
            So(cross_connected("af", "ce"), ShouldBeTrue)
            So(cross_connected("af", "be"), ShouldBeTrue)
            So(cross_connected("ai", "eh"), ShouldBeTrue)
        })
        
        Convey("failed", func() {
            So(cross_connected("bc", "hg"), ShouldBeFalse)
            So(cross_connected("fj", "ek"), ShouldBeFalse)
        })
        
    })
}

