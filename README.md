# MapReduce

A simple MapReduce implementation in Go that finds the top 10 most active users from a large dataset.

## How It Works

1. **Map** — Reads `input.txt` line by line, emits `(username, 1)` pairs into per-user intermediate files under `intermediate/`.
2. **Reduce** — Reads each intermediate file, counts the lines, and aggregates totals per user.
3. **Output** — Sorts by count and prints the top 10 users.

> [!NOTE]
> The generator simulates 10 hardcoded usernames for simplicity. In a real discord server you'd have thousands of unique users — that scale can't be meaningfully simulated locally, but the MapReduce pipeline handles it identically regardless of unique user count. The intermediate directory just grows to one file per unique user.

## Setup

**Prerequisites:** Go 1.26+, Python 3

```bash
# Clone
git clone git@github.com:sYanXO/mapReduce.git
cd mapReduce

# Generate test data
python3 generate.py          # defaults to 1 million lines
python3 generate.py 500000   # or specify a custom count

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
