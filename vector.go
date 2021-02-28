package vector

import (
	"errors"
	"math"
)

var (
	// You can't add/subtract/dot product/cross product two vectors of different sizes.
	// https://www.ling.upenn.edu/courses/cogs501/LinearAlgebraReview.html
	ErrDiffSizes         = errors.New("vector: two vectors of different sizes")
	ErrAvailableOnly3Dim = errors.New("vector: available only for 3-dimensional vectors")
)

type Vector []float64

// Magnitude = sqrt(x1^2 + x2^2 + ... + xN^2)
func (vec Vector) Magnitude() float64 {
	sum := 0.0
	for _, val := range vec {
		sum += val * val
	}
	return math.Sqrt(sum)
}

// Add = v1 + v2 = [x1v1 + x1v2, v2v1 + x2v2, ... xNv1 + xNv2]
func Add(v1, v2 Vector) (Vector, error) {
	n := len(v1)
	if n != len(v2) {
		return nil, ErrDiffSizes
	}

	res := make(Vector, n)
	for i := 0; i < n; i++ {
		res[i] = v1[i] + v2[i]
	}
	return res, nil
}

// Subtract = v1 - v2 = [x1v1 - x1v2, v2v1 - x2v2, ... xNv1 - xNv2]
func Subtract(v1, v2 Vector) (Vector, error) {
	n := len(v1)
	if n != len(v2) {
		return nil, ErrDiffSizes
	}

	res := make(Vector, n)
	for i := 0; i < n; i++ {
		res[i] = v1[i] - v2[i]
	}
	return res, nil
}

// Multiply = v1 * v2 = [x1v1 * x1v2, v2v1 * x2v2, ... xNv1 * xNv2]
func Multiply(v1, v2 Vector) (Vector, error) {
	n := len(v1)
	if n != len(v2) {
		return nil, ErrDiffSizes
	}

	res := make(Vector, n)
	for i := 0; i < n; i++ {
		res[i] = v1[i] * v2[i]
	}
	return res, nil
}

// Divide = v1 / v2 = [x1v1 / x1v2, v2v1 / x2v2, ... xNv1 / xNv2]
func Divide(v1, v2 Vector) (Vector, error) {
	n := len(v1)
	if n != len(v2) {
		return nil, ErrDiffSizes
	}

	res := make(Vector, n)
	for i := 0; i < n; i++ {
		res[i] = v1[i] / v2[i]
	}
	return res, nil
}

func DotProduct(v1, v2 Vector) (float64, error) {
	n := len(v1)
	if n != len(v2) {
		return 0, ErrDiffSizes
	}

	var prod float64
	for i := 0; i < n; i++ {
		prod += v1[i] * v2[i]
	}
	return prod, nil
}

// CrossProduct
// Implemented only for 3 dimensional vector
func CrossProduct(v1, v2 Vector) (Vector, error) {
	if len(v1) != 3 || len(v2) != 3 {
		return nil, ErrAvailableOnly3Dim
	}

	return Vector{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}, nil
}
