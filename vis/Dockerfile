FROM node:16-slim
WORKDIR /app
COPY package-lock.json /app/package-lock.json
COPY package.json /app/package.json
RUN npm i
COPY . /app/
CMD ["npm", "run", "dev"]