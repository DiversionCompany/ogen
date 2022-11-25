// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreatePet implements createPet operation.
//
// Creates a new Pet and persists it to storage.
//
// POST /pets
func (UnimplementedHandler) CreatePet(ctx context.Context, req *CreatePetReq) (r CreatePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePetCategories implements createPetCategories operation.
//
// Creates a new Category and attaches it to the Pet.
//
// POST /pets/{id}/categories
func (UnimplementedHandler) CreatePetCategories(ctx context.Context, req *CreatePetCategoriesReq, params CreatePetCategoriesParams) (r CreatePetCategoriesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePetFriends implements createPetFriends operation.
//
// Creates a new Pet and attaches it to the Pet.
//
// POST /pets/{id}/friends
func (UnimplementedHandler) CreatePetFriends(ctx context.Context, req *CreatePetFriendsReq, params CreatePetFriendsParams) (r CreatePetFriendsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePetOwner implements createPetOwner operation.
//
// Creates a new User and attaches it to the Pet.
//
// POST /pets/{id}/owner
func (UnimplementedHandler) CreatePetOwner(ctx context.Context, req *CreatePetOwnerReq, params CreatePetOwnerParams) (r CreatePetOwnerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePet implements deletePet operation.
//
// Deletes the Pet with the requested ID.
//
// DELETE /pets/{id}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) (r DeletePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePetOwner implements deletePetOwner operation.
//
// Delete the attached Owner of the Pet with the given ID.
//
// DELETE /pets/{id}/owner
func (UnimplementedHandler) DeletePetOwner(ctx context.Context, params DeletePetOwnerParams) (r DeletePetOwnerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPet implements listPet operation.
//
// List Pets.
//
// GET /pets
func (UnimplementedHandler) ListPet(ctx context.Context, params ListPetParams) (r ListPetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPetCategories implements listPetCategories operation.
//
// List attached Categories.
//
// GET /pets/{id}/categories
func (UnimplementedHandler) ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (r ListPetCategoriesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPetFriends implements listPetFriends operation.
//
// List attached Friends.
//
// GET /pets/{id}/friends
func (UnimplementedHandler) ListPetFriends(ctx context.Context, params ListPetFriendsParams) (r ListPetFriendsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPet implements readPet operation.
//
// Finds the Pet with the requested ID and returns it.
//
// GET /pets/{id}
func (UnimplementedHandler) ReadPet(ctx context.Context, params ReadPetParams) (r ReadPetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPetOwner implements readPetOwner operation.
//
// Find the attached User of the Pet with the given ID.
//
// GET /pets/{id}/owner
func (UnimplementedHandler) ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (r ReadPetOwnerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePet implements updatePet operation.
//
// Updates a Pet and persists changes to storage.
//
// PATCH /pets/{id}
func (UnimplementedHandler) UpdatePet(ctx context.Context, req *UpdatePetReq, params UpdatePetParams) (r UpdatePetRes, _ error) {
	return r, ht.ErrNotImplemented
}
