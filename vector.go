/*
This file is part of BlockyUtil project. It is licensed under The MIT License.

The MIT License (MIT)
Copyright (c) 2024 BlockyDeer <blockydeer@outlook.com>
Permission is hereby granted, free of charge, to any person obtaining a copy of this
software and associated documentation files (the “Software”), to deal in the Software
without restriction, including without limitation the rights to use, copy, modify,
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to the following
conditions:
The above copyright notice and this permission notice shall be included in all copies
or substantial portions of the Software.
The Software is provided “as is”, without warranty of any kind, express or implied,
including but not limited to the warranties of merchantability, fitness for a
particular purpose and noninfringement. In no event shall the authors or copyright
holders be liable for any claim, damages or other liability, whether in an action of
contract, tort or otherwise, arising from, out of or in connection with the software
or the use or other dealings in the Software.
*/
package blockyutil

import "math"

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Vec2d = Vec2[float64]
type Vec2f = Vec2[float32]
type Vec2i = Vec2[int32]
type Vec2u = Vec2[uint32]

type Vec2[T Numeric] struct {
	X T
	Y T
}

func CreateVec[T Numeric](x T, y T) Vec2[T] {
	return Vec2[T]{
		X: x,
		Y: y,
	}
}

func ZeroVec[T Numeric]() Vec2[T] {
	return CreateVec(T(0), T(0))
}

func (vec *Vec2[T]) EndPoint(startPoint Vec2[T]) Vec2[T] {
	return vec.Add(startPoint)
}
func (vec *Vec2[T]) Add(other Vec2[T]) Vec2[T] {
	return CreateVec(vec.X+other.X, vec.Y+other.Y)
}
func (vec *Vec2[T]) Multiply(rate float64) Vec2[T] {
	return CreateVec(T(float64(vec.X)*rate), T(float64(vec.Y)*rate))
}
func (vec *Vec2[T]) Sub(other Vec2[T]) Vec2[T] {
	return CreateVec(vec.X-other.X, vec.Y-vec.Y)
}
func (vec *Vec2[T]) Dot(other Vec2[T]) float64 {
	return float64(vec.X*other.X + vec.Y*other.Y)
}
func (vec *Vec2[T]) Cross(other Vec2[T]) Vec2[T] {
	return CreateVec(vec.X*other.Y, vec.Y*other.X)
}
func (vec *Vec2[T]) Mod() float64 {
	return math.Sqrt(float64(vec.X*vec.X + vec.Y * vec.Y))
}
func (vec *Vec2[T]) Normalize() Vec2[T] {
	mod := vec.Mod()
	if mod != 0 {
		x := vec.X / T(mod)
		y := vec.Y / T(mod)

		return CreateVec(x, y)
	} else {
		return ZeroVec[T]()
	}
}
