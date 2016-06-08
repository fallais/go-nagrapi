FROM        frolvlad/alpine-glibc:latest
MAINTAINER  François ALLAIS <francois.allais@sogeti.com>

ADD go-nagrapi /usr/bin

VOLUME /usr/bin
EXPOSE     5555
CMD        [ "/usr/bin/go-nagrapi", "--s", "/usr/bin/status.dat" ]
