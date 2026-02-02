package processing

import "payment-simulator/internal/cache"

func AssignPoNumber() string{
	return cache.GetNewPoNumber()
}
