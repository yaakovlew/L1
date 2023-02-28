package eigth_task

func Eigth(count int64, i int, value int64) int64 {
	if value == 1 {
		return count | (value << i)
	} else {
		var amount int64 = 1
		for amount <= count {
			amount = (amount << 1)
		}
		amount--
		amount = amount - (1 << i)
		return count & amount
	}
}
