FROM node:8-slim

WORKDIR /server

COPY package-lock.json package.json /server/
RUN npm install

COPY . /server
RUN npx eslint "src/**"

EXPOSE 3000
CMD [ "npm", "start" ]