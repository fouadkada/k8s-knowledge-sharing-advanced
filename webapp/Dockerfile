FROM nginx:1.19.2-alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf *
COPY build/ ./

EXPOSE 80