package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"usulroster/internal/helpers"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Nested struct {
	SourcedId string `json:"sourcedId" bson:"sourcedId,omitempty"`
	Type      string `json:"type" bson:"type,omitempty"`
}

// Gets a collection of docs
func GetCollection(
	c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request,
) ([]bson.M, []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter, err := helpers.GetFilters(r.URL.Query(), pf)
	if err != nil {
		log.Error(err)
	}
	options, errP := helpers.GetOptions(r.URL.Query(), pf)
	cur, err := c.Find(
		ctx,
		filter,
		options,
	)
	if err != nil {
		log.Error(err)
	}
	defer cur.Close(ctx)
	totalCount, err := c.CountDocuments(
		ctx,
		filter,
	)
	if err != nil {
		log.Error(err)
	}
	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err)
		}
		results = append(results, result)
	}
	w.Header().Set("X-Total-Count", strconv.FormatInt(totalCount, 10))
	w.Header().Set("Link", helpers.GetLinkHeaders(totalCount, r))
	return results, errP
}

// Gets a specific item based off the sourcedId
func GetDoc(
	c *mongo.Collection, pf []string,
	w http.ResponseWriter, r *http.Request,
) (bson.M, []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "sourcedId", Value: chi.URLParam(r, "id")}}
	options, errP := helpers.GetOption(r.URL.Query(), pf)
	cur := c.FindOne(
		ctx,
		filter,
		options,
	)
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		log.Error(err)
	}
	return result, errP
}

// Upserts a specific item based off the sourcedId
func PutDoc(c *mongo.Collection, data interface{},
	w http.ResponseWriter, r *http.Request) {
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix response
		render.JSON(w, r, err)
		return
	}
	filter := bson.D{primitive.E{Key: "sourcedId", Value: chi.URLParam(r, "id")}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.UpdateOne(
		ctx,
		filter,
		bson.D{primitive.E{Key: "$set", Value: data}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Info(err)
	}
	render.JSON(w, r, res)
}

// Performs an upsert operation to a nested array
func PutNestedDoc(
	c *mongo.Collection, data interface{},
	obj string, field string,
	w http.ResponseWriter, r *http.Request,
) {
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		log.Info(err)
		// TODO: fix response
		render.JSON(w, r, err)
		return
	}
	filter := bson.D{
		primitive.E{Key: "sourcedId", Value: chi.URLParam(r, "id")},
		primitive.E{Key: obj + "." + field, Value: chi.URLParam(r, "subId")},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, _ := c.CountDocuments(
		ctx,
		filter,
	)
	// update
	if count > 0 {
		res, err := c.UpdateOne(
			ctx,
			filter,
			bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: obj + ".$", Value: &data}}}},
		)
		if err != nil {
			// TODO: return 500?
			log.Info(err)
		}
		// TODO: return success update
		render.JSON(w, r, res)
		return
	}
	// insert
	res, err := c.UpdateOne(
		ctx,
		bson.D{primitive.E{Key: "sourcedId", Value: chi.URLParam(r, "id")}},
		bson.D{primitive.E{Key: "$addToSet", Value: bson.D{primitive.E{Key: obj, Value: &data}}}},
	)
	if err != nil {
		log.Info(err)
		// TODO: return 500?
	}
	// TODO: return success insert
	render.JSON(w, r, res)
}
