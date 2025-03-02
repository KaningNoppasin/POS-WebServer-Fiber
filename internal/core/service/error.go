package service

import "errors"

var (
	ErrProductNotFound         = errors.New("product not found")
	ErrFailedToRetrieveProduct = errors.New("failed to retrieve product")
	ErrFailedToCreateProduct   = errors.New("failed to create product")
	ErrFailedToUpdateProduct   = errors.New("failed to update product")
	ErrFailedToDeleteProduct   = errors.New("failed to delete product")
)
