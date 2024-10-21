# GitHub Activity CLI

A command-line interface tool that displays recent GitHub activity for any user.

Sample solution for the [github-user-activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh/).

## Features

- Fetch and display recent GitHub user activity
- Clean terminal output format
- Error handling for API failures and invalid usernames


build from source:

```bash
git clone https://github.com/mazyaryousefinia/github-activity
cd github-activity
go build
```

## Usage

```bash
# Basic usage
github-activity <username>

# Example
github-activity mazyaryousefinia
```

### Sample Output
```
- Pushed 1 commit(s) to mazyaryousefinia/cli-task-tracker
- Starred dabit3/prompt-engineering-for-javascript-developers
```

