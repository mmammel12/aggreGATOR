# aggreGATOR

aggregator project from boot.dev

## Requirements

- Postgres
- Go

## Installation

1. clone the repo
2. cd to the root directory of the repo
3. run `go install`
4. set up a config file
    a. create a file named `.gatorconfig.json` in your home directory
    b. minimum config is `{ "db_url": "<your db url>" }`
        i. example for local usage: `{ "db_url": "postgres://<local username>:@localhost:5432/gator?sslmode=disable" }`
5. `aggreGATOR` command should now be available

## Usage

### Commands

- reset - reset the database
    - example: `aggreGATOR reset`
- register <user> - register a user (also logs you in as this name)
    - example: `aggreGATOR register marty`
- login <user> - change active user (must be a registered user)
    - example: `aggreGATOR login marty`
- addfeed <name> <url> - add an rss feed
    - example: `aggreGATOR addfeed "Hacker News RSS" "https://hnrss.org/newest"`
- feeds - list feeds
    - example: `aggreGATOR feeds`
- follow <url> - follow and existing feed
    - example: `aggreGATOR follow "https://hnrss.org/newest"`
- following - list feeds you follow
    - example: `aggreGATOR following`
- unfollow <url> - unfollow a feed
    - example: `aggreGATOR unfollow "https://hnrss.org/newest"`
- browse <optional limit> - browse feed info
    - example: `aggreGATOR browse` or `aggreGATOR browse 5`
