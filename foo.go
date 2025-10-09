package main

type celsiues float64
type farenheit float64

func F2C(f farenheit) celsiues {
	c := (f - 32) / 1.8
	return celsiues(c)
}

func C2F(c celsiues) farenheit {
	f := (c * 9.0 / 5.0) + 32
	return farenheit(f)
}
