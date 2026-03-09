# RSS Feed Aggregator

A command-line RSS feed aggregator written in Go with PostgreSQL for data storage.

This application lets users register, log in, add feeds, follow feeds, scrape posts from feeds, and browse recent posts from the feeds they follow.

## Features

- User registration and login
- Store data in PostgreSQL
- Add new RSS feeds
- Follow and unfollow feeds
- View available feeds
- Scrape posts from feeds on an interval
- Browse recent posts from followed feeds

## Requirements

Before running this project, make sure you have the following installed:

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

## Installation

### 1. Clone the repository

```bash
git clone <your-repo-url>
cd <your-project-folder>
````

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Install Goose (database migration tool)

This project uses **Goose** to manage PostgreSQL database migrations.

Install it with:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH`.

### 4. Set up PostgreSQL

Ensure PostgreSQL is installed and running.

Create a database for the project:

```sql
CREATE DATABASE rssagg;
```

### 5. Run database migrations

The database schema is located in:

```
sql/schema
```

Run the migrations using Goose:

```bash
goose -dir sql/schema postgres "postgres://<user>:<password>@localhost:5432/rssagg?sslmode=disable" up
```

This will apply all migrations and create the required tables.

### 6. Run the application

```bash
go run .
```

Or build it:

```bash
go build -o rssagg
./rssagg
```

## Usage

The program supports the following commands:

### Register a user

Creates a new user.

```bash
rssagg register <username>
```

### Log in

Logs in an existing user.

```bash
rssagg login <username>
```

### List users

Shows all registered users.

```bash
rssagg users
```

### Aggregate feeds

Scrapes posts from feeds repeatedly using the provided time interval.

**Usage:**

```bash
rssagg agg <time_between_scrapes>
```

**Example:**

```bash
rssagg agg 1m
rssagg agg 30s
```

### List feeds

Shows feeds available to follow.

```bash
rssagg feeds
```

### Add a feed

Creates a new feed and automatically follows it for the logged-in user.

**Usage:**

```bash
rssagg addfeed "<name>" "<url>"
```

**Example:**

```bash
rssagg addfeed "Hacker News" "https://news.ycombinator.com/rss"
```

### Follow a feed

Follows an already existing feed by URL.

**Usage:**

```bash
rssagg follow "<url>"
```

**Example:**

```bash
rssagg follow "https://news.ycombinator.com/rss"
```

### View followed feeds

Lists the feeds the current user is following.

```bash
rssagg following
```

### Unfollow a feed

Unfollows a feed for the logged-in user.

**Usage:**

```bash
rssagg unfollow "<url>"
```

**Example:**

```bash
rssagg unfollow "https://news.ycombinator.com/rss"
```

### Browse posts

Browse recent posts from feeds the logged-in user follows, ordered by most recent date.

**Usage:**

```bash
rssagg browse
rssagg browse <limit>
```

**Examples:**

```bash
rssagg browse
rssagg browse 5
```

Default limit is `2` if no value is provided.

## Notes

* Some commands require the user to be logged in first.
* PostgreSQL must be running before using the application.
* Make sure your database configuration is correct before starting the app.

## Project Structure

This project is built with:

* **Go** for the application logic
* **PostgreSQL** for persistent storage
