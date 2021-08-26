package main

import (
	"fmt"
	"github.com/thescripted/sandbox-raytracing/camera"
	"github.com/thescripted/sandbox-raytracing/geom"
	"github.com/thescripted/sandbox-raytracing/hitables"
	"github.com/thescripted/sandbox-raytracing/materials"

	"math"
	"math/rand"
)

func main() {
	const (
		width    = 1200
		height   = 800
		aliasing = 15
	)

	// Camera Init
	var (
		lookFrom    = geom.Vec3{13, 2, 3}
		lookAt      = geom.Vec3{0, 0, 0}
		verticalUp  = geom.Vec3{0, 1, 0}
		distToFocus = 10.0
		aperture    = 0.1
		fieldOfView = 20.0
		aspect      = float64(width) / float64(height)
	)

	cam := camera.New(
		lookFrom,
		lookAt,
		verticalUp,
		fieldOfView,
		aspect,
		aperture,
		distToFocus,
	)

	world := generate()

	// Header for PPM file
	fmt.Printf(
		"P3\n"+
			"%d %d\n"+
			"255\n", width, height,
	)

	// Ray Casting
	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			col := geom.Vec3{0, 0, 0}
			for s := 0; s < aliasing; s++ {
				u := (float64(i) + rand.Float64()) / float64(width)
				v := (float64(j) + rand.Float64()) / float64(height)
				getRay := cam.GetRay(u, v)
				col = col.Add(color(getRay, world, 0))
			}
			col = col.Scale(1.0 / float64(aliasing))
			col = geom.Vec3{math.Sqrt(col.X()), math.Sqrt(col.Y()), math.Sqrt(col.Z())}
			ir := int(255.99 * col.X())
			ig := int(255.99 * col.Y())
			ib := int(255.99 * col.Z())
			fmt.Printf("%d %d %d\n", ir, ig, ib)

		}
	}
}

func color(r geom.Ray, world hitables.Hitable, depth int) geom.Vec3 {
	if record, ok := world.Hit(r, 0.001, 1000); ok {
		scattered, attenuation, ok := record.Material.Scatter(r, record)
		if depth < 50 && ok {
			return attenuation.Times(color(scattered, world, depth+1))
		} else {
			return geom.Vec3{0, 0, 0}
		}
	}
	unitDir := r.Direction().Unit()
	t := 0.5 * (unitDir.Y() + 1)
	return geom.Vec3{1, 1, 1}.Scale(1 - t).Add(geom.Vec3{0.5, 0.7, 1}.Scale(t))

}

func generate() hitables.Objects {
	world := hitables.Objects{
		hitables.Sphere{
			Center:   geom.Vec3{0, -1000, 0},
			Radius:   1000,
			Material: materials.Lambertian{Albedo: geom.Vec3{0.5, 0.5, 0.5}},
		},
		hitables.Sphere{
			Center:   geom.Vec3{0, 1, 0},
			Radius:   1.0,
			Material: materials.Dielectric{RefIdx: 1.5},
		},
		hitables.Sphere{
			Center:   geom.Vec3{-4, 1, 0},
			Radius:   1.0,
			Material: materials.Lambertian{Albedo: geom.Vec3{0.4, 0.2, 0.1}},
		},
		hitables.Sphere{
			Center:   geom.Vec3{4, 1, 0},
			Radius:   1.0,
			Material: materials.Metal{Albedo: geom.Vec3{0.7, 0.65, 0.5}},
		},
	}

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := geom.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(geom.Vec3{4, 0.2, 0}).Len() > 0.9 {
				if chooseMat < 0.8 { // Matte
					world = append(world, hitables.Sphere{
						Center: center,
						Radius: 0.2,
						Material: materials.Lambertian{Albedo: geom.Vec3{
							rand.Float64() * rand.Float64(),
							rand.Float64() * rand.Float64(),
							rand.Float64() * rand.Float64(),
						}},
					})
				} else if chooseMat < 0.95 { // Metal
					world = append(world, hitables.Sphere{
						Center: center,
						Radius: 0.2,
						Material: materials.Metal{
							Albedo: geom.Vec3{
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
							},
							Fuzz: 0.5 * rand.Float64(),
						},
					})
				} else { // Glass
					world = append(world, hitables.Sphere{
						Center:   center,
						Radius:   0.2,
						Material: materials.Dielectric{RefIdx: 1.5},
					})
				}
			}
		}
	}
	return world
}
