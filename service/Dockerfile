FROM node:22-alpine

RUN apk add --update curl

WORKDIR /app

COPY . .

EXPOSE 3000

CMD ["node", "index.js"]