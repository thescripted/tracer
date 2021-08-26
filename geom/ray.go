package geom

type Ray struct {
	A Vec3
	B Vec3
}

func (r Ray) Origin() Vec3 {
	return r.A
}

func (r Ray) Direction() Vec3 {
	return r.B
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.A.Add(r.B.Scale(t))
}
