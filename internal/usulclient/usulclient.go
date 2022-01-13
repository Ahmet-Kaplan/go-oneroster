package usulclient

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"usulroster/internal/auth"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func create(desc string, c *mongo.Collection) {

	p, err := randomHex(32)
	if err != nil {
		log.Error(err)
		return
	}
	cred := auth.Creds{
		ClientId:     uuid.New().String(),
		ClientSecret: p,
		Description:  desc,
	}
	b, err := json.Marshal(cred)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	en := encrypt(p)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = c.InsertOne(
		ctx,
		bson.D{primitive.E{Key: "clientId", Value: cred.ClientId},
			primitive.E{Key: "clientSecret", Value: string(en)},
			primitive.E{Key: "description", Value: cred.Description}},
	)
	if err != nil {
		log.Error(err)
	}
	return
}

func encrypt(s string) []byte {
	b := []byte(s)
	en, err := bcrypt.GenerateFromPassword(b, 12)
	if err != nil {
		fmt.Println(err)
		n := make([]byte, 0)
		return n
	}
	return en
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func list(client_desc string, c *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{}
	if client_desc != "" {
		filter = bson.D{primitive.E{Key: "description", Value: client_desc}}
	}
	cur, err := c.Find(
		ctx,
		filter,
		options.Find().SetProjection(bson.D{primitive.E{Key: "clientSecret", Value: 0}, primitive.E{Key: "_id", Value: 0}}),
	)
	if err != nil {
		log.Error(err)
		return
	}
	var creds []auth.Creds
	for cur.Next(ctx) {
		var cred auth.Creds
		err := cur.Decode(&cred)
		if err != nil {
			log.Error(err)
		}
		creds = append(creds, cred)
	}
	b, err := json.Marshal(creds)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func remove(clientId string, c *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := c.DeleteOne(
		ctx,
		bson.D{primitive.E{Key: "clientId", Value: clientId}},
	)
	if err != nil {
		log.Error(err)
	}
	log.Info(res)
}
