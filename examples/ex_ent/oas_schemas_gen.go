// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"
)

type CreatePetCategoriesReq struct {
	Name string `json:"name"`
	Pets []int  `json:"pets"`
}

// GetName returns the value of Name.
func (s *CreatePetCategoriesReq) GetName() string {
	return s.Name
}

// GetPets returns the value of Pets.
func (s *CreatePetCategoriesReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *CreatePetCategoriesReq) SetName(val string) {
	s.Name = val
}

// SetPets sets the value of Pets.
func (s *CreatePetCategoriesReq) SetPets(val []int) {
	s.Pets = val
}

type CreatePetFriendsReq struct {
	Name       string      `json:"name"`
	Weight     OptInt      `json:"weight"`
	Birthday   OptDateTime `json:"birthday"`
	Categories []int       `json:"categories"`
	Owner      int         `json:"owner"`
	Friends    []int       `json:"friends"`
}

// GetName returns the value of Name.
func (s *CreatePetFriendsReq) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *CreatePetFriendsReq) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *CreatePetFriendsReq) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *CreatePetFriendsReq) GetCategories() []int {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *CreatePetFriendsReq) GetOwner() int {
	return s.Owner
}

// GetFriends returns the value of Friends.
func (s *CreatePetFriendsReq) GetFriends() []int {
	return s.Friends
}

// SetName sets the value of Name.
func (s *CreatePetFriendsReq) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *CreatePetFriendsReq) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *CreatePetFriendsReq) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *CreatePetFriendsReq) SetCategories(val []int) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *CreatePetFriendsReq) SetOwner(val int) {
	s.Owner = val
}

// SetFriends sets the value of Friends.
func (s *CreatePetFriendsReq) SetFriends(val []int) {
	s.Friends = val
}

type CreatePetOwnerReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pets []int  `json:"pets"`
}

// GetName returns the value of Name.
func (s *CreatePetOwnerReq) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *CreatePetOwnerReq) GetAge() int {
	return s.Age
}

// GetPets returns the value of Pets.
func (s *CreatePetOwnerReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *CreatePetOwnerReq) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *CreatePetOwnerReq) SetAge(val int) {
	s.Age = val
}

// SetPets sets the value of Pets.
func (s *CreatePetOwnerReq) SetPets(val []int) {
	s.Pets = val
}

type CreatePetReq struct {
	Name       string      `json:"name"`
	Weight     OptInt      `json:"weight"`
	Birthday   OptDateTime `json:"birthday"`
	Categories []int       `json:"categories"`
	Owner      int         `json:"owner"`
	Friends    []int       `json:"friends"`
}

// GetName returns the value of Name.
func (s *CreatePetReq) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *CreatePetReq) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *CreatePetReq) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *CreatePetReq) GetCategories() []int {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *CreatePetReq) GetOwner() int {
	return s.Owner
}

// GetFriends returns the value of Friends.
func (s *CreatePetReq) GetFriends() []int {
	return s.Friends
}

// SetName sets the value of Name.
func (s *CreatePetReq) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *CreatePetReq) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *CreatePetReq) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *CreatePetReq) SetCategories(val []int) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *CreatePetReq) SetOwner(val int) {
	s.Owner = val
}

// SetFriends sets the value of Friends.
func (s *CreatePetReq) SetFriends(val []int) {
	s.Friends = val
}

// DeletePetNoContent is response for DeletePet operation.
type DeletePetNoContent struct{}

func (*DeletePetNoContent) deletePetRes() {}

// DeletePetOwnerNoContent is response for DeletePetOwner operation.
type DeletePetOwnerNoContent struct{}

func (*DeletePetOwnerNoContent) deletePetOwnerRes() {}

type ListPetCategoriesOKApplicationJSON []PetCategoriesList

func (ListPetCategoriesOKApplicationJSON) listPetCategoriesRes() {}

type ListPetFriendsOKApplicationJSON []PetFriendsList

func (ListPetFriendsOKApplicationJSON) listPetFriendsRes() {}

type ListPetOKApplicationJSON []PetList

func (ListPetOKApplicationJSON) listPetRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pet_CategoriesCreate
type PetCategoriesCreate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *PetCategoriesCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCategoriesCreate) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *PetCategoriesCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCategoriesCreate) SetName(val string) {
	s.Name = val
}

func (*PetCategoriesCreate) createPetCategoriesRes() {}

// Ref: #/components/schemas/Pet_CategoriesList
type PetCategoriesList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *PetCategoriesList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCategoriesList) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *PetCategoriesList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCategoriesList) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/PetCreate
type PetCreate struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Weight     OptInt                `json:"weight"`
	Birthday   OptDateTime           `json:"birthday"`
	Categories []PetCreateCategories `json:"categories"`
	Owner      PetCreateOwner        `json:"owner"`
}

// GetID returns the value of ID.
func (s *PetCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreate) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetCreate) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetCreate) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *PetCreate) GetCategories() []PetCreateCategories {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *PetCreate) GetOwner() PetCreateOwner {
	return s.Owner
}

// SetID sets the value of ID.
func (s *PetCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreate) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetCreate) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetCreate) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *PetCreate) SetCategories(val []PetCreateCategories) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *PetCreate) SetOwner(val PetCreateOwner) {
	s.Owner = val
}

func (*PetCreate) createPetRes() {}

// Ref: #/components/schemas/PetCreate_Categories
type PetCreateCategories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *PetCreateCategories) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreateCategories) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *PetCreateCategories) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreateCategories) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/PetCreate_Owner
type PetCreateOwner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *PetCreateOwner) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreateOwner) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *PetCreateOwner) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *PetCreateOwner) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreateOwner) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *PetCreateOwner) SetAge(val int) {
	s.Age = val
}

// Ref: #/components/schemas/Pet_FriendsCreate
type PetFriendsCreate struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetFriendsCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetFriendsCreate) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetFriendsCreate) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetFriendsCreate) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetFriendsCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetFriendsCreate) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetFriendsCreate) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetFriendsCreate) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

func (*PetFriendsCreate) createPetFriendsRes() {}

// Ref: #/components/schemas/Pet_FriendsList
type PetFriendsList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetFriendsList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetFriendsList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetFriendsList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetFriendsList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetFriendsList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetFriendsList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetFriendsList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetFriendsList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/PetList
type PetList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/Pet_OwnerCreate
type PetOwnerCreate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *PetOwnerCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetOwnerCreate) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *PetOwnerCreate) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *PetOwnerCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetOwnerCreate) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *PetOwnerCreate) SetAge(val int) {
	s.Age = val
}

func (*PetOwnerCreate) createPetOwnerRes() {}

// Ref: #/components/schemas/Pet_OwnerRead
type PetOwnerRead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *PetOwnerRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetOwnerRead) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *PetOwnerRead) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *PetOwnerRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetOwnerRead) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *PetOwnerRead) SetAge(val int) {
	s.Age = val
}

func (*PetOwnerRead) readPetOwnerRes() {}

// Ref: #/components/schemas/PetRead
type PetRead struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetRead) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetRead) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetRead) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetRead) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetRead) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetRead) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

func (*PetRead) readPetRes() {}

// Ref: #/components/schemas/PetUpdate
type PetUpdate struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetUpdate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetUpdate) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetUpdate) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetUpdate) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetUpdate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetUpdate) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetUpdate) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetUpdate) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

func (*PetUpdate) updatePetRes() {}

type R400 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// GetCode returns the value of Code.
func (s *R400) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R400) GetStatus() string {
	return s.Status
}

// SetCode sets the value of Code.
func (s *R400) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R400) SetStatus(val string) {
	s.Status = val
}

func (*R400) createPetCategoriesRes() {}
func (*R400) createPetFriendsRes()    {}
func (*R400) createPetOwnerRes()      {}
func (*R400) createPetRes()           {}
func (*R400) deletePetOwnerRes()      {}
func (*R400) deletePetRes()           {}
func (*R400) listPetCategoriesRes()   {}
func (*R400) listPetFriendsRes()      {}
func (*R400) listPetRes()             {}
func (*R400) readPetOwnerRes()        {}
func (*R400) readPetRes()             {}
func (*R400) updatePetRes()           {}

type R404 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// GetCode returns the value of Code.
func (s *R404) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R404) GetStatus() string {
	return s.Status
}

// SetCode sets the value of Code.
func (s *R404) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R404) SetStatus(val string) {
	s.Status = val
}

func (*R404) deletePetOwnerRes()    {}
func (*R404) deletePetRes()         {}
func (*R404) listPetCategoriesRes() {}
func (*R404) listPetFriendsRes()    {}
func (*R404) listPetRes()           {}
func (*R404) readPetOwnerRes()      {}
func (*R404) readPetRes()           {}
func (*R404) updatePetRes()         {}

type R409 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// GetCode returns the value of Code.
func (s *R409) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R409) GetStatus() string {
	return s.Status
}

// SetCode sets the value of Code.
func (s *R409) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R409) SetStatus(val string) {
	s.Status = val
}

func (*R409) createPetCategoriesRes() {}
func (*R409) createPetFriendsRes()    {}
func (*R409) createPetOwnerRes()      {}
func (*R409) createPetRes()           {}

type R500 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// GetCode returns the value of Code.
func (s *R500) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R500) GetStatus() string {
	return s.Status
}

// SetCode sets the value of Code.
func (s *R500) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R500) SetStatus(val string) {
	s.Status = val
}

func (*R500) createPetCategoriesRes() {}
func (*R500) createPetFriendsRes()    {}
func (*R500) createPetOwnerRes()      {}
func (*R500) createPetRes()           {}
func (*R500) deletePetOwnerRes()      {}
func (*R500) deletePetRes()           {}
func (*R500) listPetCategoriesRes()   {}
func (*R500) listPetFriendsRes()      {}
func (*R500) listPetRes()             {}
func (*R500) readPetOwnerRes()        {}
func (*R500) readPetRes()             {}
func (*R500) updatePetRes()           {}

type UpdatePetReq struct {
	Name       string      `json:"name"`
	Weight     OptInt      `json:"weight"`
	Birthday   OptDateTime `json:"birthday"`
	Categories []int       `json:"categories"`
	Owner      int         `json:"owner"`
	Friends    []int       `json:"friends"`
}

// GetName returns the value of Name.
func (s *UpdatePetReq) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *UpdatePetReq) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *UpdatePetReq) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *UpdatePetReq) GetCategories() []int {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *UpdatePetReq) GetOwner() int {
	return s.Owner
}

// GetFriends returns the value of Friends.
func (s *UpdatePetReq) GetFriends() []int {
	return s.Friends
}

// SetName sets the value of Name.
func (s *UpdatePetReq) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *UpdatePetReq) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *UpdatePetReq) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *UpdatePetReq) SetCategories(val []int) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *UpdatePetReq) SetOwner(val int) {
	s.Owner = val
}

// SetFriends sets the value of Friends.
func (s *UpdatePetReq) SetFriends(val []int) {
	s.Friends = val
}
