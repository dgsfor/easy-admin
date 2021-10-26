FROM alpine:v3.7-1

MAINTAINER dgsfor@gmail.com

WORKDIR /www/easy-api-template.xx.com
COPY easy-api-template easy-api-template
ADD conf conf
RUN chmod +x easy-api-template

# api
EXPOSE 8081

CMD ["./easy-api-template", "--env", "prod"]
