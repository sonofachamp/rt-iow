package rt

type Ray struct {
	A, B Vector
}

func NewRay(a, b Vector) Ray {
	return Ray{
		A: a,
		B: b,
	}
}

func (r Ray) Origin() Vector {
	return r.A
}

func (r Ray) Direction() Vector {
	return r.B
}

func (r Ray) PointAtParameter(t float64) Vector {
	return r.A.Add(r.B.MultiplyByScalar(t))
}
