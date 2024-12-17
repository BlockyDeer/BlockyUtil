package blockyutil

import "math"

type Vec3d = Vec3[float64]
type Vec3f = Vec3[float32]
type Vec3i = Vec3[int32]
type Vec3u = Vec3[uint32]

type Vec3[T Numeric] struct {
	X T
	Y T
	Z T
}

func CreateVec3[T Numeric](x T, y T, z T) Vec3[T] {
	return Vec3[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

func ZeroVec3[T Numeric]() Vec3[T] {
	return CreateVec3(T(0), T(0), T(0))
}

func (vec *Vec3[T]) EndPoint(startPoint Vec3[T]) Vec3[T] {
	return vec.Add(startPoint)
}

func (vec *Vec3[T]) Add(other Vec3[T]) Vec3[T] {
	return CreateVec3(vec.X+other.X, vec.Y+other.Y, vec.Z+other.Z)
}

func (vec *Vec3[T]) Multiply(rate float64) Vec3[T] {
	return CreateVec3(T(float64(vec.X)*rate), T(float64(vec.Y)*rate), T(float64(vec.Z)*rate))
}

func (vec *Vec3[T]) Sub(other Vec3[T]) Vec3[T] {
	return CreateVec3(vec.X-other.X, vec.Y-other.Y, vec.Z-other.Z)
}

func (vec *Vec3[T]) Dot(other Vec3[T]) float64 {
	return float64(vec.X*other.X + vec.Y*other.Y + vec.Z*other.Z)
}

func (vec *Vec3[T]) Mod() float64 {
	return math.Sqrt(float64(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z))
}

func (vec *Vec3[T]) Normalize() Vec3[T] {
	mod := vec.Mod()
	if mod != 0 {
		x := vec.X / T(mod)
		y := vec.Y / T(mod)
		z := vec.Z / T(mod)

		return CreateVec3(x, y, z)
	} else {
		return ZeroVec3[T]()
	}
}
