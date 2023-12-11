FROM node:16-alpine

WORKDIR /usr/src/app

COPY client /usr/src/app/client

EXPOSE 3000

RUN npm install -g http-server nodemon

CMD ["nodemon", "--exec", "http-server", "client", "-p", "3000"]
