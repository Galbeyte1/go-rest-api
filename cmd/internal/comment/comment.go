package comment

import (
	"context"
	"fmt"
)

// Comment - a representatoin of the comment
//	structure for our service
type Comment struct {
	ID string
	Slug string // A Slug is the unique identifying part of a web address, typically at the end of the URL
	Body string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error) 
}


// Define the service everyone will interact w/ 
//	to add, delete etc comments
// Service - is the struct on which all our logic
//		will be built on top of
type Service struct {
	Store Store
}

// Constuctor fucntion for service, returning pointer to a service
// Let it be known Go does not have the conecept of constructor function built in
// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
