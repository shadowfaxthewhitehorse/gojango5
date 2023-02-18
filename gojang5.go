// implementation of solution of Convex Hull problem
//
// uses Graham Scan algorithm

package main

import (
    "fmt"
    "sort"
)

type Point struct {
    x, y float64
}

func crossProduct(p, q, r Point) float64 {
    return (q.x-p.x)*(r.y-p.y) - (q.y-p.y)*(r.x-p.x)
}

func grahamScan(points []Point) []Point {
    if len(points) < 3 {
        return points
    }

    var hull []Point

    // Find the point with the smallest y-coordinate (ties broken by smallest x-coordinate)
    var anchor Point = points[0]
    for _, p := range points {
        if p.y < anchor.y || (p.y == anchor.y && p.x < anchor.x) {
            anchor = p
        }
    }

    // Sort the points by polar angle with respect to the anchor point
    sortByPolarAngle := func(p, q Point) bool {
        cp := crossProduct(anchor, p, q)
        if cp == 0 {
            return p.x < q.x
        }
        return cp > 0
    }
    sort.Slice(points, func(i, j int) bool {
        return sortByPolarAngle(points[i], points[j])
    })

    hull = append(hull, points[0], points[1])

    for i := 2; i < len(points); i++ {
        for len(hull) >= 2 && crossProduct(hull[len(hull)-2], hull[len(hull)-1], points[i]) <= 0 {
            hull = hull[:len(hull)-1]
        }
        hull = append(hull, points[i])
    }

    return hull
}

func main() {
    points := []Point{
        {0, 3},
        {1, 1},
        {2, 2},
        {4, 4},
        {0, 0},
        {1, 2},
        {3, 1},
        {3, 3},
    }

    hull := grahamScan(points)
    fmt.Println("Convex Hull:")
    for _, p := range hull {
        fmt.Printf("(%v, %v)\n", p.x, p.y)
    }
}
