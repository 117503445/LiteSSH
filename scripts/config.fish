if status is-interactive
    set fish_greeting # Disable greeting

    # set -x all_proxy "socks5://127.0.0.1:1080"; set -x http_proxy $all_proxy; set -x https_proxy $all_proxy
    set -x PATH ~/.local/bin ~/go/bin $PATH
end