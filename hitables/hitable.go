package hitables

import (
	"github.com/thescripted/sandbox-raytracing/geom"
	"github.com/thescripted/sandbox-raytracing/materials"
)

// Hitable is an interface for all objects that must react to rays.
type Hitable interface {
	Hit(ray geom.Ray, tMin, tMax float64) (HitRecord, bool)
}

// HitRecord contains recorded information about where an object was hit.
type HitRecord struct {
	t        float64
	Point    geom.Vec3
	Normal   geom.Vec3
	Material materials.Material
}
