package triangles

const (
	MAX_TRIANGLE = 100000
)

var (
	Triangles [MAX_TRIANGLE + 1]bool
)

func Triangle(number int) bool {
	return Triangles[number]
}

func triangles() {
	for i := 0; i < len(Triangles); i++ {
		Triangles[i] = false
	}
	i := 1
	for {
		t := (i * (i + 1)) >> 1
		if t >= len(Triangles) {
			break
		}
		Triangles[t] = true
		i++
	}
}

func Init() {
	triangles()
}
