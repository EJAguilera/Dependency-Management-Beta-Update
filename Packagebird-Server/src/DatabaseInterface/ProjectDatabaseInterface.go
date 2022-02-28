package DatabaseInterface

import (
	"context"
	"fmt"
	"log"
	"os"
	configFile "packagebird-server/src/config"
	structures "packagebird-server/src/structures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProject(client mongo.Client, name string, description string) (bool, error) {
	// First establish connection to database and collection
	collection := client.Database("packagebird").Collection("project")

	// Check that project isn't already present in the collection
	var result structures.Project
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: name}}).Decode(&result)

	// !!! This conditional needs fixed, not elegant at all.
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Project source directory path
			projectSourcePath := configFile.Config.ProjectSourcePath
			nameSourcePath := fmt.Sprintf("%v/%v", projectSourcePath, name)
			if err := os.Mkdir(nameSourcePath, 0755); err != nil {
				log.Printf("Failed to create project directory at path specified\n%v", err)
				return false, err
			}

			// If no project found with name, create new project
			log.Printf("Creating new project with name %v in database...", name)

			newProject := structures.Project{Name: name, Description: description, UUID: "", LatestVersion: 0, Packages: nil, SourceFile: nameSourcePath}
			_, insertErr := collection.InsertOne(context.TODO(), newProject)

			if insertErr != nil {
				log.Printf("Error encountered with inserting new project into database...")
				return false, insertErr
			}

			return true, nil
		} else {
			// Unknown error encountered
			log.Printf("Error encountered searching for project in database...")
			return false, err
		}
	} else {
		// Project already exist
		log.Printf("Project with name %v already exist in database...", name)
		return false, nil
	}
}

func LookupProject(client mongo.Client, name string, description string) (bool, error) {
	collection := client.Database("packagebird").Collection("project")

	log.Printf("Looking up project with name: %v", name)
	_, err := collection.Find(context.TODO(), bson.D{primitive.E{Key: "name", Value: name}})

	if err == mongo.ErrNoDocuments {
		log.Printf("Project with name %v is not in the database", name)
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		// Find error encountered
		log.Printf("Error encountered searching for project in database")
		return false, err
	}
}

func UpdateProject(c mongo.Client, name string, dependencies []string) error {
	collection := c.Database("packagebird").Collection("project")

	var result structures.Project
	filter := bson.M{"name": name}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	result.Packages = append(result.Packages, dependencies...)

	_, err = collection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.M{"dependencies": result.Packages}}})
	if err != nil {
		log.Printf("%v", err)
		return err
	} else {
		log.Printf("Updated dependencies for project: %v", name)
		return nil
	}
}

func IncrementProjectVersion(c mongo.Client, name string) error {
	collection := c.Database("packagebird").Collection("project")
	project, err := GetProject(*collection, name)
	if err != nil {
		log.Printf("Cannot find and retrieve project %v from MongoDB", name)
		return err
	}

	version := project.LatestVersion + 1
	_, err = collection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.M{"version": version}}})

	if err != nil {
		log.Printf("Error incrementing project %v version: %v", name, err)
		return err
	}
	return nil
}

func GetProject(c mongo.Collection, name string) (*structures.Project, error) {
	var result structures.Project
	filter := bson.M{"name": name}
	if err := c.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
