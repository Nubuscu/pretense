FROM node:16-slim as base
WORKDIR /app
COPY package-lock.json /app/package-lock.json
COPY package.json /app/package.json
RUN npm i
COPY . /app/

FROM base as prod
RUN npm run build
CMD ["node", "build/index.js"]