package main

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie oft n in der Liste vorkommt.
func CountElement(list []int, n int) int {
	result := 0
	for _, element := range list {
		if element == n {
			result++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste größer als n sind.
func CountGreaterThan(list []int, n int) int {
	result := 0
	for _, element := range list {
		if element > n {
			result++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste ungleich n sind.
func CountNotN(list []int, n int) int {
	result := 0
	for _, element := range list {
		if element != n {
			result++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste durch n teilbar sind.
func CountDivisors(list []int, n int) int {
	result := 0
	for _, element := range list {
		if element%n == 0 {
			result++
		}
	}
	return result
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in den beiden Listen gleich sind.
func CountEqualElements(list1, list2 []int) int {
	result := 0
	for i := range list1 {
		if i >= len(list2) || list1[i] == list2[i] {
			result++
		}
	}
	if len(list2) > len(list1) {
		result += len(list2) - len(list1)
	}
	return result
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in list1 größer
// als in list 2 sind. Sind die Listen nicht gleich lang, soll die Stelle ignoriert werden.
func CountGreaterElements(list1, list2 []int) int {
	result := 0
	for i := range list1 {
		if i >= len(list2) || list1[i] > list2[i] {
			result++
		}
	}
	return result
}
