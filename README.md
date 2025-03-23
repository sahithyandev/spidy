# Spidy

A simple web crawler I am building to learn about web crawling and golang.

## Features

Here are the current features of Spidy:

- Queue based crawling: A disk-persisted queue is used to store the URLs to be crawled.
- Priorirty based crawling: URLs are prioritized based on their occurence.
- Respects robots.txt: If a host have a robots.txt, Spidy respects it to avoid crawling disallowed URLs.

## Development

### Prerequisites

- Go 1.23 or later

### Project setup

SQLite is used as the database. The database is created and seeded automatically if not present.

## Author

Sahithyan Kandathasan (https://sahithyan.dev)
