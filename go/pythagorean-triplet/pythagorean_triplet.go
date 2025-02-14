package pythagorean

import "slices"

type Triplet [3]int

type Trip struct {
	a, b, c int
}

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (trips []Triplet) {
	return findTrips(min, max, 0)
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (trips []Triplet) {
	return findTrips(1, p, p)
}

func findTrips(min, max, p int) (trips []Triplet) {
	found := map[Trip]bool{}
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if p == 0 || a+b+c == p {
					if a*a+b*b == c*c {
						trip := []int{a, b, c}
						slices.Sort(trip)
						tripStr := Trip{trip[0], trip[1], trip[2]}
						if found[tripStr] {
							continue
						}
						found[tripStr] = true
						var tripArr [3]int
						copy(tripArr[:], trip)
						trips = append(trips, tripArr)
					}
				}
			}
		}
	}
	return trips
}
