package basics

import "fmt"

const (
	yes        = "yes"
	no         = "no"
	invalidID  = "invalid id"
	cannotTell = "cannot tell"
)

type ProductFinder interface {
	FindOneByID(id string) (*Product, error)
}

// ProductFinder.FindOneByID can return multiple error types.
// E.g.: InvalidIDError, ProductNotFoundError, nil, or any other type.
// Depending on the type of error returned, DoesProductExist should return
// a specific string value.
// If the error is nil, it should return "yes
// If the id is invalid, it should return "invalid id"
// If the product does not exist, it should return "no"
// For any other error type, it should return "cannot tell"
// Hint: you can use the constants defined above.
func DoesProductExist(finder ProductFinder, productID string) string {
	_, err := finder.FindOneByID(productID)
	switch err.(type) {
	case nil:
		return yes
	case InvalidIDError:
		return invalidID
	case ProductNotFoundError:
		return no
	default:
		return cannotTell
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
