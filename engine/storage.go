package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/domain"
)

type (
	ForumRepository interface {
		Get(id domain.ForumID) *domain.Forum
		Put(forum *domain.Forum)
		List(query *Query) []*domain.Forum
	}

	GreetingRepository interface {
		Put(greeting *domain.Greeting)
		List(query *Query) []*domain.Greeting
	}

	TopicRepository interface {
		Get(id domain.TopicID) *domain.Topic
		Put(topic *domain.Topic)
		List(query *Query) []*domain.Topic
	}

	UserRepository interface {
		Get(id domain.UserID) *domain.User
		Put(user *domain.User)
		List(query *Query) []*domain.User
	}

	StorageFactory interface {
		GetGreetingRepository(c context.Context) GreetingRepository
	}
)
