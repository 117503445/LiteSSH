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

RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay.git yay && cd yay && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

RUN su - builder -c "yay -Su code-server --noconfirm"

RUN chsh -s /usr/bin/fish
COPY ./scripts/config.fish /root/.config/fish/config.fish

WORKDIR /workspace