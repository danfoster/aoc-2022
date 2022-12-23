package geom

type Point3D struct {
	X int
	Y int
	Z int
}

func (a Point3D) Add(b Point3D) Point3D {
	return Point3D{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}
