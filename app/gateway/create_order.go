package gateway

import (
	"fmt"
)

type Gateway struct {
}

func (g *Gateway) CreateOrder(msg string) {
	fmt.Println(msg)
}
