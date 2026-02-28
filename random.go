package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomUserGenerator helps generate search queries for finding random users
type RandomUserGenerator struct {
	random *rand.Rand
}

// NewRandomUserGenerator creates a new random user generator
func NewRandomUserGenerator() *RandomUserGenerator {
	return &RandomUserGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateSearchQuery generates a random search query for finding users
func (r *RandomUserGenerator) GenerateSearchQuery() string {
	// Different strategies for finding random users
	strategies := []func() string{
		r.searchByFollowers,
		r.searchByRepositories,
		r.searchByCreationDate,
		r.searchByLocation,
		r.searchByLanguage,
	}

	// Pick a random strategy
	strategy := strategies[r.random.Intn(len(strategies))]
	return strategy()
}

// searchByFollowers searches for users with a random follower count range
func (r *RandomUserGenerator) searchByFollowers() string {
	min := r.random.Intn(50)
	max := min + r.random.Intn(100) + 10
	return fmt.Sprintf("followers:%d..%d", min, max)
}

// searchByRepositories searches for users with a random repository count
func (r *RandomUserGenerator) searchByRepositories() string {
	min := r.random.Intn(20)
	max := min + r.random.Intn(50) + 5
	return fmt.Sprintf("repos:%d..%d", min, max)
}

// searchByCreationDate searches for users created on random dates
func (r *RandomUserGenerator) searchByCreationDate() string {
	// Random date within the last 5 years
	daysAgo := r.random.Intn(365 * 5)
	date := time.Now().AddDate(0, 0, -daysAgo).Format("2006-01-02")
	return fmt.Sprintf("created:%s", date)
}

// searchByLocation searches for users in random locations
func (r *RandomUserGenerator) searchByLocation() string {
	locations := []string{
		"USA", "UK", "Canada", "Germany", "France", "Japan", "Australia",
		"Brazil", "India", "China", "Russia", "Italy", "Spain", "Netherlands",
		"Sweden", "Poland", "Ukraine", "Argentina", "Mexico", "Indonesia",
	}
	location := locations[r.random.Intn(len(locations))]
	return fmt.Sprintf("location:%s", location)
}

// searchByLanguage searches for users who use specific programming languages
func (r *RandomUserGenerator) searchByLanguage() string {
	languages := []string{
		"JavaScript", "Python", "Java", "Go", "TypeScript", "C++", "Ruby",
		"PHP", "C#", "Swift", "Kotlin", "Rust", "Scala", "Dart", "R",
	}
	language := languages[r.random.Intn(len(languages))]
	minRepos := r.random.Intn(5) + 1
	return fmt.Sprintf("language:%s repos:>%d", language, minRepos)
}

// RandomRepoGenerator helps generate search queries for finding random repositories
type RandomRepoGenerator struct {
	random *rand.Rand
}

// NewRandomRepoGenerator creates a new random repository generator
func NewRandomRepoGenerator() *RandomRepoGenerator {
	return &RandomRepoGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateSearchQuery generates a random search query for finding repositories
func (r *RandomRepoGenerator) GenerateSearchQuery() string {
	strategies := []func() string{
		r.searchByStars,
		r.searchByForks,
		r.searchByRepoCreationDate,
		r.searchByRepoLanguage,
		r.searchByTopic,
	}

	strategy := strategies[r.random.Intn(len(strategies))]
	return strategy()
}

// searchByStars searches for repositories with a random star count range
func (r *RandomRepoGenerator) searchByStars() string {
	min := r.random.Intn(100)
	max := min + r.random.Intn(500) + 10
	return fmt.Sprintf("stars:%d..%d", min, max)
}

// searchByForks searches for repositories with a random fork count range
func (r *RandomRepoGenerator) searchByForks() string {
	min := r.random.Intn(20)
	max := min + r.random.Intn(100) + 5
	return fmt.Sprintf("forks:%d..%d", min, max)
}

// searchByRepoCreationDate searches for repositories created on random dates
func (r *RandomRepoGenerator) searchByRepoCreationDate() string {
	daysAgo := r.random.Intn(365 * 5)
	date := time.Now().AddDate(0, 0, -daysAgo).Format("2006-01-02")
	return fmt.Sprintf("created:%s", date)
}

// searchByRepoLanguage searches for repositories using specific programming languages
func (r *RandomRepoGenerator) searchByRepoLanguage() string {
	languages := []string{
		"JavaScript", "Python", "Java", "Go", "TypeScript", "C++", "Ruby",
		"PHP", "C#", "Swift", "Kotlin", "Rust", "Scala", "Dart", "R",
		"Haskell", "Elixir", "Clojure", "Lua", "Perl",
	}
	language := languages[r.random.Intn(len(languages))]
	minStars := r.random.Intn(10) + 1
	return fmt.Sprintf("language:%s stars:>%d", language, minStars)
}

// searchByTopic searches for repositories with specific topics
func (r *RandomRepoGenerator) searchByTopic() string {
	topics := []string{
		"cli", "api", "web", "database", "machine-learning", "devops",
		"testing", "security", "blockchain", "game", "mobile", "cloud",
		"microservices", "data-science", "automation", "networking",
	}
	topic := topics[r.random.Intn(len(topics))]
	minStars := r.random.Intn(20) + 1
	return fmt.Sprintf("topic:%s stars:>%d", topic, minStars)
}
