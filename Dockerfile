FROM alpine:3.4

COPY bin/ferrarin /usr/local/bin

WORKDIR /

CMD ["/bin/sh","-c","dos2unix","-vf","/usr/local/bin/ferrarin"]

ENTRYPOINT ["ferrarin"]

