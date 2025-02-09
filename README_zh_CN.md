# LiteSSH

> Lightweight multi-server management with a powerful VSCode editing experience, making remote work more efficient and seamless.

LiteSSH offers the following advantages:

- **Powerful file editing experience**: Based on VSCode, supporting syntax highlighting, intelligent suggestions, and shortcuts.
- **Multi-server support**: Copy and move files between multiple servers.
- **Open source**: Prevents malicious key leaks and allows for further customization.
- **Web-based**: Runs in the browser, eliminating the need to install clients on multiple devices.
- **Configuration synchronization**: Saves server configurations and keys on the backend, avoiding repetitive configurations on multiple devices.
- **Quick response**: Maintains long-term SSH connections with servers on the backend, rather than establishing connections only when the client is opened.
- **Plugin ecosystem**: Compatible with VSCode plugins.

VSCode Remote is very convenient for managing servers. Its layout includes a file bar on the left and file content and terminal on the right. Thanks to the integration of the file bar, terminal, and VSCode's powerful text editing capabilities, VSCode Remote offers a much better user experience compared to XSHELL, XFTP, WinSCP, and FinalShell. However, VSCode Remote has the following drawbacks:

- **Lack of multi-server support**: You can only connect to servers one by one, and you cannot copy or move files between servers.
- **Not open source**: Although VSCode itself is open source, VSCode Remote is not.
- **Agent installation**: Requires installing an Agent on the server, leading to various issues:
  - Agent installation/update download failures.
  - Agent is based on Node and cannot be installed on certain specific distributions.
  - Agent cannot run on very old distributions.
  - Agent installation may violate server security management rules.
- **Client installation**: Requires installing and continuously updating VSCode on each device.
- **Difficult configuration synchronization**: Requires repeatedly configuring server connection information and keys on each device.
- **Slow response**: When connecting to a server, a new window is created and a connection is established, with each server taking several seconds to connect. After a network interruption, the process must be painfully repeated.

LiteSSH, based on VSCode, avoids the dependency on Agents, thus offering higher efficiency in operational scenarios. However, the absence of an Agent means LiteSSH plugins cannot interact with the server environment, so complex remote development still requires VSCode Remote. (Of course, remote development can also be achieved through code-server.)

## Quick Start

Prepare the configuration file and Docker Compose declaration file:

```sh
git clone https://github.com/117503445/LiteSSH.git
cd LiteSSH/docs/example
```

For users in mainland China, you can pull the Docker image from Alibaba Cloud ACR:

```sh
docker pull registry.cn-hangzhou.aliyuncs.com/117503445/litessh && docker tag registry.cn-hangzhou.aliyuncs.com/117503445/litessh 117503445/litessh
```

Start the service:

```sh
docker compose up -d
```

Open <http://localhost:4444/?folder=/remote> in your browser and enter the password `123456`.

You can now start managing server1, server2, and server3.

![index](./docs/assets/index.png)

Click "Open in Remote Terminal" on a file/folder to open a remote terminal.

![terminal](./docs/assets/terminal.png)

You can also directly enter `r server3` in the terminal to remotely connect to server3.

## Configuration Reference

Use the TOML format configuration file. Take the configuration in the quick start as an example:

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

`code-server-password` is the password for code-server. The default is empty. When `code-server-password` is empty, code-server can be logged into without a password.

`nodes` is an array, where each element represents a server. `nodes.server1` defines the configuration for server1.

Each server supports the following configurations:

| Configuration Item | Default Value | Required | Description |
| --- | --- | --- | --- |
| host | - | Yes | The hostname or IP address of the server |
| port | 22 | No | The SSH port of the server |
| user | root | No | The username of the server |
| pri | - | Yes | The path to the private key |
| path | ~ | No | The mounted directory on the server. When `path` is empty, it defaults to `~`; when `path` starts with `/`, it represents an absolute path; otherwise, it is a relative path under `~`, such as `.ssh` representing `~/.ssh`. |

Only private key login is supported; password login is not supported.

## Implementation

Understanding how LiteSSH is implemented can help you locate and resolve related issues.

The LiteSSH image is based on ArchLinux and comes pre-installed with code-server. Additionally, the LiteSSH plugin is compiled, packaged, and installed to add the right-click "Open in Remote Terminal" feature.

When the container starts, the entrypoint is launched first. The entrypoint does the following two things:
- Starts code-server and sets the password according to the configuration file.
- Starts the litessh service and, based on the configuration file, uses Rclone to mount the specified directory of each server under `/remote`.

After the user opens a remote terminal, the LiteSSH plugin opens a new terminal and enters `r $path`, where `$path` is the directory path clicked by the user. The `r` program receives the `$path` parameter, combines it with the configuration file, generates and executes the SSH command to open the remote terminal.

The `/workspace/logs` directory in the container contains logs for code-server, litessh, and each Rclone service.