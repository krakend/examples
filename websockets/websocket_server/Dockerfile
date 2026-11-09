FROM node:22-alpine
ENV NODE_ENV=production

WORKDIR /app

COPY ["package.json", "./"]

RUN npm install --production

COPY . .

CMD ["node", "server.cjs"]
