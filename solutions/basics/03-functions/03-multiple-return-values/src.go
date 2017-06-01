package basics

func ComputePolygonCenter(coordinates [][2]float64) (x, y float64) {
	var nbPoints, sumX, sumY float64
	for _, point := range coordinates {
		nbPoints++
		sumX += point[0]
		sumY += point[1]
	}
	return sumX / nbPoints, sumY / nbPoints
}
