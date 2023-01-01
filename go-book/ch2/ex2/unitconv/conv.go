package unitconv

func PaToAtm(pa Pascal) Atmospheric {
	return Atmospheric(pa / 101300)
}

func AtmToPa(atm Atmospheric) Pascal {
	return Pascal(atm * 101300)
}
