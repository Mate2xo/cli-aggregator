# Blog aggreGator

Basic RSS feed collector; toy project to discover Go.

## Installation

Requires:

- PostgresSQL
- Go runtime

Then `$ go install gator` where you have cloned this repo,

## Usage

First create a config dotfile in your `home` directory : `~/.gatorconfig.json`.
You only need to set your DB URL to get started, so this file should look like:

```json
{
  "db_url": "postgres://username:@localhost:5432/your_db?sslmode=disable"
}
```

Then you can launch the program as `gator <command> <arguments>`

## Available commands

- `agg <duration>`: running service that loops through the current user's followed feeds, and fetches their last version
  - e.g.: `$ gator agg 2m` -> the service will loop on followed feeds to fetch every two minutes

- `register <username>`: creates a User in the DB, and sets it as the current user
  - e.g.: `$ gator register pikachu`
- `login <username>`: sets the given user as the current user, if it exists
  - e.g.: `$ gator login pikachu`
- `users`: lists currently registered users

- `addfeed <feed name> <feed URL>`: registers a given feed, which can then be fetched by the `agg` service
  - e.g.: `$ gator addfeed TechCrunch https://techcrunch.com/feed/`
- `feeds`: list currently registered feeds
- `follow <url>`: the current user is associated with the given feed, if it exists, thus "following" it
  - e.g.: `$ gator follow https://techcrunch.com/feed/`
- `following`: the currently followed feeds by the current user
- `unfollow <url>`: remove a feed from the current user's following list
  - e.g.: `$ gator unfollow https://techcrunch.com/feed/`

- `browse`: show the latest article details from the current user's feeds
