package vec

import "math"

type Vec struct {
	X, Y, Z float64
}

func New(x, y, z float64) Vec {
	return Vec{x, y, z}
}

func (v Vec) Add(w Vec) Vec {
	return Vec{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vec) Sub(w Vec) Vec {
	return Vec{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vec) Scaled(f float64) Vec {
	return Vec{v.X * f, v.Y * f, v.Z * f}
}

func (v Vec) LenSqr() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec) Len() float64 {
	return math.Sqrt(v.LenSqr())
}

func (v Vec) Normalized() Vec {
	return v.Scaled(1 / v.Len())
}

func Dot(u, v Vec) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func Cross(u, v Vec) Vec {
	return Vec{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

func Zero() Vec {
	return Vec{0, 0, 0}
}

func One() Vec {
	return Vec{1, 1, 1}
}
