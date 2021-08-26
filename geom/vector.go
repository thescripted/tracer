package geom

import "math"

type Vec3 [3]float64

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v[0] + v2[0], v[1] + v2[1], v[2] + v2[2]}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v[0] - v2[0], v[1] - v2[1], v[2] - v2[2]}
}

func (v Vec3) Times(v2 Vec3) Vec3 {
	return Vec3{v[0] * v2[0], v[1] * v2[1], v[2] * v2[2]}
}

func (v Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{v[0] / v2[0], v[1] / v2[1], v[2] / v2[2]}
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v[1]*v2[2] - v[2]*v2[1],
		-(v[0]*v2[2] - v[2]*v2[0]),
		v[0]*v2[1] - v[1]*v2[0],
	}
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec3) LenSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec3) Unit() Vec3 {
	k := 1.0 / v.Len()
	return Vec3{v[0] * k, v[1] * k, v[2] * k}
}
