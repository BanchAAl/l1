package sort

//QSort реализация быстрой сортировки
func QSort(ar []int) {
	//дальше нечего разбивать
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	// рубим массив на две части
	QSort(ar[:split])
	QSort(ar[split:])
}

//partition делим массив на части. опорный элемент выбираем посередине.
func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		//нашли левый элемент больше опорного и правый меньше опорного - надо поменять их местами
		swap(ar, left, right)
	}
}

//swap меняем элементы местами.
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
