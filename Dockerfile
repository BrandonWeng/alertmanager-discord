FROM scratch
MAINTAINER Brandon Weng <wengbrandon@gmail.com>
COPY alertmanager-discord.linux /alertmanager-discord
ENTRYPOINT [ "/alertmanager-discord" ]
