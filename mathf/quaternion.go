package mathf

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/goplus/spbase/mathf/impl"
)

type Quaternion struct {
	X, Y, Z, W Float
}

func (q Quaternion) toImpl() impl.Quaternion {
	return *(*impl.Quaternion)(unsafe.Pointer(&q))
}

func (q Quaternion) fromImpl(iq impl.Quaternion) Quaternion {
	return *(*Quaternion)(unsafe.Pointer(&iq))
}

func NewQuaternion(x, y, z, w float64) Quaternion {
	return Quaternion{Float(x), Float(y), Float(z), Float(w)}
}

func NewQuaternionFromAxisAngle(axis Vec3, angle float64) Quaternion {
	halfAngle := angle * 0.5
	s := math.Sin(halfAngle)
	return Quaternion{
		X: Float(axis.X * Float(s)),
		Y: Float(axis.Y * Float(s)),
		Z: Float(axis.Z * Float(s)),
		W: Float(math.Cos(halfAngle)),
	}
}

func (q Quaternion) Add(other Quaternion) Quaternion {
	return q.fromImpl(q.toImpl().Add(other.toImpl()))
}

func (q Quaternion) Sub(other Quaternion) Quaternion {
	return q.fromImpl(q.toImpl().Sub(other.toImpl()))
}

func (q Quaternion) Mul(other Quaternion) Quaternion {
	return q.fromImpl(q.toImpl().Mul(other.toImpl()))
}

func (q Quaternion) Mulf(f float64) Quaternion {
	return q.fromImpl(q.toImpl().Mulf(f))
}

func (q Quaternion) Dot(other Quaternion) float64 {
	return q.toImpl().Dot(other.toImpl())
}

func (q Quaternion) Length() float64 {
	return q.toImpl().Length()
}

func (q Quaternion) LengthSquared() float64 {
	return q.toImpl().LengthSquared()
}

func (q Quaternion) Normalize() Quaternion {
	return q.fromImpl(q.toImpl().Normalized())
}

func (q Quaternion) IsNormalized() bool {
	return q.toImpl().IsNormalized()
}

func (q Quaternion) IsFinite() bool {
	return q.toImpl().IsFinite()
}

func (q Quaternion) Inverse() Quaternion {
	return q.fromImpl(q.toImpl().Inverse())
}

func (q Quaternion) Conjugate() Quaternion {
	return q.fromImpl(q.toImpl().Inverse()) // Conjugate is the same as inverse for unit quaternions
}

func (q Quaternion) Lerp(to Quaternion, weight float64) Quaternion {
	return q.fromImpl(q.toImpl().Slerpni(to.toImpl(), weight))
}

func (q Quaternion) Lerpf(to Quaternion, weight float64) Quaternion {
	return q.fromImpl(q.toImpl().Slerpni(to.toImpl(), weight))
}

func (q Quaternion) Slerp(to Quaternion, weight float64) Quaternion {
	return q.fromImpl(q.toImpl().Slerp(to.toImpl(), weight))
}

func (q Quaternion) Slerpf(to Quaternion, weight float64) Quaternion {
	return q.fromImpl(q.toImpl().Slerp(to.toImpl(), weight))
}

func (q Quaternion) GetEuler() Vec3 {
	angles := q.toImpl().EulerAngles(2) // EulerOrderYXZ = 2
	return NewVec3(float64(angles[0]), float64(angles[1]), float64(angles[2]))
}

func (q Quaternion) RotateVec3(v Vec3) Vec3 {
	// Convert quaternion to basis and use it to rotate the Vec
	qv := NewQuaternion(float64(v.X), float64(v.Y), float64(v.Z), 0)
	qr := q.Mul(qv).Mul(q.Conjugate())
	return NewVec3(float64(qr.X), float64(qr.Y), float64(qr.Z))
}

func (q Quaternion) Neg() Quaternion {
	return q.fromImpl(q.toImpl().Neg())
}

func (q Quaternion) String() string {
	return fmt.Sprintf("[x=%g, y=%g, z=%g, w=%g]", q.X, q.Y, q.Z, q.W)
}
