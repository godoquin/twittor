package models

type Relation struct {
	UserID         string `bson:"userid" json:"userId,omitempty"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId,omitempty"`
}
