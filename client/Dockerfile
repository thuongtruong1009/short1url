FROM node:18-alpine

WORKDIR /client

COPY ./package*.json /client/

RUN npm install && npm cache clean --force

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]


# ENV PATH ./node_modules/.bin/:$PATH