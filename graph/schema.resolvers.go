package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"LFGbackend/graph/model"
	"LFGbackend/resolvers"
	"context"
)

// SignInAsPlayer is the resolver for the signInAsPlayer field.
func (r *mutationResolver) SignInAsPlayer(ctx context.Context, player model.PlayerInput) (*model.Player, error) {
	return resolvers.UpsertPlayerResolver(r.Server, r.Db, ctx, player)
}

// SendMessage is the resolver for the sendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, text string) (bool, error) {
	return resolvers.SendMessageResolver(text, ctx, r.Server)
}

// GetPlayers is the resolver for the getPlayers field.
func (r *queryResolver) GetPlayers(ctx context.Context, players []*model.PlayerInput) ([]*model.Player, error) {
	return resolvers.GetPlayerResolver(r.Db, ctx, players)
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, page int, count int) ([]*model.PostInfo, error) {
	return resolvers.GetPostsResolver(r.Server, ctx, page, count)
}

// JoinPost is the resolver for the joinPost field.
func (r *subscriptionResolver) JoinPost(ctx context.Context, player model.PlayerInput, id string) (<-chan *model.Post, error) {
	return resolvers.JoinPostResolver(ctx, r.Server, player, id)
}

// CreatePost is the resolver for the createPost field.
func (r *subscriptionResolver) CreatePost(ctx context.Context, mode model.GameMode, player model.PlayerInput, need int, minRank model.Rank) (<-chan *model.Post, error) {
	return resolvers.CreatePostResolver(ctx, r.Server, r.Db, mode, player, need, minRank)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
