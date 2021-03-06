package geo

// LinearRing type
type LinearRing struct {
	LineString
}

/* CONSTRUCTOR */

// NewLinearRing LinearRing constructor
func NewLinearRing(points ...Point) LinearRing {
	var lr LinearRing

	firstPoint := points[0]
	lastPoint := points[len(points)-1]

	// check if ring is closed, else append first point to it
	if !firstPoint.Equals(lastPoint) {
		points = append(points, firstPoint)
	}

	lr.LineString = NewLineString(points...)

	return lr
}

/* GEOMETRY */

// GeometryType returns GeometryType as a string
func (l LinearRing) GeometryType() string {
	return "LinearRing"
}

/* INTERFACE GUARD */
var (
	_ Geometry = (*LinearRing)(nil)
	_ Curve    = (*LinearRing)(nil)
)
