/* Auto generated code. Do not modify it */
package models

type Products []Product

func (items *Products) IndexOf(item Product) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *Products) Includes(item Product) bool {
	return items.IndexOf(item) != -1
}

func (items *Products) Any(predicate func(Product) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items Products) Filter(predicate func(item Product) bool) Products {
	result := Products{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
