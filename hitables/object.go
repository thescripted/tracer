package hitables

import "github.com/thescripted/sandbox-raytracing/geom"

type Objects []Sphere

func (o Objects) Hit(ray geom.Ray, tMin, tMax float64) (HitRecord, bool) {
	isHit := false
	globalRecord := HitRecord{}
	closest := tMax
	for _, sphere := range o {
		if record, ok := sphere.Hit(ray, tMin, closest); ok {
			isHit = true
			closest = record.t
			globalRecord = record
		}
	}
	return globalRecord, isHit
}
