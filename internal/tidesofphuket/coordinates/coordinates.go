package coordinates

type Coordinates struct {
	Beach string
	Lat   string
	Lon   string
}

func GetPatongCoordinates() *Coordinates {
	return &Coordinates{
		Beach: "Patong/Karon/Kata",
		Lat:   "7.900432",
		Lon:   "98.296332",
	}
}

func GetMaiKaoCoordinates() *Coordinates {
	return &Coordinates{
		Beach: "MaiKao",
		Lat:   "8.134379",
		Lon:   "98.298922",
	}
}

func GetAonangCoordinates() *Coordinates {
	return &Coordinates{
		Beach: "Aonang",
		Lat:   "8.043317",
		Lon:   "98.809804",
	}
}
