# Spin

Quickly spin up a new project.

## Introduction

A Go implementation of [SAO](https://github.com/saojs/sao).

_WIP_.

## Install

```bash
curl -sf https://gobinaries.com/egoist/spin | sh
```

Or [download a binary manually](https://github.com/egoist/spin/releases).

## Usage

```bash
spin user/repo new-project
```

### Private Repository

Use `--clone` which uses `git clone` under the hood to download your repository:

```bash
spin user/private-repo new-project --clone
```

### Update Cached Repository

Spin will cache the repository after first run, you can use `-u, --update` to update it:

```bash
spin owner/repo new-project --update
```

## License

MIT &copy; [EGOIST (Kevin Titor)](https://egoist.sh)
