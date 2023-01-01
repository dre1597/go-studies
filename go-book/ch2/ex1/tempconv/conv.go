package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + Celsius(FreezingF))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - FreezingF) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - FreezingK)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + Celsius(FreezingK))
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-FreezingK)*9/5 + 32)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + Fahrenheit(FreezingK))
}
