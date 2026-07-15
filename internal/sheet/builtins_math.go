package sheet

import "math"

// mathResult wraps a numeric result, mapping a NaN or infinite value (a domain
// error or overflow, e.g. SQRT of a negative) to #NUM!, per Excel.
func mathResult(x floatVal) Value {
	f := float64(x)
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return errorValue(ErrNum)
	}
	return numberValue(x)
}

// unaryNumeric adapts a single-operand numeric op into a builtin, propagating a
// non-numeric operand as #VALUE! and a domain error as #NUM!.
func unaryNumeric(op func(x floatVal) floatVal) func(args []Value) Value {
	return func(args []Value) Value {
		n, v := args[0].asNumber()
		if v.isError() {
			return v
		}
		return mathResult(op(floatVal(n)))
	}
}

// binaryNumeric adapts a two-operand numeric op into a builtin, propagating a
// non-numeric operand as #VALUE! and a domain error as #NUM!.
func binaryNumeric(op func(a, b floatVal) floatVal) func(args []Value) Value {
	return func(args []Value) Value {
		a, va := args[0].asNumber()
		if va.isError() {
			return va
		}
		b, vb := args[1].asNumber()
		if vb.isError() {
			return vb
		}
		return mathResult(op(floatVal(a), floatVal(b)))
	}
}

// Thin bridges to the standard library, typed in the engine's floatVal (math
// operates on float64).
func mFloor(x floatVal) floatVal { return floatVal(math.Floor(float64(x))) }
func mTrunc(x floatVal) floatVal { return floatVal(math.Trunc(float64(x))) }
func mSqrt(x floatVal) floatVal  { return floatVal(math.Sqrt(float64(x))) }
func mExp(x floatVal) floatVal   { return floatVal(math.Exp(float64(x))) }
func mLn(x floatVal) floatVal    { return floatVal(math.Log(float64(x))) }
func mLog10(x floatVal) floatVal { return floatVal(math.Log10(float64(x))) }
func mSin(x floatVal) floatVal   { return floatVal(math.Sin(float64(x))) }
func mCos(x floatVal) floatVal   { return floatVal(math.Cos(float64(x))) }
func mTan(x floatVal) floatVal   { return floatVal(math.Tan(float64(x))) }
func mAsin(x floatVal) floatVal  { return floatVal(math.Asin(float64(x))) }
func mAcos(x floatVal) floatVal  { return floatVal(math.Acos(float64(x))) }
func mAtan(x floatVal) floatVal  { return floatVal(math.Atan(float64(x))) }
func mSinh(x floatVal) floatVal  { return floatVal(math.Sinh(float64(x))) }
func mCosh(x floatVal) floatVal  { return floatVal(math.Cosh(float64(x))) }
func mTanh(x floatVal) floatVal  { return floatVal(math.Tanh(float64(x))) }

// mPow raises a to the b-th power.
func mPow(a, b floatVal) floatVal { return floatVal(math.Pow(float64(a), float64(b))) }

// sqrtPi is √(x·π), matching Excel SQRTPI.
func sqrtPi(x floatVal) floatVal { return floatVal(math.Sqrt(float64(x) * math.Pi)) }

// toDegrees and toRadians convert between radians and degrees.
func toDegrees(radians floatVal) floatVal { return radians * 180 / floatVal(math.Pi) }
func toRadians(degrees floatVal) floatVal { return degrees * floatVal(math.Pi) / 180 }

// atan2Excel is ATAN2 in Excel's argument order (x, y), the reverse of Go's
// math.Atan2(y, x).
func atan2Excel(x, y floatVal) floatVal { return floatVal(math.Atan2(float64(y), float64(x))) }

// sign is -1, 0, or 1 by the sign of x.
func sign(x floatVal) floatVal {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// fnPi is the constant π (no arguments).
func fnPi([]Value) Value { return numberValue(floatVal(math.Pi)) }

// fnLog is log(x); a second argument sets the base (default 10).
func fnLog(args []Value) Value {
	x, v := args[0].asNumber()
	if v.isError() {
		return v
	}
	if len(args) == 1 {
		return mathResult(floatVal(math.Log10(x)))
	}
	base, bv := args[1].asNumber()
	if bv.isError() {
		return bv
	}
	return mathResult(floatVal(math.Log(x) / math.Log(base)))
}

// fnQuotient is the integer part of a / b; a zero divisor is #DIV/0!.
func fnQuotient(args []Value) Value {
	a, va := args[0].asNumber()
	if va.isError() {
		return va
	}
	b, vb := args[1].asNumber()
	if vb.isError() {
		return vb
	}
	if b == 0 {
		return errorValue(ErrDiv)
	}
	return numberValue(floatVal(math.Trunc(a / b)))
}

// fnProduct multiplies its numeric operands; an empty set is 1.
func fnProduct(args []Value) Value {
	nums, bad, ok := numerics(args)
	if !ok {
		return bad
	}
	product := 1.0
	for _, n := range nums {
		product *= n
	}
	return numberValue(floatVal(product))
}

// fnSumsq totals the squares of its numeric operands.
func fnSumsq(args []Value) Value {
	nums, bad, ok := numerics(args)
	if !ok {
		return bad
	}
	total := 0.0
	for _, n := range nums {
		total += n * n
	}
	return numberValue(floatVal(total))
}
