package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// newFollowCommand creates the 'follow' command
func newFollowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "follow <username>",
		Short: "Follow a specific GitHub user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			username := args[0]

			client := NewGitHubClient(os.Getenv("GITHUB_TOKEN"))

			fmt.Printf("Following user: %s\n", username)
			if err := client.FollowUser(username); err != nil {
				return fmt.Errorf("failed to follow user: %w", err)
			}

			return nil
		},
	}
}

// newFollowAllCommand creates the 'follow-all' command
func newFollowAllCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "follow-all <username>",
		Short: "Follow all users who follow the specified username",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			username := args[0]

			client := NewGitHubClient(os.Getenv("GITHUB_TOKEN"))

			fmt.Printf("Fetching followers of %s...\n", username)
			followers, err := client.GetFollowers(username, 0)
			if err != nil {
				return fmt.Errorf("failed to get followers: %w", err)
			}

			if len(followers) == 0 {
				fmt.Printf("User %s has no followers\n", username)
				return nil
			}

			fmt.Printf("Found %d followers. Following them...\n", len(followers))

			successCount := 0
			for i, follower := range followers {
				fmt.Printf("[%d/%d] Following %s...\n", i+1, len(followers), follower.GetLogin())

				if err := client.FollowUser(follower.GetLogin()); err != nil {
					fmt.Printf("  Warning: %v\n", err)
				} else {
					successCount++
				}
			}

			fmt.Printf("\nSuccessfully followed %d out of %d users\n", successCount, len(followers))
			return nil
		},
	}
}

// newFollowRandomCommand creates the 'follow-random' command
func newFollowRandomCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "follow-random",
		Short: "Follow one randomly selected GitHub user",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := NewGitHubClient(os.Getenv("GITHUB_TOKEN"))

			fmt.Println("Finding a random user to follow...")
			users, err := client.GetRandomUsers(1)
			if err != nil {
				return fmt.Errorf("failed to get random user: %w", err)
			}

			if len(users) == 0 {
				return fmt.Errorf("no random users found")
			}

			user := users[0]
			fmt.Printf("Selected random user: %s\n", user.GetLogin())

			if err := client.FollowUser(user.GetLogin()); err != nil {
				return fmt.Errorf("failed to follow user: %w", err)
			}

			return nil
		},
	}
}

// newFollowNCommand creates the 'follow-n' command
func newFollowNCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "follow-n <count>",
		Short: "Follow N randomly selected GitHub users",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			count, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid count: %s", args[0])
			}

			if count <= 0 {
				return fmt.Errorf("count must be greater than 0")
			}

			if count > 100 {
				return fmt.Errorf("count must be 100 or less to avoid rate limits")
			}

			client := NewGitHubClient(os.Getenv("GITHUB_TOKEN"))

			fmt.Printf("Finding %d random users to follow...\n", count)
			users, err := client.GetRandomUsers(count)
			if err != nil {
				return fmt.Errorf("failed to get random users: %w", err)
			}

			if len(users) == 0 {
				return fmt.Errorf("no random users found")
			}

			fmt.Printf("Found %d users. Following them...\n", len(users))

			successCount := 0
			for i, user := range users {
				fmt.Printf("[%d/%d] Following %s...\n", i+1, len(users), user.GetLogin())

				if err := client.FollowUser(user.GetLogin()); err != nil {
					fmt.Printf("  Warning: %v\n", err)
				} else {
					successCount++
				}
			}

			fmt.Printf("\nSuccessfully followed %d out of %d users\n", successCount, len(users))
			return nil
		},
	}
}
