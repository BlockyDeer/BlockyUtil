package blockyutil

// ensure that the expression will not less than 0
func ZeroFloor(value float64) float64 {
	if value > 0 {
		return value
	} else {
		return 0
	}
}

// ensure that the expression will not less than a specific value
func Rooftop(value float64, top float64) float64 {
	if value < top {
		return top
	} else {
		return value
	}
}

// check if the value is between the minimum and the maximum or it really is the minimum or the maximum
func Between(value float64, minimum float64, maximum float64) bool {
	return value <= maximum && value >= minimum
}
