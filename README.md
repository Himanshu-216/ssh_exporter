# SSH Exporter for Prometheus

This SSH Exporter provides Prometheus-compatible metrics to track SSH activity on your server. It includes metrics such as the number of active SSH sessions, active sessions per user, login attempts, and more. This exporter is designed to help monitor SSH usage, failed login attempts, and user activity, which can be used to create alerts and perform analysis on your server's SSH activity.

## How to Run Binary

```bash
make build
```

```bash
./bin/ssh_exporter --web.listen-address="0.0.0.0:9898"
```

## How to Run with Docker

```bash
docker pull himanshu162pnt723/ssh_exporter
```
```bash
docker run --privileged --rm \
  -p 9898:9898 \
  -v /var/log/wtmp:/var/log/wtmp:ro \
  -v /var/log/btmp:/var/log/btmp:ro \
  -v /var/run/utmp:/var/run/utmp:ro \
  -v /var/log/lastlog:/var/log/lastlog:ro \
  -v /etc/passwd:/etc/passwd:ro \
  himanshu162pnt723/ssh_exporter
```
## Features

- **Active SSH Sessions**: Tracks the number of active SSH sessions on the server.
- **Sessions by User**: Tracks the number of active SSH sessions per user.
- **SSH Logins Today**: Tracks the number of SSH logins that have occurred today.
- **Last Login Time**: Provides the last login time for each user (in Unix timestamp).
- **Failed Logins by IP**: Tracks failed SSH login attempts by source IP address.
- **Successful Logins by IP**: Tracks successful SSH login attempts by source IP address.

You can use these Prometheus queries to visualize and alert on SSH activity.

- **Total active SSH sessions**:

    ```promQL
    ssh_active_sessions
    ```

- **Active sessions by user**:

    ```promQL
    ssh_active_sessions_by_user
    ```

- **Logins today**:

    ```promQL
    ssh_logins_today
    ```

- **Last login time for a user**:

    ```promQL
    ssh_user_last_login{user="username"}
    ```

- **Failed logins today**:

    ```promQL
    SSHFailedLoginsByIP
    ```

- **Successful logins today**:

    ```promQL
    SSHSuccessfulLoginsByIP
    ```

## Contributing

Feel free to open an issue or submit a pull request if you find any bugs or have suggestions for new features.

## Releases

You can download the latest release of the SSH Exporter from the [Releases](https://github.com/Himanshu-216/ssh_exporter/releases) page on GitHub.

- [Latest Release](https://github.com/Himanshu-216/ssh_exporter/releases/latest)






