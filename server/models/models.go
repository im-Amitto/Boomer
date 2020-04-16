package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
}

type Cases struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	VehicleNumber   string             `json:"vehiclenumber"`
	AssignedOfficer primitive.ObjectID `json:"assignedofficer"`
	ReportedTime    string             `json:"reportedtime"`
	Image           string             `json:"image"`
	Status          bool               `json:"status"`
}

type PoliceOfficer struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name       string             `json:"name"`
	LastActive string             `json:"lastactive"`
	Status     bool               `json:"status"`
}
