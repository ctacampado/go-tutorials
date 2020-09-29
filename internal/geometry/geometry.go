package geometry

import "math"

// Point represents a 2D point
type Point struct {
	id   string
	X, Y float64
}

// Distance is a traditional function
func Distance(p, q Point) float64 {
	// d=√((x2-x1)²+(y2-y1)²)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance is the same thing as the above function
// but as a method to struct Point
func (p Point) Distance(q Point) float64 {
	// d=√((x2-x1)²+(y2-y1)²)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// UpdateX tries to update X but fails
func (p Point) UpdateX(x float64) {
	p.X = x
}

// Add two Points
func (p Point) Add(q Point) Point { return Point{X: p.X + q.X, Y: p.Y + q.Y} }

// Sub two Points
func (p Point) Sub(q Point) Point { return Point{X: p.X - q.X, Y: p.Y - q.Y} }

// ScalyBy scale up with factor
func (p *Point) ScalyBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// SetID sets the point id
func (p *Point) SetID(id string) { p.id = id }

// GetID sets the point id
func (p *Point) GetID() string { return p.id }

// Path represents a sequence of line segments
type Path []Point

// PathDistance is a traditional function
func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Distance retruns the distance traveled along the path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// TranslateBy translates points in a path given an offset and boolean
// 		false - subtract
//    true  - add
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}
