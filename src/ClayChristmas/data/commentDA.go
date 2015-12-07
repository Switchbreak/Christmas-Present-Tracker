package data

import (
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

func GetComments(appContext appengine.Context, partyID string, personID string) ([]model.Comment, error) {
	query := datastore.NewQuery("Comment").Ancestor(getParentKey(appContext, partyID, personID)).Order( "Date" )

	var comments []model.Comment
	keys, err := query.GetAll(appContext, &comments)
	if err != nil {
		return nil, err
	}

	for index, key := range keys {
		comments[index].ID = key.StringID()
	}

	return comments, err
}

func CreateComment(appContext appengine.Context, partyID string, personID string, comment *model.Comment) (*model.Comment, error) {
	key := datastore.NewIncompleteKey(appContext, "Comment", getParentKey(appContext, partyID, personID))

	key, err := datastore.Put(appContext, key, comment)
	comment.ID = key.StringID()

	return comment, err
}

func GetComment(appContext appengine.Context, partyID string, personID string, commentID string) (*model.Comment, error) {
	key := datastore.NewKey(appContext, "Comment", commentID, 0, getParentKey(appContext, partyID, personID))

	var comment model.Comment
	err := datastore.Get(appContext, key, &comment)

	comment.ID = commentID
	return &comment, err
}

func UpdateComment(appContext appengine.Context, partyID string, personID string, comment *model.Comment) (*model.Comment, error) {
	key := datastore.NewKey(appContext, "Comment", comment.ID, 0, getParentKey(appContext, partyID, personID))

	key, err := datastore.Put(appContext, key, comment)
	comment.ID = key.StringID()

	return comment, err
}

func DeleteComment(appContext appengine.Context, partyID string, personID string, comment *model.Comment) error {
	key := datastore.NewKey(appContext, "Comment", comment.ID, 0, getParentKey(appContext, partyID, personID))

	return datastore.Delete(appContext, key)
}

func getParentKey(appContext appengine.Context, partyID string, personID string) *datastore.Key {
	parentKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	if personID != "" {
		parentKey = datastore.NewKey(appContext, "InvitedPerson", personID, 0, nil)
	}
	
	return parentKey
}
