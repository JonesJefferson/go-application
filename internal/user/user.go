package user

import (
	"context"
	"encoding/json"
	"fmt"

	userschema "github.com/tutorial/api/internal/user/mongodb/schema"
	"github.com/tutorial/api/mongodb"
	"github.com/tutorial/api/proto/_generated/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser creates a User and stores in mongodb
func (s *Server) createUser(ctx context.Context, mongoClient *mongo.Client, req *user.CreateUserRequest) (error) {
	userInfo := &userschema.User{}
	userInfo.ConvertToSchema(req.GetUser())
	_, err := mongodb.InsertOne(ctx, s.database, "admin", "Users", userInfo)
	// handle the error
    if err != nil {
        return err
    } 
	return nil
}

// GetUser fetches User details from mongodb
func (s *Server) getUser(ctx context.Context, mongoClient *mongo.Client, userName string) (*user.User, error) {
	res, err := mongodb.FindOne(ctx, s.database, "admin", "Users", bson.M{"name": userName})
	if err != nil {
		return nil, err
	}
	var userInfo *userschema.User
	if err = res.(*mongo.SingleResult).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo.ConvertToProto(), nil
}

// UpdateUser updates existing user in mongodb
func (s *Server) updateUser(ctx context.Context, mongoClient *mongo.Client, req *user.UpdateUserRequest) (error) {
	var updateParams map[string]interface{}
	userBytes, _ := json.Marshal(req.GetUser())
    json.Unmarshal(userBytes, &updateParams)
	filter := bson.M{"name": req.GetName()}
    fields := bson.M{"$set": updateParams}
	fmt.Println(fields)
	fmt.Println(filter)
	err := mongodb.UpdateOne(ctx, s.database, "admin", "Users", filter, fields)
	if err != nil {
        return err
    }
	return nil
}

// DeleteUser deleltes the user from mongodb
func (s *Server) deleteUser(ctx context.Context, mongoClient *mongo.Client, userName string) (string, error) {
	err := mongodb.DeleteOne(ctx, s.database, "admin", "Users", bson.M{"name": userName})
	if err != nil {
		return "", err
	}
	return "Deleted record successfully", nil 
}