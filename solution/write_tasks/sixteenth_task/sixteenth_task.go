package sixteenth_task

// быстрая сортировка
func quickSort(arr []int, left, right int) {
	if left < right {
		//выбор границы подмассива
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	//берем самый правый элемент
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		//если правый элемент больше чем текущий, то увеличиваем i
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	//меняем местами самый правый и то место, где все стоящие слева будут меньше него, а справа - больше
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func Sixteen(mas []int) {
	quickSort(mas, 0, len(mas)-1)
}
