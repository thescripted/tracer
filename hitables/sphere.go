package hitables

import (
	"github.com/thescripted/sandbox-raytracing/geom"
	"github.com/thescripted/sandbox-raytracing/materials"
	"math"
)

// A Sphere is a spherical Hitable.
type Sphere struct {
	Center   geom.Vec3
	Radius   float64
	Material materials.Material
}

func (s Sphere) Hit(ray geom.Ray, tMin, tMax float64) (HitRecord, bool) {
	oc := ray.Origin().Sub(s.Center)
	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * oc.Dot(ray.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	record := HitRecord{}
	if discriminant > 0 {
		root := (-b - math.Sqrt(discriminant)) / (2 * a)
		if root < tMax && root > tMin {
			record.t = root
			record.Point = ray.PointAtParameter(record.t)
			record.Normal = record.Point.Sub(s.Center).Scale(1 / s.Radius)
			record.Material = s.Material
			return record, true
		}

		root = (-b + math.Sqrt(discriminant)) / (2 * a)
		if root < tMax && root > tMin {
			record.t = root
			record.Point = ray.PointAtParameter(record.t)
			record.Normal = record.Point.Sub(s.Center).Scale(1 / s.Radius)
			record.Material = s.Material
			return record, true
		}
	}
	return record, false
}
