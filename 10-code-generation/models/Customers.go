/* Auto generated code. Do not modify it */
package models

type Customers []Customer

func (items *Customers) IndexOf(item Customer) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *Customers) Includes(item Customer) bool {
	return items.IndexOf(item) != -1
}

func (items *Customers) Any(predicate func(Customer) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items Customers) Filter(predicate func(item Customer) bool) Customers {
	result := Customers{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
