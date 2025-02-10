package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {

	//Arange
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	//Act
	got, err := geo.GetMyLocation(city)

	//Assert
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}
