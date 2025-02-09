# LiteSSH

> 轻量级的多服务器管理，强大的 VSCode 编辑体验，让远程工作更加高效无缝

LiteSSH 具有以下优点

- 强大的文件编辑体验：基于 VSCode，支持语法高亮、智能提示、快捷键
- 多服务器支持：在多个服务器之间复制、移动文件
- 开源：防止恶意的密钥泄露，允许进一步的自定义
- Web：在浏览器中运行，无需在多台设备上安装客户端
- 配置同步：在后端统一保存服务器配置和密钥，无需在多台设备上重复配置
- 迅速响应：后端与服务器长期维护 SSH 连接，而不是打开客户端后才开始连接
- 插件生态：兼容 VSCode 插件

VSCode Remote 用于管理服务器时非常便利。其布局为左侧文件栏，右边文件内容和终端。得益于文件栏、终端的联动和 VSCode 本身强大的文本编辑功能，VSCode Remote 具有远好于 XSHELL、XFTP、WinSCP、FinalShell 的使用体验。然而，VSCode Remote 具有以下缺点

- 缺乏多服务器支持：只能一个个点开需要连接的服务器，而且无法在服务器之间复制、移动文件
- 不开源：尽管 VSCode 本体是开源的，但是 VSCode Remote 不开源
- 安装 Agent：需要在服务器上安装 Agent，因此产生了各种问题。
    - Agent 安装/更新时下载失败
    - Agent 基于 Node，无法安装在某些特定的发行版上
    - Agent 无法运行在太老的发行版上
    - Agent 安装可能会违反服务器的安全管理规则
- 客户端安装：需要在每个设备上安装、持续更新 VSCode
- 配置难以同步：需要在每个设备上反复配置服务器连接信息、密钥
- 响应慢：连接服务器时，会新建窗口、建立连接，每个服务器都需等待好几秒。网络中断后，需要痛苦地重复以上流程

LiteSSH 在 VSCode 的基础上，避免了对 Agent 的依赖，因此在运维场景下拥有更高的效率。但是，Agent 的缺失使 LiteSSH 插件无法与服务器环境交互，因此复杂的远程开发仍然需要 VSCode Remote。（当然，也可以通过 code-server 实现远程开发）

## 快速开始

准备 Docker Compose 文件

```sh
git clone https://github.com/117503445/litessh.git
```

docker pull registry.cn-hangzhou.aliyuncs.com/117503445/litessh && docker tag registry.cn-hangzhou.aliyuncs.com/117503445/litessh 117503445/litessh
