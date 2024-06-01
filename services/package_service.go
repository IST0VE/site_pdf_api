package services

import (
	"context"
	"time"

	"github.com/IST0VE/site_pdf_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePackage(pkg *models.Package) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return packageCollection.InsertOne(ctx, pkg)
}

func GetPackageByID(pkgID primitive.ObjectID) (*models.Package, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var pkg models.Package
	err := packageCollection.FindOne(ctx, bson.M{"_id": pkgID}).Decode(&pkg)
	return &pkg, err
}

func GetAllPackages() ([]*models.Package, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var packages []*models.Package
	cursor, err := packageCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pkg models.Package
		if err := cursor.Decode(&pkg); err != nil {
			return nil, err
		}
		packages = append(packages, &pkg)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}

func UpdatePackage(pkgID primitive.ObjectID, updatePackage *models.Package) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{
		"$set": updatePackage,
	}
	return packageCollection.UpdateByID(ctx, pkgID, update)
}

func DeletePackage(pkgID primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return packageCollection.DeleteOne(ctx, bson.M{"_id": pkgID})
}
