package mongodb

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SingleResult interface {
    Decode(v interface{}) error
}

type UpdateResult interface {
    Decode(v interface{}) error
}

func Connect(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {
	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
					30 * time.Second)
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){
	 
	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func(){

	// client.Disconnect method also has deadline.
	// returns error if any,
	if err := client.Disconnect(ctx); err != nil{
		panic(err)
		 }
	}()
}

// InsertOne inserts document into mongodb
func InsertOne (ctx context.Context, client *mongo.Client, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
 
    // select database and collection ith Client.Database method
    // and Database.Collection method
    collection := client.Database(dataBase).Collection(col)
     
    // InsertOne accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}

// FindOne finds matching entries.
func FindOne(ctx context.Context, client *mongo.Client, dataBase, col string, filter interface{}) (SingleResult, error) {
	collection := client.Database(dataBase).Collection(col)

	singleResult := collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}
	return singleResult, nil
}

// UpdateOne updates an entry.
func UpdateOne(ctx context.Context, client *mongo.Client, dataBase, col string, filter interface{}, fields map[string]interface{}) (error) {
	collection := client.Database(dataBase).Collection(col)
	_, err := collection.UpdateOne(ctx, filter, fields)
	if err != nil { 
        return err                    
    }
	return nil
}

// DeleteOne deletes an entry.
func DeleteOne(ctx context.Context, client *mongo.Client, dataBase, col string, filter interface{}) (error) {
	collection := client.Database(dataBase).Collection(col)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil { 
        return err                    
    }
	return nil
}