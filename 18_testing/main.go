package sqrt

func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, nil
	}
	if val == 0.0 {
		return 0.0, nil
	}
	quess, epsilon := 1.0, 0.00001
	for i := 0; i < 100; i++ {
		if Abs(quess*quess-val) < epsilon {
			return quess, nil
		}
		quess = (val/quess + quess) / 2.0
	}
	return quess, nil
}

func main() {
	
}
