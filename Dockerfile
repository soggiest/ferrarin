FROM golang:alpine

ADD /bin/ferrarin /ferrarin 
#RUN chmod 755 /ferrarin
ENTRYPOINT ["/ferrarin"]

