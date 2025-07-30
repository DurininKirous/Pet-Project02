# 📁 bigscanner

**bigscanner** is a lightweight CLI utility written in Go to find the largest file in a given directory (recursively).

## Features

- Scans directories to find the largest file
- Optional `--max-depth` to limit traversal depth
- Supports `--json`, `--human` output and `--abs-path`
- Easy to use with Docker
- Ready for CI/CD with GitHub Actions and Docker Hub

## Usage

```bash
bigscanner --dir ./some/path [--max-depth N] [--json] [--abs-path] [--human]
