package main

// https://projecteuler.net/problem=816
// go fmt ./... && go vet ./... && go test && go run 816.go

//
// Precalcuate all of the s(n) values so we don't have to recalcuate them.
//
// The actual distance from one point to another is:
//   sqrt( (x1-x2)^2 + (y1-y2)^2 )
// This can be slow to calcuate. Since we need to compare *lots* of point
// distances, use something faster. This equation gives a value that is
// proportional to the actual distance so is useful for comparisons:
//   abs(x1-x2) + abs(y1-y2)
// This can also be written as:
//   max(a.x, b.x) - min(a.x, b.x) + max(a.y, b.y) - min(a.y, b.y)
// Once we know which two points are closest we can then do the more
// expensive calculation to find their actual distance.
//
// There are 2 million points to consider (k=2000000). If we do a naive
// comparison that will execute in O(n^2) time. That will be far too slow.
// That is also far more comparisons than are actually necessary. If the
// min distance found so far is 20 then we know that (1, ??) and (999, ??)
// cannot be closer no matter what the Y value. Sort the list of points by
// the X value and stop comparing once the two X values are farther apart
// than the current min discovered distance.
//

import (
	"fmt"
	"math"
	"sort"
)

var (
	sCache = map[int]int64{
		0: 290797,
	}
	cacheMaxN = 0
)

type point struct {
	x int64
	y int64
}

// sFunc returns s(n) = s(n-1)^2 % 50515093 where s(0) = 290797
func sFunc(n int) int64 {
	val := sCache[n-1]
	return (val * val) % 50515093
}

// loadCache loads the first maxN values into sCache
func loadCache(maxN int) {
	for i := cacheMaxN + 1; i <= maxN*2; i++ {
		sCache[i] = sFunc(i)
	}
	cacheMaxN = maxN
}

// s returns the cached value of s(n)
func s(n int) int64 {
	return sCache[n]
}

// P returns a point with coordinates P(n) = (s(2n), s(2n+1))
func P(n int) (int64, int64) {
	i := 2 * n
	return s(i), s(i + 1)
}

// distance returns the actual distance between two points
func distance(a, b point) float64 {
	deltaX := a.x - b.x
	deltaY := a.y - b.y

	sumSquares := deltaX*deltaX + deltaY*deltaY

	return math.Sqrt(float64(sumSquares))
}

// distanceFast returns an approximation of the distance between two points
func distanceFast(a, b point) int64 {
	return max(a.x, b.x) - min(a.x, b.x) + max(a.y, b.y) - min(a.y, b.y)
}

// minDistance returns the minimum distance among a constellation of points. It is assumed the points are sorted by X increasing.
func minDistance(points []point) float64 {
	minDist := distanceFast(points[0], points[1])
	minI := 0
	minJ := 1

	for i := 1; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if points[j].x-points[i].x >= minDist {
				break
			}
			dist := distanceFast(points[i], points[j])
			if dist < minDist {
				minDist = dist
				minI = i
				minJ = j
			}
		}
	}

	return distance(points[minI], points[minJ])
}

// d returns the shortest distance of any two (distinct) points among P(0)...P(k-1)
func d(k int) float64 {
	points := []point{}

	for n := 0; n <= k-1; n++ {
		p := point{}
		p.x, p.y = P(n)
		points = append(points, p)
	}

	// Sort points by the x value
	sort.SliceStable(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	return minDistance(points)
}

func main() {
	fmt.Printf("Welcome to 816\n\n")

	k := 2000000
	loadCache(k)

	shortest := d(k)
	fmt.Println("For k =", k, "shortest distance =", shortest)
}
