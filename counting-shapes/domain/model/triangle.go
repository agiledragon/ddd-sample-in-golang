package model

func IsTriangle(points Points, lines []Line) bool {
    connected := hasConnected(lines)
    in_same_line := inSameLine(lines)
    p0 := Point(points[0])
    p1 := Point(points[1])
    p2 := Point(points[2])
    if connected(p0, p1) &&
        connected(p1, p2) &&
        connected(p0, p2) &&
        not(in_same_line(p0, p1, p2)) {
        return true
    }
    return false
}

