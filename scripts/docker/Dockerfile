FROM archlinux

ARG CHINA_MIRROR
RUN if [ -z "$CHINA_MIRROR" ]; then \
      echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist; \
    else \
      echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist; \
    fi

RUN pacman-key --init && \
    pacman-key --populate && \
    pacman -Sy archlinux-keyring --noconfirm && \
    pacman -Su --noconfirm

RUN pacman -Syu go fish btop git openssh nano vim micro base-devel tmux yazi --noconfirm

RUN if [ -n "$CHINA_MIRROR" ]; then \
        go env -w GOPROXY=https://goproxy.cn,direct; \
    fi

ENV PATH="$PATH:/root/go/bin"

RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay.git yay && cd yay && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

RUN su - builder -c "yay -Su code-server --noconfirm"
RUN go install github.com/117503445/vsc-init@master
# ENTRYPOINT [ "fish" ]
RUN vsc-init

RUN chsh -s /usr/bin/fish
COPY ./scripts/config.fish /root/.config/fish/config.fish

RUN pacman -Syu npm pnpm --noconfirm
RUN if [ -n "$CHINA_MIRROR" ]; then \
        npm config set registry https://registry.npmmirror.com; \
        pnpm config set registry https://registry.npmmirror.com; \
    fi
RUN npm install --global @vscode/vsce

COPY ./ext /tmp/ext
WORKDIR /tmp/ext
RUN vsce package --allow-missing-repository
RUN code-server --install-extension litessh-0.0.1.vsix

COPY go.mod go.sum /tmp/litessh/
WORKDIR /tmp/litessh
RUN go mod download
COPY . /tmp/litessh
RUN go install ./cmd/entrypoint
RUN go install ./cmd/litessh
RUN go install ./cmd/r
RUN rm -rf /tmp/litessh

WORKDIR /workspace
ENTRYPOINT [ "entrypoint" ]