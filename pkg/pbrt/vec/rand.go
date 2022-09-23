package vec

import "math/rand"

func uniform(min, max float64, rng *rand.Rand) float64 {
	return min + (max-min)*rng.Float64()
}

func RandomUniform(min, max float64, rng *rand.Rand) Vec {
	return Vec{
		uniform(min, max, rng),
		uniform(min, max, rng),
		uniform(min, max, rng),
	}
}

func RandomInUnitDisk(rng *rand.Rand) Vec {
	for {
		p := Vec{
			uniform(-1, 1, rng),
			uniform(-1, 1, rng),
			0,
		}

		if p.LenSqr() >= 1 {
			continue
		}

		return p
	}
}

func RandomInUnitSphere(rng *rand.Rand) Vec {
	for {
		p := RandomUniform(-1, 1, rng)
		if p.LenSqr() >= 1 {
			continue
		}
		return p
	}
}
