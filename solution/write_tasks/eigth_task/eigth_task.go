package eigth_task

func Eigth(count int, i int, value int) int {
	if value == 1 {
		return count | (value << i)
	} else {
		amount := 1
		for amount <= count {
			amount = (amount << 1)
		}
		amount--
		amount = amount - (1 << i)
		return count & amount
	}
}
