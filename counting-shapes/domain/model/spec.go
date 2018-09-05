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
    ab := Points([]Point{a, b})
    cd := Points([]Point{c, d})
    ad := Points([]Point{a, d})
    bc := Points([]Point{b, c})
    
    return ring_order_connected(a, b, c, d) &&
        not(cross_connected(ab, cd)) &&
        not(cross_connected(ad, bc)) &&
        not(in_same_line(a, b, c)) &&
        not(in_same_line(a, b, d)) &&
        not(in_same_line(a, c, d)) &&
        not(in_same_line(b, c, d))
}

