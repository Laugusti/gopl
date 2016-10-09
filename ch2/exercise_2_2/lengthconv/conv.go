package lengthconv

func MiToKm(mi Mile) Kilometer { return Kilometer(mi * 1.60934) }

func KmToMi(km Kilometer) Mile { return Mile(km * 0.621371) }
