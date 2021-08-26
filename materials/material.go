package materials

import (
	"github.com/thescripted/sandbox-raytracing/geom"
	"github.com/thescripted/sandbox-raytracing/hitables"
	"math"
	"math/rand"
)

type Material interface {
	Scatter(ray geom.Ray, record hitables.HitRecord) (scattered geom.Ray, attenuation geom.Vec3, ok bool)
}
type Metal struct {
	Albedo geom.Vec3
	Fuzz   float64
}

type Lambertian struct {
	Albedo geom.Vec3
}

type Dielectric struct {
	RefIdx float64
}

func (m Metal) Scatter(rayIn geom.Ray, record hitables.HitRecord) (scattered geom.Ray, attenuation geom.Vec3, ok bool) {
	reflected := Reflect(rayIn.Direction().Unit(), record.Normal)
	scattered = geom.Ray{
		A: record.Point,
		B: reflected.Add(RandInUnitSphere().Scale(m.Fuzz)),
	}
	attenuation = m.Albedo
	if scattered.Direction().Dot(record.Normal) > 0 {
		return scattered, attenuation, true
	}
	return geom.Ray{}, geom.Vec3{}, false
}

func (l Lambertian) Scatter(rayIn geom.Ray, record hitables.HitRecord) (scattered geom.Ray, attenuation geom.Vec3, ok bool) {
	target := record.Point.Add(record.Normal).Add(RandInUnitSphere())
	scattered = geom.Ray{
		A: record.Point,
		B: target.Sub(record.Point),
	}
	attenuation = l.Albedo
	return scattered, attenuation, true
}

func (d Dielectric) Scatter(rayIn geom.Ray, record hitables.HitRecord) (scattered geom.Ray, attenuation geom.Vec3, ok bool) {
	var outwardNormal geom.Vec3
	var ni_over_nt, cosine, reflectProb float64
	reflected := Reflect(rayIn.Direction(), record.Normal)
	attenuation = geom.Vec3{1, 1, 1}

	if rayIn.Direction().Dot(record.Normal) > 0 {
		outwardNormal = record.Normal.Scale(-1)
		ni_over_nt = d.RefIdx
		cosine = (rayIn.Direction().Dot(record.Normal) * d.RefIdx) / (rayIn.Direction().Len())
	} else {
		outwardNormal = record.Normal
		ni_over_nt = 1.0 / d.RefIdx
		cosine = -rayIn.Direction().Dot(record.Normal) / (rayIn.Direction().Len())
	}

	refracted, ok := Refract(rayIn.Direction(), outwardNormal, ni_over_nt)

	if ok {
		reflectProb = schlick(cosine, d.RefIdx)
	} else {
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		scattered = geom.Ray{
			A: record.Point,
			B: reflected,
		}
	} else {
		scattered = geom.Ray{
			A: record.Point,
			B: refracted,
		}
	}
	return scattered, attenuation, true
}

func Reflect(v, n geom.Vec3) geom.Vec3 {
	return v.Sub(n.Scale(2 * n.Dot(v)))
}

func Refract(v, n geom.Vec3, ni_over_nt float64) (geom.Vec3, bool) {
	uv := v.Unit()
	dt := uv.Dot(n)
	discriminant := 1.0 - ni_over_nt*ni_over_nt*(1-dt*dt)
	if discriminant > 0 {
		firstpart := uv.Sub(n.Scale(dt))
		firstpart = firstpart.Scale(ni_over_nt)
		secondpart := n.Scale(math.Sqrt(discriminant))
		refracted := firstpart.Sub(secondpart)
		return refracted, true
	}
	return geom.Vec3{}, false
}

func RandInUnitSphere() geom.Vec3 {
	var p geom.Vec3
	for {
		p = geom.Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.Scale(2).Sub(geom.Vec3{1, 1, 1})
		if p.LenSq() < 1 {
			break
		}
	}
	return p
}

func schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
