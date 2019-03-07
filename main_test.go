package main

import (
	"fmt"
	"math"
	"testing"
)

/* func TestOk(t *testing.T){
err:=
if err!=nil{
t.Errorf("Test for OK fail")
}
} */

func TestRad2deg(t *testing.T) {
	got := fmt.Sprint(Rad2deg(math.Pi))
	if got != "180" {
		t.Errorf("rad2deg(pi) = %s; want 180", got)
	}
}
func TestDeg2rad(t *testing.T) {
	got := Deg2rad(180)
	if got != math.Pi {
		t.Errorf("rad2deg(pi) = %f; want 3.14", got)
	}
}
func TestLatLonCalc(t *testing.T) {
	lat := 59.9199375
	lon := 30.3410583
	b := 134.0
	d := 632.0
	got := fmt.Sprint(LatLonCalc(lat, lon, b, d))
	if got != "55.75003721967312 37.61320899547458" {
		t.Errorf("rad2deg(pi) = %s; want 55.75003721967312 37.61320899547458", got)
	}
}
