package main

import (
	"fmt"
	"strings"
)

const (
	base32 = "0123456789bcdefghjkmnpqrstuvwxyz"
)

func encode(latitude float64, longitude float64, precision int) string {
	minLat, maxLat := -90.0, 90.0
	minLon, maxLon := -180.0, 180.0
	geohash := make([]byte, 0, precision)
	even := true

	for len(geohash) < precision {
		if even {
			mid := (minLon + maxLon) / 2
			if longitude > mid {
				geohash = append(geohash, '1')
				minLon = mid
			} else {
				geohash = append(geohash, '0')
				maxLon = mid
			}
		} else {
			mid := (minLat + maxLat) / 2
			if latitude > mid {
				geohash = append(geohash, '1')
				minLat = mid
			} else {
				geohash = append(geohash, '0')
				maxLat = mid
			}
		}
		even = !even
	}

	return string(geohash)
}

func base32Encode(geohash string) string {
	var result strings.Builder
	for i := 0; i < len(geohash); i += 5 {
		end := i + 5
		if end > len(geohash) {
			end = len(geohash)
		}
		chunk := geohash[i:end]
		index := binaryToDecimal(chunk)
		result.WriteByte(base32[index])
	}
	return result.String()
}

func binaryToDecimal(binary string) int {
	result := 0
	for _, bit := range binary {
		result = result*2 + int(bit-'0')
	}
	return result
}

func main() {
	latitude := 40.741895
	longitude := -73.989308
	precision := 6

	geohashBinary := encode(latitude, longitude, precision*5)
	geohash := base32Encode(geohashBinary)

	fmt.Println(geohash)
}
