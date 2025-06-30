# go-follow-cli

A command-line tool for managing GitHub following relationships using the GitHub API.

## Features

- Follow specific GitHub users
- Follow all followers of a user
- Follow random GitHub users
- Simple and maintainable codebase

## Prerequisites

- Go 1.21 or higher
- GitHub personal access token with `user:follow` permission

## Installation

```bash
# Clone the repository
git clone https://github.com/opd-ai/go-follow-cli.git
cd go-follow-cli

# Build the application
go build -o go-follow-cli

# Or install directly
go install github.com/opd-ai/go-follow-cli@latest
```

## GitHub Token Setup

1. Go to GitHub Settings → Developer settings → Personal access tokens
2. Generate a new token with `user:follow` permission
3. Set the environment variable:

```bash
export GITHUB_TOKEN=your_token_here
```

## Usage

### Follow a specific user

```bash
go-follow-cli follow octocat
```

### Follow all followers of a user

```bash
go-follow-cli follow-all torvalds
```

### Follow a random user

```bash
go-follow-cli follow-random
```

### Follow N random users

```bash
go-follow-cli follow-n 5
```

## Commands

- `follow <username>` - Follow a specific GitHub user
- `follow-all <username>` - Follow all users who follow the specified username
- `follow-random` - Follow one randomly selected GitHub user
- `follow-n <count>` - Follow N randomly selected GitHub users (max 100)

## Error Handling

The tool provides clear error messages for common issues:

- Missing GitHub token
- Invalid usernames
- Rate limit warnings
- Network errors
- Already following user

## Development

### Running Tests

```bash
go test ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter
golangci-lint run
```

## Rate Limits

The tool automatically handles GitHub API rate limits by:
- Monitoring remaining requests
- Sleeping when limits are low
- Providing feedback to users

## License

MIT License

## Contributing

Contributions are welcome! Please ensure:
- Code follows Go conventions
- All tests pass
- Clear commit messages
- Updated documentation