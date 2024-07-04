package main

import (
	"go-minimalists/features/product"
	"go-minimalists/features/user"
)

var migrations = []interface{}{
	&user.User{},
	&product.Product{},
}
