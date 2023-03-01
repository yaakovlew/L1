package eigth_task

func Eigth(count int64, i int, value int64) int64 {
	//Если нужно поменять на 1, то тогда можно использовать "побитовое или" с маской 1....0
	if value == 1 {
		return count | (value << i)
	} else {
		//Если нужно поменять на 0, то тогда можно использовать "побитовое и" с маской 1..0..1,
		//чтобы не изменить другие значения
		var amount int64 = 1
		for amount <= count {
			amount = (amount << 1)
		}
		amount--
		amount = amount - (1 << i)
		return count & amount
	}
}
