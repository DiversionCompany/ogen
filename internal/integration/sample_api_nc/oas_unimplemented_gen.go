// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DataGetFormat implements dataGetFormat operation.
//
// Retrieve data.
//
// GET /name/{id}/{foo}1234{bar}-{baz}!{kek}
func (UnimplementedHandler) DataGetFormat(ctx context.Context, params DataGetFormatParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// DefaultTest implements defaultTest operation.
//
// POST /defaultTest
func (UnimplementedHandler) DefaultTest(ctx context.Context, req *DefaultTest, params DefaultTestParams) (r int32, _ error) {
	return r, ht.ErrNotImplemented
}

// ErrorGet implements errorGet operation.
//
// Returns error.
//
// GET /error
func (UnimplementedHandler) ErrorGet(ctx context.Context) (r *ErrorStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarGet implements foobarGet operation.
//
// Dumb endpoint for testing things.
//
// GET /foobar
func (UnimplementedHandler) FoobarGet(ctx context.Context, params FoobarGetParams) (r FoobarGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarPost implements foobarPost operation.
//
// Dumb endpoint for testing things.
//
// POST /foobar
func (UnimplementedHandler) FoobarPost(ctx context.Context, req OptPet) (r FoobarPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarPut implements PUT /foobar operation.
//
// PUT /foobar
func (UnimplementedHandler) FoobarPut(ctx context.Context) (r *FoobarPutDef, _ error) {
	return r, ht.ErrNotImplemented
}

// NoAdditionalPropertiesTest implements noAdditionalPropertiesTest operation.
//
// GET /noAdditionalPropertiesTest
func (UnimplementedHandler) NoAdditionalPropertiesTest(ctx context.Context) (r *NoAdditionalPropertiesTest, _ error) {
	return r, ht.ErrNotImplemented
}

// NullableDefaultResponse implements nullableDefaultResponse operation.
//
// GET /nullableDefaultResponse
func (UnimplementedHandler) NullableDefaultResponse(ctx context.Context) (r *NilIntStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// OneofBug implements oneofBug operation.
//
// POST /oneofBug
func (UnimplementedHandler) OneofBug(ctx context.Context, req *OneOfBugs) error {
	return ht.ErrNotImplemented
}

// PatternRecursiveMapGet implements GET /patternRecursiveMap operation.
//
// GET /patternRecursiveMap
func (UnimplementedHandler) PatternRecursiveMapGet(ctx context.Context) (r PatternRecursiveMap, _ error) {
	return r, ht.ErrNotImplemented
}

// PetCreate implements petCreate operation.
//
// Creates pet.
//
// POST /pet
func (UnimplementedHandler) PetCreate(ctx context.Context, req OptPet) (r *Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// PetFriendsNamesByID implements petFriendsNamesByID operation.
//
// Returns names of all friends of pet.
//
// GET /pet/friendNames/{id}
func (UnimplementedHandler) PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) (r []string, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGet implements petGet operation.
//
// Returns pet from the system that the user has access to.
//
// GET /pet
func (UnimplementedHandler) PetGet(ctx context.Context, params PetGetParams) (r PetGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetAvatarByID implements petGetAvatarByID operation.
//
// Returns pet avatar by id.
//
// GET /pet/avatar
func (UnimplementedHandler) PetGetAvatarByID(ctx context.Context, params PetGetAvatarByIDParams) (r PetGetAvatarByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetAvatarByName implements petGetAvatarByName operation.
//
// Returns pet's avatar by name.
//
// GET /pet/{name}/avatar
func (UnimplementedHandler) PetGetAvatarByName(ctx context.Context, params PetGetAvatarByNameParams) (r PetGetAvatarByNameRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetByName implements petGetByName operation.
//
// Returns pet by name from the system that the user has access to.
//
// GET /pet/{name}
func (UnimplementedHandler) PetGetByName(ctx context.Context, params PetGetByNameParams) (r *Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// PetNameByID implements petNameByID operation.
//
// Returns pet name by pet id.
//
// GET /pet/name/{id}
func (UnimplementedHandler) PetNameByID(ctx context.Context, params PetNameByIDParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUpdateNameAliasPost implements POST /pet/updateNameAlias operation.
//
// POST /pet/updateNameAlias
func (UnimplementedHandler) PetUpdateNameAliasPost(ctx context.Context, req OptPetName) (r *PetUpdateNameAliasPostDef, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUpdateNamePost implements POST /pet/updateName operation.
//
// POST /pet/updateName
func (UnimplementedHandler) PetUpdateNamePost(ctx context.Context, req OptString) (r *PetUpdateNamePostDef, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUploadAvatarByID implements petUploadAvatarByID operation.
//
// Uploads pet avatar by id.
//
// POST /pet/avatar
func (UnimplementedHandler) PetUploadAvatarByID(ctx context.Context, req PetUploadAvatarByIDReq, params PetUploadAvatarByIDParams) (r PetUploadAvatarByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RecursiveArrayGet implements GET /recursiveArray operation.
//
// GET /recursiveArray
func (UnimplementedHandler) RecursiveArrayGet(ctx context.Context) (r RecursiveArray, _ error) {
	return r, ht.ErrNotImplemented
}

// RecursiveMapGet implements GET /recursiveMap operation.
//
// GET /recursiveMap
func (UnimplementedHandler) RecursiveMapGet(ctx context.Context) (r *RecursiveMap, _ error) {
	return r, ht.ErrNotImplemented
}

// SecurityTest implements securityTest operation.
//
// GET /securityTest
func (UnimplementedHandler) SecurityTest(ctx context.Context) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// StringIntMapGet implements GET /stringIntMap operation.
//
// GET /stringIntMap
func (UnimplementedHandler) StringIntMapGet(ctx context.Context) (r *StringIntMap, _ error) {
	return r, ht.ErrNotImplemented
}

// TestFloatValidation implements testFloatValidation operation.
//
// POST /testFloatValidation
func (UnimplementedHandler) TestFloatValidation(ctx context.Context, req *TestFloatValidation) error {
	return ht.ErrNotImplemented
}

// TestNullableOneofs implements testNullableOneofs operation.
//
// GET /testNullableOneofs
func (UnimplementedHandler) TestNullableOneofs(ctx context.Context) (r TestNullableOneofsRes, _ error) {
	return r, ht.ErrNotImplemented
}
