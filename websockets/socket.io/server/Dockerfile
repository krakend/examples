FROM node:22-alpine
ENV NODE_ENV=production

WORKDIR /app

COPY ["package.json", "./"]

RUN npm install --production

COPY server.js .

CMD ["node", "server.js"]
