FROM node:18-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUn npm run build

EXPOSE 3000

CMD ["npm", "start"]