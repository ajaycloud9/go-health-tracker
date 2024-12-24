FROM ajaysc30/go-health-tracker:ui-base-22.11.0

# Set the working directory inside the container
RUN mkdir -p /Go
COPY ui /Go/ui

WORKDIR /Go/ui/go-health-frontend
 
# Install dependencies
RUN npm install

# Expose the port your app runs on
EXPOSE 3000/tcp
 
# Define the command to run your app
CMD ["npm", "start"]