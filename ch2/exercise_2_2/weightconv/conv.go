package weightconv

func KgToLb(k Kilogram) Pound { return Pound(k * 2.20462) }

func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.453592) }
