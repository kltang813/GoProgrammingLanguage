package weightconv

func KgToLbs(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}

func LbsToKg(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}
