package user

import (
	"context"
	"log"
	"github.com/tutorial/api/proto/_generated/proto/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server defines the user server
type Server struct {
	database *mongo.Client
}

// NewServer returns the user server
func NewServer (ctx context.Context, database *mongo.Client) user.UserServiceServer{
	var server = &Server{
		database: database,
	}
	return server 
}

// CreateUser creates a User and stores in mongodb
func (s *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.Printf("Received: %s", req.GetUser())
	err := s.createUser(ctx, s.database, req)
	if err != nil {
		return nil, err
	} else {
		return &user.CreateUserResponse{User: req.GetUser()}, nil
	}
}

// GetUser fetches User details from mongodb
func (s *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Printf("Received: %s", req.GetName())
	resp, err := s.getUser(ctx, s.database, req.GetName())
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{User: resp}, nil
}

// UpdateUser updates existing user in mongodb
func (s *Server) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	log.Printf("Received: %s", req.GetUser())
	err := s.updateUser(ctx, s.database, req)
	if err != nil {
		return nil, err
	}
	updatedUser, err := s.getUser(ctx, s.database, req.GetUser().Name)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResponse{User: updatedUser}, nil
}

// DeleteUser deleltes the user from mongodb
func (s *Server) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	log.Printf("Received: %s", req.GetName())
	status, err := s.deleteUser(ctx, s.database, req.GetName())
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResponse{Status: status}, nil
}