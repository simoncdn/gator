# Gator

A command-line RSS/Atom feed aggregator and reader built with Go and PostgreSQL.

## Features

- **User Management**: Register, login, and manage user accounts
- **Feed Management**: Add RSS/Atom feeds and follow/unfollow them
- **Content Aggregation**: Automatically fetch and store posts from followed feeds
- **Content Browsing**: Browse and read posts from your followed feeds
- **Database Integration**: PostgreSQL with automated migrations using Goose

## Prerequisites

- Go 1.24.2 or later
- PostgreSQL database
- Goose migration tool

## Installation

1. Clone the repository:
```bash
git clone https://github.com/simoncdn/gator.git
cd gator
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your PostgreSQL database and configure the connection string in your config file.

4. Run database migrations:
```bash
goose up
```

5. Build the application:
```bash
go build -o gator
```

## Configuration

The application requires a configuration file with your database connection details. The config is managed through the `internal/config` package.

## Usage

### User Management
```bash
# Register a new user
./gator register <username>

# Login as a user
./gator login <username>

# List all users
./gator users

# Reset user data
./gator reset
```

### Feed Management
```bash
# Add a new RSS/Atom feed
./gator addfeed <feed_url>

# List all available feeds
./gator feeds

# Follow a feed
./gator follow <feed_name>

# List feeds you're following
./gator following

# Unfollow a feed
./gator unfollow <feed_name>
```

### Content Management
```bash
# Aggregate posts from all followed feeds
./gator agg

# Browse posts from followed feeds
./gator browse [limit]
```

## Example Workflow

1. Register and login:
```bash
./gator register john
./gator login john
```

2. Add and follow feeds:
```bash
./gator addfeed https://news.ycombinator.com/rss
./gator follow HackerNews
```

3. Fetch and browse content:
```bash
./gator agg
./gator browse 5
```

## Database Schema

The application uses PostgreSQL with the following main tables:
- `users`: User accounts
- `feeds`: RSS/Atom feed information
- `feed_follows`: User-feed relationships
- `posts`: Aggregated posts from feeds

## Dependencies

- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [github.com/google/uuid](https://github.com/google/uuid) - UUID generation

