package main

import (
	"flag"
	"fmt"
	"math"
)

const r = 6371

func main() {
	lat := flag.Float64("lat", 59.9199375, "Lantitude value")
	lon := flag.Float64("lon", 30.3410583, "Lontitude value")
	b := flag.Float64("bear", 134, "Bearing to the point (degree)")
	d := flag.Float64("dist", 632, "Distance to the point (km)")
	flag.Parse()
	fmt.Println(LatLonCalc(*lat, *lon, *b, *d))
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
