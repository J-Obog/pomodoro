FROM node:lts-alpine3.14

RUN mkdir /react-client

WORKDIR /react-client

COPY . /react-client/

RUN npm install

EXPOSE 3000

CMD ["npm", "start"]
