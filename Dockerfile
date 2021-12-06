FROM alpine:latest

RUN mkdir -p /opt/goapp

WORKDIR /opt/goapp

COPY ./configs/db/migrate ./configs/db/migrate
COPY .env .
COPY ./build/gorestfull .

RUN chmod +x ./gorestfull

EXPOSE 8080

ENTRYPOINT [ "/opt/goapp/gorestfull" ]
CMD [ "serve" ] 
