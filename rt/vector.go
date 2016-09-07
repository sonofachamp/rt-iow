package rt

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func V(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vector) Negate() Vector {
	return Vector{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
	}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vector) Multiply(b Vector) Vector {
	return Vector{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func (a Vector) Divide(b Vector) Vector {
	return Vector{
		X: a.X / b.X,
		Y: a.Y / b.Y,
		Z: a.Z / b.Z,
	}
}

func (a Vector) MultiplyByScalar(b float64) Vector {
	return Vector{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func (a Vector) DivideByScalar(b float64) Vector {
	return Vector{
		X: a.X / b,
		Y: a.Y / b,
		Z: a.Z / b,
	}
}

func UnitVector(a Vector) Vector {
	return a.DivideByScalar(a.Length())
}
