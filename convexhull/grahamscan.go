package convexhull

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}

type PointList []Point

func MakePoint(x float64, y float64) Point {
	return Point{X: x, Y: y}
}

func PrintStack(s *Stack) {
	v := s.top
	fmt.Printf("Stack: ")
	for v != nil {
		fmt.Printf("%v ", v.value)
		v = v.next
	}
	fmt.Println("")
}

//Implement sort interface
func (p PointList) Len() int {
	return len(p)
}

func (p PointList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PointList) Less(i, j int) bool {
	area := Area2(p[0], p[i], p[j])

	if area == 0 {
		x := math.Abs(p[i].X-p[0].X) - math.Abs(p[j].X-p[0].X)
		y := math.Abs(p[i].Y-p[0].Y) - math.Abs(p[j].Y-p[0].Y)

		if x < 0 || y < 0 {
			return true
		} else if x > 0 || y > 0 {
			return false
		} else {
			return false
		}
	}

	return area > 0
}

func (p PointList) FindLowestPoint() {
	m := 0
	for i := 1; i < len(p); i++ {
		//If lowest points are on the same line, take the rightmost point
		if (p[i].Y < p[m].Y) || ((p[i].Y == p[m].Y) && p[i].X > p[m].X) {
			m = i
		}
	}
	p[0], p[m] = p[m], p[0]
}

func (points PointList) Compute() (PointList, bool) {
	if len(points) < 3 {
		return nil, false
	}

	stack := new(Stack)
	points.FindLowestPoint()
	sort.Sort(&points)

	stack.Push(points[0])
	stack.Push(points[1])

	fmt.Println("-START-COMPUTE--------------------------------------")
	fmt.Printf("Sorted Points: %v\n", points)

	i := 2
	for i < len(points) {
		pi := points[i]

		PrintStack(stack)

		p1 := stack.top.next.value.(Point)
		p2 := stack.top.value.(Point)

		if isLeft(p1, p2, pi) {
			stack.Push(pi)
			i++
		} else {
			stack.Pop()
		}
	}

	//Copy the hull
	ret := make(PointList, stack.Len())
	top := stack.top
	count := 0
	for top != nil {
		ret[count] = top.value.(Point)
		top = top.next
		count++
	}

	fmt.Println("-END-COMPUTE----------------------------------------------")
	return ret, true
}

func isLeft(p0, p1, p2 Point) bool {
	return Area2(p0, p1, p2) > 0
}

func Area2(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

func (points PointList) PrintLines() {

	fmt.Println("-START-HULL--------------------------------------")
	for _, p := range points {
		fmt.Printf("x %f y %f\n", p.X, p.Y)
	}
	fmt.Println("-START-HULL--------------------------------------")
}