# SSH Exporter

A Prometheus exporter that exposes SSH session and login metrics for monitoring SSH activity on your Linux servers.

## Features

This exporter collects and exposes the following SSH-related metrics:

- `ssh_active_sessions`: Number of active SSH sessions.
- `ssh_active_sessions_by_user{user}`: Number of active SSH sessions by each user.
- `ssh_logins_today`: Number of successful SSH logins today.
- `ssh_user_last_login{user}`: Last login time per user (Unix timestamp).
- `ssh_failed_logins_today_total`: Number of failed SSH login attempts today.

## Getting Started

### Installation

Download the precompiled binary for your platform from the [Releases](https://github.com/Himanshu-216/ssh_exporter/releases) page, or build from source:

```bash
GOOS=<os_name> GOARCH=<arch> go build -o ssh-exporter main.go
