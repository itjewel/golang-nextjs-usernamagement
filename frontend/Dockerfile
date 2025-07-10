# Install dependencies
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package files and install deps
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Build the app
RUN npm run build

# Start Next.js server
EXPOSE 3000
CMD ["npm", "start"]
