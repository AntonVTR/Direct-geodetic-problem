package main

import (
	"fmt"
	"math"
)

const r = 6371

func main() {
	fmt.Println("Start")
	lat := 59.9199375
	lon := 30.3410583
	var b, d float64
	b = 134
	d = 632
	fmt.Println(LatLonCalc(lat, lon, b, d))
}

func LatLonCalc(lat float64, lon float64, bearing float64, dist float64) (float64, float64) {
	s := float64(dist) / float64(r)
	latD := Deg2rad(lat)
	bearD := Deg2rad(bearing)
	lat1 := Rad2deg(math.Asin(math.Cos(s)*math.Sin(latD) + math.Sin(s)*math.Cos(latD)*math.Cos(bearD)))
	lon1 := lon + Rad2deg(math.Atan(math.Sin(s)*math.Sin(bearD)/(math.Cos(s)*math.Cos(latD)-math.Sin(s)*math.Sin(latD)*math.Cos(bearD))))
	return lat1, lon1
}

func Rad2deg(rad float64) float64 {
	return rad * 180 / math.Pi

}
func Deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}
