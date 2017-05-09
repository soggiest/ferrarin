FROM golang:alpine

ADD /bin/ferrarin /ferrarin 
ENTRYPOINT ["/ferrarin"]

