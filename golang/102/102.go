package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	x int
	y int
}

func loadTriangles() [][]vertex {
	var triangles [][]vertex

	raw, _ := os.ReadFile("102_triangles.txt")
	lines := strings.Split(string(raw), "\n")
	for _, line := range lines {
		vertices := strings.Split(line, ",")
		if len(vertices) <= 1 {
			continue
		}
		var v []vertex
		x, _ := strconv.Atoi(vertices[0])
		y, _ := strconv.Atoi(vertices[1])
		v = append(v, vertex{x, y})
		x, _ = strconv.Atoi(vertices[2])
		y, _ = strconv.Atoi(vertices[3])
		v = append(v, vertex{x, y})
		x, _ = strconv.Atoi(vertices[4])
		y, _ = strconv.Atoi(vertices[5])
		v = append(v, vertex{x, y})
		triangles = append(triangles, v)
	}

	return triangles
}

func intersectX(A, B vertex) (float64, bool) {
	if (A.y >= 0 && B.y <= 0) || (B.y >= 0 && A.y <= 0) {
		if A.x == B.x {
			return float64(A.x), true
		}
		rise := A.y - B.y
		run := A.x - B.x
		m := float64(rise) / float64(run)
		b := float64(A.y) - m*float64(A.x)
		return -b / m, true
	}

	return 0, false
}

func intersectY(A, B vertex) (float64, bool) {
	if (A.x >= 0 && B.x <= 0) || (B.x >= 0 && A.x <= 0) {
		if A.y == B.y {
			return float64(A.y), true
		}
		rise := A.y - B.y
		run := A.x - B.x
		m := float64(rise) / float64(run)
		b := float64(A.y) - m*float64(A.x)
		return b, true
	}

	return 0, false
}

// Is the origin *interior* to the triangle?
func containsOrigin(triangle []vertex) bool {
	A := triangle[0]
	B := triangle[1]
	C := triangle[2]
	Origin := vertex{0, 0}

	// If a vertex is on the origin then the
	// triangle does not have the origin on the
	// *interior*.
	if A == Origin || B == Origin || C == Origin {
		return false
	}

	// If one of the edges is on an axis then the
	// triangle does not have the origin on the
	// *interior*.
	if (A.x == 0 && B.x == 0) || (A.x == 0 && C.x == 0) {
		return false
	}
	if (A.y == 0 && B.y == 0) || (A.y == 0 && C.y == 0) {
		return false
	}

	// If a point is inside a closed shape then a line drawn
	// from that point to a point known to be outside of the
	// shape will cross an odd number of lines. Since this
	// triangle is a convex hull, that odd number will be
	// exactly 1.
	// With the notable exception that if the line drawn
	// from the interior point intersects a vertex then
	// it will show as having crossed 2 lines.
	intersectCount := 0
	intersect := false
	x := 0.0
	y := 0.0

	// From origin out along the Positive X-axis
	intersectCount = 0
	x, intersect = intersectX(A, B)
	if intersect && x >= 0 {
		intersectCount++
	}
	x, intersect = intersectX(B, C)
	if intersect && x >= 0 {
		intersectCount++
	}
	x, intersect = intersectX(A, C)
	if intersect && x >= 0 {
		intersectCount++
	}
	if intersectCount == 2 {
		if A.y == 0 || B.y == 0 || C.y == 0 {
			intersectCount--
		}
	}
	if intersectCount != 1 {
		return false
	}

	// From origin out along the Positive Y-axis
	intersectCount = 0
	y, intersect = intersectY(A, B)
	if intersect && y >= 0 {
		intersectCount++
	}
	y, intersect = intersectY(B, C)
	if intersect && y >= 0 {
		intersectCount++
	}
	y, intersect = intersectY(A, C)
	if intersect && y >= 0 {
		intersectCount++
	}
	if intersectCount == 2 {
		if A.x == 0 || B.x == 0 || C.x == 0 {
			intersectCount--
		}
	}
	if intersectCount != 1 {
		return false
	}

	// From origin out along the Negative X-axis
	intersectCount = 0
	x, intersect = intersectX(A, B)
	if intersect && x <= 0 {
		intersectCount++
	}
	x, intersect = intersectX(B, C)
	if intersect && x <= 0 {
		intersectCount++
	}
	x, intersect = intersectX(A, C)
	if intersect && x <= 0 {
		intersectCount++
	}
	if intersectCount == 2 {
		if A.y == 0 || B.y == 0 || C.y == 0 {
			intersectCount--
		}
	}
	if intersectCount != 1 {
		return false
	}

	// From origin out along the Negative Y-axis
	intersectCount = 0
	y, intersect = intersectY(A, B)
	if intersect && y <= 0 {
		intersectCount++
	}
	y, intersect = intersectY(B, C)
	if intersect && y <= 0 {
		intersectCount++
	}
	y, intersect = intersectY(A, C)
	if intersect && y <= 0 {
		intersectCount++
	}
	if intersectCount == 2 {
		if A.x == 0 || B.x == 0 || C.x == 0 {
			intersectCount--
		}
	}
	if intersectCount != 1 {
		return false
	}

	return true
}

func main() {
	count := 0
	for _, triangle := range loadTriangles() {
		if containsOrigin(triangle) {
			count++
		}
	}
	fmt.Println("Triangles that contain the origin:", count)
}
