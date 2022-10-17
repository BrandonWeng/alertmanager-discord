FROM scratch
MAINTAINER Brandon Weng <wengbrandon@gmail.com>
COPY alertmanager-discord.linux /alertmanager-discord

ENV LISTEN_ADDRESS=0.0.0.0:9094
EXPOSE 9094
ENTRYPOINT [ "/alertmanager-discord" ]
