FROM golang:alpine

ADD /bin/ferrarin /ferrarin 
ENTRYPOINT ["/bin/sh", "/ferrarin"]

