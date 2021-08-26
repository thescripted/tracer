package camera

import (
	"github.com/thescripted/sandbox-raytracing/geom"
	"math"
	"math/rand"
)

type Camera struct {
	origin          geom.Vec3
	lowerLeftCorner geom.Vec3
	horizontal      geom.Vec3
	vertical        geom.Vec3
	aperture        float64
	u               geom.Vec3
	v               geom.Vec3
	w               geom.Vec3
}

// TODO: Create more sensible defaults. Users shouldn't need to worry about setting focusDist/aperture/etc.
func New(LookFrom, lookAt, vup geom.Vec3, vfov, aspect, aperture, focusDist float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := LookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)

	return Camera{
		origin:          LookFrom,
		lowerLeftCorner: LookFrom.Sub(u.Scale(halfWidth * focusDist)).Sub(v.Scale(halfHeight * focusDist)).Sub(w.Scale(focusDist)),
		horizontal:      u.Scale(2 * halfWidth * focusDist),
		vertical:        v.Scale(2 * halfHeight * focusDist),
		aperture:        aperture,
		u:               u,
		v:               v,
		w:               w,
	}
}

func (c Camera) GetRay(u, v float64) geom.Ray {
	rd := randomInUnitDisk().Scale(c.aperture / 2)
	offset := c.u.Scale(rd.X()).Add(c.v.Scale(rd.Y()))
	return geom.Ray{
		A: c.origin.Add(offset),
		B: c.lowerLeftCorner.Add(c.horizontal.Scale(u)).Add(c.vertical.Scale(v)).Sub(c.origin).Sub(offset),
	}
}

func randomInUnitDisk() geom.Vec3 {
	var p geom.Vec3
	for {
		p = geom.Vec3{rand.Float64(), rand.Float64(), 0}.Scale(2).Sub(geom.Vec3{1, 1, 0})
		if p.LenSq() < 1 {
			break
		}
	}
	return p
}
