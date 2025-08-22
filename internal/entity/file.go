package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	ResourceID primitive.ObjectID  `bson:"resource_id"`
	Name       string              `bson:"name"`
	MimeType   string              `bson:"mime_type"`
	Size       int64               `bson:"size"`
	Content    []byte              `bson:"content,omitempty"`
	GridFSID   *primitive.ObjectID `bson:"gridfs_id,omitempty"`
}
