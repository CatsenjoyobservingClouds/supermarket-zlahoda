#FROM node:16-alpine AS build
#
#WORKDIR /app
#
#COPY package.json .
#COPY package-lock.json .
#RUN npm install
#COPY . .

FROM node:18.16.0

WORKDIR /app

COPY package.json ./
COPY tsconfig.json ./
RUN npm i

COPY . .

CMD ["npm", "run", "start"]