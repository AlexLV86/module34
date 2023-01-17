package calc

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}

// Выполняет вычисления +, -, * или / с двумя вещественными числами
// Добавить обработку ошибок
func (c calculator) Calculate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return c.add(a, b)
	case "-":
		return c.sub(a, b)
	case "*":
		return c.mul(a, b)
	case "/":
		return c.div(a, b)
	default:
		return 0
	}

}
func (c calculator) add(a, b float64) float64 {
	return (a + b)
}
func (c calculator) sub(a, b float64) float64 {
	return (a - b)
}
func (c calculator) mul(a, b float64) float64 {
	return (a * b)
}
func (c calculator) div(a, b float64) float64 {
	if b == 0.0 {
		return 0
	}
	return (a / b)
}
