# LiteSSH

> Lightweight multi-server management with a powerful VSCode editing experience makes remote work more efficient and seamless.

LiteSSH has the following advantages:

- Powerful file editing experience: Based on VSCode, it supports syntax highlighting, intelligent prompts, and shortcuts.
- Multi-server support: Copy and move files between multiple servers.
- Open-source: Prevents malicious key leaks and allows for further customization.
- Web-based: Runs in a browser without needing to install clients on multiple devices.
- Configuration synchronization: Saves server configurations and keys at the backend, eliminating the need for repeated setups across devices.
- Quick response: Maintains long-term SSH connections with servers from the backend instead of establishing connections after opening the client.
- Plugin ecosystem: Compatible with VSCode plugins.

VSCode Remote is very convenient for managing servers, featuring a layout with a file bar on the left side and file content and terminal on the right. Thanks to the integration of the file bar, terminal linkage, and VSCode's robust text editing capabilities, VSCode Remote offers a far superior user experience compared to XSHELL, XFTP, WinSCP, FinalShell. However, VSCode Remote has the following drawbacks:

- Lack of multi-server support: Can only open one server connection at a time and cannot copy or move files between servers.
- Not open-source: Although VSCode itself is open-source, VSCode Remote is not.
- Installation of Agent: Requires installing an agent on the server, leading to various issues such as download failures during installation/updates, compatibility issues with certain Linux distributions, and potential violations of server security policies.
- Client installation: Needs to be installed and continuously updated on each device.
- Difficult configuration synchronization: Server connection information and keys need to be repeatedly configured on each device.
- Slow response: Establishes new windows and connections when connecting to servers, requiring several seconds for each server. Reconnection is required after network interruptions.

LiteSSH improves upon VSCode by avoiding dependency on Agents, thereby achieving higher efficiency in operation and maintenance scenarios. However, the lack of an Agent means LiteSSH plugins cannot interact with the server environment, so complex remote development still requires VSCode Remote (or can alternatively use code-server for remote development).

## Quick Start

Prepare the configuration file and Docker Compose declaration file:

```sh
git clone https://github.com/117503445/LiteSSH.git
cd LiteSSH/docs/example
```

For users in mainland China, pull the Docker image from Alibaba Cloud ACR:

```sh
docker pull registry.cn-hangzhou.aliyuncs.com/117503445/litessh && docker tag registry.cn-hangzhou.aliyuncs.com/117503445/litessh 117503445/litessh
```

Start the service:

```sh
docker compose up -d
```

Open in your browser <http://localhost:4444/?folder=/remote>, enter the password 123456, and start managing server1, server2, and server3.

Click "Open in Remote Terminal" on Files/Folders to open the remote terminal.

You can also directly input `r server3` in the terminal to remotely connect to server3.

## Configuration Reference

Use TOML format configuration files. Here is an example based on the quick start configuration:

```toml
code-server-password = "123456"

[nodes.server1]
host = "server1"
pri = "/root/.ssh/id_ed25519"
# default path is ~
# default user is root
# default port is 22

[nodes.server2]
host = "server2"
pri = "/root/.ssh/id_ed25519"
path = ".ssh" # use relative path of ~

[nodes.server3]
host = "server3"
pri = "/root/.ssh/id_ed25519"
path = "/etc" # use absolute path
```

The `code-server-password` is the password for code-server. It defaults to empty. When `code-server-password` is empty, you can log in to code-server without a password.

`nodes` is an array where each element represents a server. `nodes.server1` defines the configuration for server1.

Each server supports the following configurations:

| Config | Default Value | Required | Description |
| --- | --- | --- | --- |
| host | - | Yes | The hostname or IP address of the server |
| port | 22 | No | The SSH port of the server |
| user | root | No | The username of the server |
| pri | - | Yes | Path to the private key |
| path | ~ | No | The mounted directory on the server |

Only private key login is supported, not password login.

## Implementation

Understanding how LiteSSH is implemented can help you troubleshoot related issues.

The LiteSSH image is based on ArchLinux and comes pre-installed with code-server. Additionally, it compiles and installs the LiteSSH plugin to add the feature of opening a remote terminal via right-click.

Upon container startup, the entrypoint performs two tasks:
- Launching code-server and setting the password according to the configuration file.
- Starting the LiteSSH service and mounting each server's specified directories under `/remote` using Rclone based on the configuration file.

After opening the remote terminal, the LiteSSH plugin opens a new terminal and inputs `r $path`, where `$path` is the directory path clicked by the user. Upon receiving the `$path` parameter, the `r` program generates and executes an SSH command based on the configuration file to open the remote terminal. 

The `/workspace/logs` directory inside the container contains logs for code-server, LiteSSH, and each Rclone service.