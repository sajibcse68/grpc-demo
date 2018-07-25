FROM ubuntu:14.04

RUN mkdir /server
ADD server/myapp /server/

CMD ["/server/myapp"]

EXPOSE 8088