package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string for localhost mongoDB
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "boomer"

// case object/instance
var caseCollection *mongo.Collection

// officer object/instance
var officerCollection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	caseCollection = client.Database(dbName).Collection("cases")
	officerCollection = client.Database(dbName).Collection("officer")

	fmt.Println("Collection instance created!")
}

// GetAllCase get all the case route
func GetAllCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := bson.D{}
	payload := getAll(caseCollection, query)
	if payload == nil {
		payload = []bson.M{}
	}
	json.NewEncoder(w).Encode(payload)
}

// GetAllOfficer get all the officer route
func GetAllOfficer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := bson.D{}
	payload := getAll(officerCollection, query)
	if payload == nil {
		payload = []bson.M{}
	}
	json.NewEncoder(w).Encode(payload)
}

// CreateCase create case route
func CreateCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormValue returns the first value for the given key `myFile`
	vehicleNumber := r.FormValue("vehicleNumber")
	var caseDetail models.Cases = models.Cases{}
	caseDetail.VehicleNumber = vehicleNumber
	if !caseExist(caseDetail) {
		// FormFile returns the first file for the given key `myFile`
		file, _, err := r.FormFile("myFile")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		// Create a temporary file within our temp-images directory that follows a particular naming pattern
		tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
		if err != nil {
			log.Fatal(err)
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		caseDetail.Image = tempFile.Name()
		caseDetail.ReportedTime = time.Now().String()
		insertOneCase(caseDetail)
		json.NewEncoder(w).Encode(bson.M{"msg": "Case Registered Successfully"})
	} else {
		json.NewEncoder(w).Encode(bson.M{"msg": "Case Already Exist"})

	}
}

// CreateOfficer create officer route
func CreateOfficer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var officer models.PoliceOfficer
	_ = json.NewDecoder(r.Body).Decode(&officer)
	insertOneOfficer(officer)
	json.NewEncoder(w).Encode(bson.M{"msg": "Officer added successfully"})
}

// CaseComplete update case status
func CaseComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var caseDetail models.Cases
	_ = json.NewDecoder(r.Body).Decode(&caseDetail)
	caseComplete(caseDetail)
	json.NewEncoder(w).Encode(bson.M{"msg": "Case Successfully Solved"})
}

// get all from the passed DB collection and return it
func getAll(collection *mongo.Collection, query bson.D) []primitive.M {
	cur, err := collection.Find(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one policeOfficer in the DB
func insertOneOfficer(officer models.PoliceOfficer) {

	insertResult, err := officerCollection.InsertOne(context.Background(), officer)

	if err != nil {
		log.Fatal(err)
	}

	officer.ID = insertResult.InsertedID.(primitive.ObjectID)
	caseToBeAssigned := getFreeCase()
	if !primitive.ObjectID.IsZero(caseToBeAssigned.ID) {
		updateOfficerStatus(officer.ID, true)
		assignOfficer(caseToBeAssigned.ID, officer.ID)
	}
}

// Insert one case in the DB
func insertOneCase(caseDetail models.Cases) {

	insertResult, err := caseCollection.InsertOne(context.Background(), caseDetail)

	if err != nil {
		log.Fatal(err)
	}
	caseDetail.ID = insertResult.InsertedID.(primitive.ObjectID)

	assignedPoliceOfficer := getFreeOfficer()
	if !primitive.ObjectID.IsZero(assignedPoliceOfficer.ID) {
		assignOfficer(caseDetail.ID, assignedPoliceOfficer.ID)
		updateOfficerStatus(assignedPoliceOfficer.ID, true)
	}
}

// case complete method, update case status and unlink officer
func caseComplete(completedCase models.Cases) {
	updateCaseStatus(completedCase.ID, true)
	updateOfficerStatus(completedCase.AssignedOfficer, false)

	caseToBeAssigned := getFreeCase()
	if !primitive.ObjectID.IsZero(caseToBeAssigned.ID) {
		updateOfficerStatus(completedCase.AssignedOfficer, true)
		assignOfficer(caseToBeAssigned.ID, completedCase.AssignedOfficer)
	}
}

// return officer who is not assigned to any case or null
func getFreeOfficer() models.PoliceOfficer {
	cur, err := officerCollection.Find(context.Background(), bson.D{{"status", false}})
	if err != nil {
		log.Fatal(err)
	}

	var results models.PoliceOfficer
	var lastActive string = "5020-04-16 01:19:56.211652991 +0530 IST m=+6.330077699"
	for cur.Next(context.Background()) {
		var result models.PoliceOfficer
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		x := result.LastActive
		if lastActive > x {
			lastActive = x
			results = result
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// return case which does not have any officer assigned
func getFreeCase() models.Cases {
	cur, err := caseCollection.Find(context.Background(), bson.D{{"assignedofficer", primitive.NilObjectID}})
	if err != nil {
		log.Fatal(err)
	}

	var results models.Cases
	var reportedTime string = "5020-04-16 01:19:56.211652991 +0530 IST m=+6.330077699"
	for cur.Next(context.Background()) {
		var result models.Cases
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		x := result.ReportedTime
		if reportedTime > x {
			reportedTime = x
			results = result
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func updateOfficerStatus(id primitive.ObjectID, status bool) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": status, "lastactive": time.Now().String()}}
	_, err := officerCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func updateCaseStatus(id primitive.ObjectID, status bool) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := caseCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func assignOfficer(id primitive.ObjectID, id2 primitive.ObjectID) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"assignedofficer": id2}}
	_, err := caseCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func caseExist(caseDetail models.Cases) bool {
	query := bson.D{{"vehiclenumber", caseDetail.VehicleNumber}, {"status", false}}
	isADuplicate := getAll(caseCollection, query)

	return isADuplicate != nil
}
