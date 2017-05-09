FROM scratch 

ADD /bin/ferrarin /ferrarin 
CMD ["/bin/sh", "/ferrarin"]

