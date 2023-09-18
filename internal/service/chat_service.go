package service

import (
	"Kavka/internal/domain/chat"
	"Kavka/internal/domain/user"
	chatRepository "Kavka/internal/repository/chat"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chatService struct {
	chatRepo *chatRepository.ChatRepository
	userRepo *user.UserRepository
}

func NewChatService(chatRepo *chatRepository.ChatRepository, userRepo *user.UserRepository) *chatService {
	return &chatService{chatRepo, userRepo}
}

func (s *chatService) GetChat(staticID primitive.ObjectID) (*chat.Chat, error) {
	foundChat, foundChatErr := s.chatRepo.FindChatOrSidesByStaticID(&staticID)
	if foundChatErr != nil {
		return nil, foundChatErr
	}

	return foundChat, nil
}

func (s *chatService) CreateDirect(userStaticID primitive.ObjectID,
	targetStaticID primitive.ObjectID,
) (*chat.Chat, error) {
	sides := [2]*primitive.ObjectID{
		&userStaticID,
		&targetStaticID,
	}

	dup, _ := s.chatRepo.FindBySides(sides)
	if dup != nil {
		return nil, chatRepository.ErrChatAlreadyExists
	}

	return s.chatRepo.Create(chat.ChatTypeDirect, &chat.DirectChatDetail{
		Sides: sides,
	})
}

func (s *chatService) CreateGroup(userStaticID primitive.ObjectID,
	title string, username string, description string,
) (*chat.Chat, error) {
	return s.chatRepo.Create(chat.ChatTypeGroup, &chat.GroupChatDetail{
		Title:       title,
		Username:    username,
		Members:     []*primitive.ObjectID{&userStaticID},
		Admins:      []*primitive.ObjectID{&userStaticID},
		Description: description,
	})
}

func (s *chatService) CreateChannel(userStaticID primitive.ObjectID,
	title string, username string, description string,
) (*chat.Chat, error) {
	return s.chatRepo.Create(chat.ChatTypeGroup, &chat.GroupChatDetail{
		Title:       title,
		Username:    username,
		Members:     []*primitive.ObjectID{&userStaticID},
		Admins:      []*primitive.ObjectID{&userStaticID},
		Description: description,
	})
}
