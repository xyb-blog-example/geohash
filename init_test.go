package geohash

import (
    "testing"
)

func TestCreateGeoMap(t *testing.T) {
    CreateGeoMap(180,90, -180, -90, 15, map[string]Node{"1":{X:116.599831, Y:39.925746}})
}
