package volumeconv

func GalToL(gal Gallon) Liter { return Liter(gal * 3.78541) }

func LToGal(l Liter) Gallon { return Gallon(l * 0.264172) }
