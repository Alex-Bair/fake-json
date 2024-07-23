# Purpose

Create fake JSON file for testing.

# Requirements
- Go

# Instructions

To generate file with one million documents:
```
go run . -count 1_000_000 | pv > data.jsonl
```

That file can then be split into multiple smaller ones with the below command:
```
split -dl 1000 --additional-suffix=.jsonl data.jsonl dat
```