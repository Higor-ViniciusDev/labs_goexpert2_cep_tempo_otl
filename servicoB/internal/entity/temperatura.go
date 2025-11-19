package entity

type Temperatura struct {
	TempCelsius float64
	TempFar     float64
	TempKelvin  float64
}

func (t *Temperatura) ConverterCelsiusParaKelvin() {
	t.TempKelvin = t.TempCelsius + 273
}
