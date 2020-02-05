FROM scratch

WORKDIR /usr/share/thebears

COPY bin/bear bear
COPY bin/cfg.toml cfg.toml

CMD ["/usr/share/thebears/bear", "-c", "/usr/share/thebears/cfg.toml", " >> /dev/null 2>&1"]