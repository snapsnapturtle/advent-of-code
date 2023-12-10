package util

// IsPointWithinPolyline taken from https://wrfranklin.org/Research/Short_Notes/pnpoly.html#Explanation
func IsPointWithinPolyline(polyline [][2]int, point [2]int) bool {
	vertices := len(polyline)
	c := false

	for i, j := 0, vertices-1; i < vertices; j, i = i, i+1 {
		if ((polyline[i][1] > point[1]) != (polyline[j][1] > point[1])) &&
			(point[0] < (polyline[j][0]-polyline[i][0])*(point[1]-polyline[i][1])/(polyline[j][1]-polyline[i][1])+polyline[i][0]) {
			c = !c
		}
	}

	return c
}
