package basics

import "fmt"

type ProductFinder interface {
	FindOneByID(id string) (*Product, error)
}

func DoesProductExist(finder ProductFinder, productID string) string {
	_, err := finder.FindOneByID(productID)
	switch err.(type) {
	case nil:
		return "yes"
	case InvalidIDError:
		return "invalid id"
	case ProductNotFoundError:
		return "no"
	default:
		return "cannot tell"
	}
}

type Product struct {
	ID   string
	Name string
}

type InvalidIDError struct {
	ID string
}

func (e InvalidIDError) Error() string {
	return fmt.Sprintf(`id "%v" is invalid`, e.ID)
}

type ProductNotFoundError struct {
	ID string
}

func (e ProductNotFoundError) Error() string {
	return fmt.Sprintf(`product with id "%v" does not exist`, e.ID)
}
