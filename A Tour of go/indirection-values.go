package main
import "fmt"
import "math"

type Vertex struct {
	X , Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFun(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFun(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFun(*p))
}