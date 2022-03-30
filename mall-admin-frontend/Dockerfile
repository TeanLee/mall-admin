# ！注意：node 版本要和本地能执行成功的版本保持一致！！！
FROM node:16.12

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

ENV VUE_APP_API_URL='http://127.0.0.1:8081'

EXPOSE 8080

CMD [ "npm", "run", "serve" ]