# Dynamic CRUD API CDK Project

This repository contains a project for creating a dynamic CRUD (Create, Read, Update, Delete) API using AWS CDK (Cloud Development Kit) with TypeScript. The API is designed to be flexible and scalable, allowing for dynamic handling of various entities.

## Project Structure

- **bin/**: Contains the entry point for the CDK application.
- **lib/**: Contains the CDK stack definition.
- **src/**: Contains the Go source code for the Lambda function handling the API requests.

## Useful Commands

- `npx cdk deploy` - Deploy this stack to your default AWS account/region
- `npx cdk diff` - Compare deployed stack with current state

## API Endpoints

The API provides the following endpoints for dynamic CRUD operations:

### Health Check

- **GET /**: Returns the health status of the API.

### Dynamic CRUD Operations

- **GET /api/:entity**: Retrieve all records of a specified entity.
  - **Parameters**: 
    - `entity` (path) - The name of the entity to retrieve records for.

- **GET /api/:entity/:id**: Retrieve a single record by ID.
  - **Parameters**: 
    - `entity` (path) - The name of the entity.
    - `id` (path) - The ID of the record to retrieve.

- **POST /api/:entity**: Create a new record for a specified entity.
  - **Parameters**: 
    - `entity` (path) - The name of the entity.
  - **Body**: JSON object representing the new record.

- **PUT /api/:entity/:id**: Update an existing record by ID.
  - **Parameters**: 
    - `entity` (path) - The name of the entity.
    - `id` (path) - The ID of the record to update.
  - **Body**: JSON object representing the updated record.

- **DELETE /api/:entity/:id**: Delete a record by ID.
  - **Parameters**: 
    - `entity` (path) - The name of the entity.
    - `id` (path) - The ID of the record to delete.

## Setting Up Local API Using SAM

To set up the local API for development using AWS SAM (Serverless Application Model), follow these steps:

1. **Install AWS SAM CLI**: Follow the [installation guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html) to install the SAM CLI.

2. **Build the SAM Application**:

   ```sh
   sam build
   ```

3. **Start the API Locally**:

   ```sh
   sam local start-api
   ```

4. **Invoke the API**: You can now send HTTP requests to the local endpoint, for example `http://localhost:3000/api/<entity-name>`.

## License

This project is licensed under the MIT License.
