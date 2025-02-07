npm install --global yo generator-code

su builder
cd ~
yo code

Ignored build scripts: @playwright/browser-chromium, esbuild. Run "pnpm approve-builds" to pick which dependencies should be allowed to run scripts.

Done in 33s

Your extension litessh has been created!

To start editing with Visual Studio Code, use the following commands:

     code litessh

Open vsc-extension-quickstart.md inside the new extension for further instructions
on how to modify, test and publish your extension.

To run the extension you need to install the recommended extension 'connor4312.esbuild-problem-matchers'.

For more information, also visit http://code.visualstudio.com and follow us @code.


[builder@vscode-lite-ssh-dev ~]$ 


pnpm install
pnpm approve-builds
pnpm run package-web

npm install -g @vscode/vsce
vsce package
code-server --install-extension litessh-0.0.1.vsix
code-server --install-extension ext1-0.0.1.vsix

pacman -Sy --noconfirm rclone
pacman -Sy --noconfirm sshpass
pacman -Sy --noconfirm fuse3


/root/.config/rclone/rclone.conf
mkdir /remote
rclone mount s: /remote --allow-non-empty --allow-other --vfs-cache-mode full -vvv
sshpass -p 123456 ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@127.0.0.1

