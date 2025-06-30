package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

// GitHubClient wraps the GitHub API client with common operations
type GitHubClient struct {
	client *github.Client
	ctx    context.Context
}

// NewGitHubClient creates a new authenticated GitHub API client
func NewGitHubClient(token string) *GitHubClient {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(ctx, tokenSource)

	return &GitHubClient{
		client: github.NewClient(tokenClient),
		ctx:    ctx,
	}
}

// FollowUser follows a specific GitHub user
func (gc *GitHubClient) FollowUser(username string) error {
	// First check if user exists
	user, _, err := gc.client.Users.Get(gc.ctx, username)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == 404 {
			return fmt.Errorf("user '%s' not found", username)
		}
		return fmt.Errorf("error checking user: %w", err)
	}

	// Check if already following
	isFollowing, _, err := gc.client.Users.IsFollowing(gc.ctx, "", username)
	if err != nil {
		return fmt.Errorf("error checking follow status: %w", err)
	}

	if isFollowing {
		return fmt.Errorf("already following user '%s'", username)
	}

	// Follow the user
	_, err = gc.client.Users.Follow(gc.ctx, username)
	if err != nil {
		return fmt.Errorf("error following user: %w", err)
	}

	fmt.Printf("Successfully followed %s (%s)\n", username, user.GetName())
	return nil
}

// GetFollowers returns a list of users who follow the specified username
func (gc *GitHubClient) GetFollowers(username string, limit int) ([]*github.User, error) {
	var allFollowers []*github.User

	opts := &github.ListOptions{
		PerPage: 100,
	}

	for {
		followers, resp, err := gc.client.Users.ListFollowers(gc.ctx, username, opts)
		if err != nil {
			if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == 404 {
				return nil, fmt.Errorf("user '%s' not found", username)
			}
			return nil, fmt.Errorf("error fetching followers: %w", err)
		}

		allFollowers = append(allFollowers, followers...)

		// Check if we've reached the limit
		if limit > 0 && len(allFollowers) >= limit {
			return allFollowers[:limit], nil
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage

		// Rate limit handling
		if resp.Rate.Remaining < 10 {
			sleepDuration := time.Until(resp.Rate.Reset.Time) + time.Second
			fmt.Printf("Rate limit low, sleeping for %v...\n", sleepDuration)
			time.Sleep(sleepDuration)
		}
	}

	return allFollowers, nil
}

// GetRandomUsers fetches random GitHub users
func (gc *GitHubClient) GetRandomUsers(count int) ([]*github.User, error) {
	users := make([]*github.User, 0, count)

	// Use the search API to find random users
	// We'll search for users created after a random date
	randomGenerator := NewRandomUserGenerator()

	for len(users) < count {
		query := randomGenerator.GenerateSearchQuery()

		searchOpts := &github.SearchOptions{
			ListOptions: github.ListOptions{
				PerPage: 30,
			},
		}

		result, resp, err := gc.client.Search.Users(gc.ctx, query, searchOpts)
		if err != nil {
			return nil, fmt.Errorf("error searching users: %w", err)
		}

		// Add non-duplicate users
		for _, user := range result.Users {
			if len(users) >= count {
				break
			}

			// Check if already following
			isFollowing, _, _ := gc.client.Users.IsFollowing(gc.ctx, "", user.GetLogin())
			if !isFollowing {
				users = append(users, user)
			}
		}

		// Rate limit handling
		if resp.Rate.Remaining < 10 {
			sleepDuration := time.Until(resp.Rate.Reset.Time) + time.Second
			fmt.Printf("Rate limit low, sleeping for %v...\n", sleepDuration)
			time.Sleep(sleepDuration)
		}
	}

	return users, nil
}
