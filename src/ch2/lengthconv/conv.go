package lengthconv

func MToFt(m Meter) Feet {
	return Feet(m / 0.3048)
}

func FtToM(f Feet) Meter {
	return Meter(f * 0.3048)
}
