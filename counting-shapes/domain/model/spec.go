package model

func IsTriangle(points Points, lines []Line) bool {
    ring_order_connected := ringOrderConnected(lines)
    in_same_line := inSameLine(lines)
    
    a := points[0]
    b := points[1]
    c := points[2]
    
    return ring_order_connected(a, b, c) &&
        not(in_same_line(a, b, c))
}

func IsQuadrangle(points Points, lines []Line) bool {
    ring_order_connected := ringOrderConnected(lines)
    cross_connected := hasCrossConnected(lines)
    in_same_line := inSameLine(lines)
    
    a := points[0]
    b := points[1]
    c := points[2]
    d := points[3]
    
    ring_order_connected_without_cross := func(w, x, y, z Point) bool {
        wx := Points([]Point{w, x})
        yz := Points([]Point{y, z})
        wz := Points([]Point{w, z})
        xy := Points([]Point{x, y})
        return ring_order_connected(w, x, y, z) &&
            not(cross_connected(wx, yz)) &&
            not(cross_connected(wz, xy))
        
    }
    
    return ring_order_connected_without_cross(a, b, c, d) &&
        not(in_same_line(a, b, c)) &&
        not(in_same_line(a, b, d)) &&
        not(in_same_line(a, c, d)) &&
        not(in_same_line(b, c, d))
}

