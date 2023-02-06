package controllers

func getOrderByStatus(order []Order, want string) []Order {
	s := 0
	for i, v := range order {
		if v.Status == want {
			swap(order, i, s)
			s++
		}
	}
	return order
}

func swap(order []Order, i, j int) {
	order[i], order[j] = order[j], order[i]
}

func getOrderById(order []Order, O_id string) (orderr Order) {
	for _, v := range order {
		if v.Id == O_id {
			orderr = v
		}
	}
	return
}
