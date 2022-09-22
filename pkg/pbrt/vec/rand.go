package vec

import "math/rand"

func uniform(min, max float64, rng *rand.Rand) float64 {
	return min + (max-min)*rng.Float64()
}

func RandomInUnitDisk(rng *rand.Rand) Vec {
again:
	p := Vec{
		uniform(-1, 1, rng),
		uniform(-1, 1, rng),
		0,
	}

	if p.LenSqr() >= 1 {
		goto again
	}

	return p
}
