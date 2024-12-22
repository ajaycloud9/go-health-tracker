FROM ajaysc30/go-health-tracker:ui-base-22.11.0

# Set the working directory inside the container
WORKDIR /go-health-frontend
 
# Copy package.json and package-lock.json
COPY package*.json ./
 
# Install dependencies
RUN npm install
 
# Copy the rest of your application files
COPY . .
 
# Expose the port your app runs on
EXPOSE 3000
 
# Define the command to run your app
CMD ["npm", "start"]