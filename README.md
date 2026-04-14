# MapReduce

A simple MapReduce implementation in Go that finds the top 10 most active users from a large dataset.

## How It Works

1. **Map** — Reads `input.txt` line by line, emits `(username, 1)` pairs into per-user intermediate files under `intermediate/`.
2. **Reduce** — Reads each intermediate file, counts the lines, and aggregates totals per user.
3. **Output** — Sorts by count and prints the top 10 users.

## Setup

**Prerequisites:** Go 1.26+, Python 3

```bash
# Clone
git clone git@github.com:sYanXO/mapReduce.git
cd mapReduce

# Generate test data (default: 1M lines)
python3 generate.py          # or: python3 generate.py 5000000

# Run
go run main.go
```

## Project Structure

```
├── main.go          # MapReduce implementation
├── generate.py      # Test data generator
├── go.mod
└── .gitignore
```
