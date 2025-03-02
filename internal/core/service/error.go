package service

import "errors"

var (
	ErrProductNotFound         = errors.New("product not found")
	ErrFailedToRetrieveProduct = errors.New("failed to retrieve product")
	ErrFailedToCreateProduct   = errors.New("failed to create product")
	ErrFailedToUpdateProduct   = errors.New("failed to update product")
	ErrFailedToDeleteProduct   = errors.New("failed to delete product")
)

var (
	ErrStockNotFound         = errors.New("stock not found")
	ErrFailedToRetrieveStock = errors.New("failed to retrieve stock")
	ErrFailedToCreateStock   = errors.New("failed to create stock")
	ErrFailedToUpdateStock   = errors.New("failed to update stock")
	ErrFailedToDeleteStock   = errors.New("failed to delete stock")
)
