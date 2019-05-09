FROM node:8-slim

WORKDIR /server

COPY package-lock.json package.json /server/
RUN npm install

COPY . /server

EXPOSE 3000
CMD [ "npm", "start" ]