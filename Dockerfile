FROM postgres

COPY ./script.sql /root/

WORKDIR /root/