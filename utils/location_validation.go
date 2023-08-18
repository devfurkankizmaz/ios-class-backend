package utils

func IsValidLatitude(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}

func IsValidLongitude(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}
