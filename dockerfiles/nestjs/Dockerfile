# Stage 1: Build the app
FROM node:18 AS builder

WORKDIR /app

# Copy dependency files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy source files
COPY . .

# Build the NestJS app
RUN npm run build

# Stage 2: Run the built app
FROM node:18-alpine

WORKDIR /app

# Copy only the compiled dist folder and dependencies
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY package*.json ./

# Expose the port NestJS runs on
EXPOSE 3000

# Command to run the app
CMD ["node", "dist/main"]
