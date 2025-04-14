package services

// esse serviço vai realizar os seguinte:
// 1. Definir uma base, tendo 10 como default
// 2.
type IExpoente interface{}

type expoente struct {
	base     int
	expoente int
}

func NewExpoente(base int, expoente int) IExpoente {
	return &expoente{
		base:     base,
		expoente: expoente,
	}
}
