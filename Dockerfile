FROM scratch 

ADD /bin/ferrarin /ferrarin 
ENTRYPOINT ["/bin/sh", "/ferrarin"]

