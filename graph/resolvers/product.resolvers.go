package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"sqlc-rest-api/graph/models"
)

// CreateProduct is the resolver for the CreateProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input models.NewProduct) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - CreateProduct"))
}

// UpdateProduct is the resolver for the UpdateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input models.UpdateProduct) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - UpdateProduct"))
}

// DeleteProduct is the resolver for the DeleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, input models.URIID) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteProduct - DeleteProduct"))
}

// GetProduct is the resolver for the GetProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, input models.URIID) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: GetProduct - GetProduct"))
}