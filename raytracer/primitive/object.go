package primitive

type Object interface {
	Material() Material
	Intersect(ray *Ray) []Hit
}
