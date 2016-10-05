package main

import (
	"./convexhull"
)

const (
	Title  = "Convex Hull in 2D"
	Width  = 840
	Height = 630
	HW     = Width / 2
	HH     = Height / 2
)

var running, drawHull bool
var points, hull convexhull.PointList
var px, py float64

func main() {
		points, hull = nil, nil
		points = make(convexhull.PointList, 0)
		hull = make(convexhull.PointList, 0)
		
		points = append(points, convexhull.MakePoint(100, 100))
		points = append(points, convexhull.MakePoint(150, 150))
		points = append(points, convexhull.MakePoint(150, 150))
		points = append(points, convexhull.MakePoint(175, 150))
		points = append(points, convexhull.MakePoint(100, 200))
		points = append(points, convexhull.MakePoint(200, 100))
		points = append(points, convexhull.MakePoint(200, 200))
		
		hull, _ = points.Compute()
		hull.PrintLines()
}


