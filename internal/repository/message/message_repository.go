package repository

import (
	"context"
	"errors"
	"github.com/kavkaco/Kavka-Core/database"
	"github.com/kavkaco/Kavka-Core/internal/domain/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrChatNotFound = errors.New("chat not found")
	ErrMsgNotFound  = errors.New("message not found")
	ErrNoAccess     = errors.New("no access")
)

type messageRepository struct {
	chatsCollection *mongo.Collection
}

func NewMessageRepository(db *mongo.Database) message.Repository {
	return &messageRepository{db.Collection(database.ChatsCollection)}
}

func (repo *messageRepository) Insert(chatID primitive.ObjectID, msg *message.Message) (*message.Message, error) {
	filter := bson.M{"id": chatID}
	update := bson.M{"$push": bson.M{"messages": msg}}

	_, err := repo.chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if database.IsRowExistsError(err) {
			return nil, ErrChatNotFound
		}
		return nil, err
	}

	return msg, nil
}

func (repo *messageRepository) Update(chatID primitive.ObjectID, messageID primitive.ObjectID, fieldsToUpdate bson.M) error {
	filter := bson.M{"id": chatID, "messages.id": messageID}
	update := bson.M{"$set": fieldsToUpdate}

	_, err := repo.chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if database.IsRowExistsError(err) {
			return ErrChatNotFound
		}

		return err
	}

	return nil
}

func (repo *messageRepository) Delete(chatID primitive.ObjectID, messageID primitive.ObjectID) error {
	filter := bson.M{"id": chatID}
	update := bson.M{"$pull": bson.M{"messages": bson.M{"id": messageID}}}

	_, err := repo.chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if database.IsRowExistsError(err) {
			return ErrChatNotFound
		}

		return err
	}

	return nil
}
