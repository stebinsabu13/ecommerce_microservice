package interfaces

type CartClient interface {
	CreateCart(uint) error
}
