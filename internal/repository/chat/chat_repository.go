package repository

import (
	"context"
	"errors"

	"github.com/kavkaco/Kavka-Core/database"
	"github.com/kavkaco/Kavka-Core/internal/domain/chat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrChatNotFound      = errors.New("chat not found")
	ErrChatAlreadyExists = errors.New("chat already exists")
)

type repository struct {
	chatsCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) chat.Repository {
	return &repository{db.Collection(database.ChatsCollection)}
}

func (repo *repository) Create(newChat chat.Chat) (*chat.Chat, error) {
	_, err := repo.chatsCollection.InsertOne(context.Background(), newChat)
	if err != nil {
		return nil, err
	}

	return &newChat, nil
}

func (repo *repository) Destroy(chatID primitive.ObjectID) error {
	filter := bson.M{"chat_id": chatID}

	_, err := repo.chatsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Where(filter bson.M) ([]chat.Chat, error) {
	cursor, err := repo.chatsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var chats []chat.Chat

	decodeErr := cursor.All(context.Background(), &chats)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return chats, nil
}

func (repo *repository) findBy(filter bson.M) (*chat.Chat, error) {
	result, err := repo.Where(filter)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		user := result[len(result)-1]

		return &user, nil
	}

	return nil, ErrChatNotFound
}

func (repo *repository) GetUserChats(userStaticID primitive.ObjectID) ([]chat.Chat, error) {
	filter := bson.M{
		"chat_detail.members": bson.M{
			"$in": bson.A{userStaticID},
		},
	}

	chats, err := repo.Where(filter)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (repo *repository) FindByID(staticID primitive.ObjectID) (*chat.Chat, error) {
	filter := bson.M{"chat_id": staticID}
	return repo.findBy(filter)
}

func (repo *repository) FindChatOrSidesByStaticID(staticID primitive.ObjectID) (*chat.Chat, error) {
	filter := bson.M{
		"$or": []interface{}{
			bson.M{"chat_detail.sides": bson.M{"$in": []primitive.ObjectID{staticID}}},
			bson.M{"chat_id": staticID},
		},
	}

	return repo.findBy(filter)
}

func (repo *repository) FindBySides(sides [2]primitive.ObjectID) (*chat.Chat, error) {
	filter := bson.M{
		"chat_detail.sides":     sides,
		"chat_detail.chat_type": bson.M{"$ne": "direct"},
	}

	return repo.findBy(filter)
}
