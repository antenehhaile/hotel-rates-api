# Hotel Rates API

## Overview
This repository contains a GoLang application that serves as an API for fetching hotel rates. It integrates with an external hotel rates provider (hotelbeds) to retrieve information about available hotels, their rates, and other related details.

## Requirements
- Go 1.16 or higher
- Dependencies listed in `go.mod`

## Local Set up
1. Clone the repository to your local machine:

```bash
git clone <repository-url>
```

2. Download all the go dependencies:

```bash
go mod download
```

3. Set up your local env
```bash
export API_KEY=${API_KEY}
export SECRET=${API_SECRET}
```

4. Run the App
```bash
go run main.go
```

The API will run on port `8080`

## Local Testing

There are 2 exposed endpoints. You can call the endpoints as follows. 

`health`
```bash
curl --location --globoff 'http://localhost:8080/health'
```   
  
`/hotels`
```bash 
curl --location --globoff 'http://localhost:8080/hotels/?checkin=2024-06-15&checkout=2024-06-16&currency=USD&guestNationality=US&hotelIds=264&occupancies=[{%22rooms%22%3A1%2C%20%22adults%22%3A%202}]'
```

## Unit Tests

Run the tests using the following command:
```bash
go test ./...
```

## Deployment with AWS ECS Fargate (GitHub Actions)
This service has been full configured and deployed to AWS fargate. 

This repository also includes a GitHub Actions workflow (`ecs-fargate-deploy.yaml`) for deploying the application to AWS ECS Fargate.

Workflow Details

- Trigger: The workflow is triggered when a new release is tagged in github
- Steps:
  - Build, Tag, and Push Image to Amazon ECR: Builds a Docker image, tags it, and pushes it to Amazon ECR.
  - Fill in the New Image ID: Updates the ECS task definition with the new Docker image ID always set to latest.
  - Deploy Amazon ECS Task Definition: Deploys the updated task definition to the ECS cluster.
  - The task definition is stored in `rate-api-definition.json` in the root directory

### To test the API deployed to AWS ECS Fargate

`/hotels`
```bash
curl --location --globoff 'apl-net-lb-77e5c69e28629a3d.elb.us-east-2.amazonaws.com/hotels/?checkin=2024-06-15&checkout=2024-06-16&currency=USD&guestNationality=US&hotelIds=77%2C168%2C264%2C265%2C297%2C311&occupancies=[{%22rooms%22%3A1%2C%20%22adults%22%3A%202}]'
```

`/health`
```bash
curl --location --globoff 'apl-net-lb-77e5c69e28629a3d.elb.us-east-2.amazonaws.com/health'
```

## Deployment with EC2 Autoscaling Group (GitHub Actions)
This service has been full configured and deployed to AWS EC2 autoscaling group via AWS CodeDeploy. 

This repository includes a GitHub Actions workflow (`ec2-auto-scaling-group.yaml`) for deploying the application to an EC2 Autoscaling Group.

### Workflow Details

- Trigger: The workflow is triggered on every push to the main branch.
- Steps: 
  - Build and Package: Compiles the GoLang application and creates a tarball for deployment.
  - Upload Artifact to S3: Uploads the deployment artifact to an S3 bucket.
  - Deploy: Initiates a deployment to the EC2 Autoscaling Group.

### Additional File 
- `.github/scripts/build.sh`: Bash script for building the GoLang application and creating a deployment artifact.
- `appspec.yml`: AWS CodeDeploy AppSpec file for specifying deployment details.
- `start-application`: A Bash script for creating a systemd service file (hotelrateapi.service) and then reloading systemd daemon and starting the service.

### To test the API deployed to AWS EC2 Auoscaling Group
`/hotels`
```bash
curl --location --globoff '3.142.68.170:8080/hotels/?checkin=2024-06-15&checkout=2024-06-16&currency=USD&guestNationality=US&hotelIds=77%2C168%2C264%2C265%2C297%2C311&occupancies=[{%22rooms%22%3A1%2C%20%22adults%22%3A%202}]'
```

`/health`
```bash
curl --location --globoff '3.142.68.170:8080/health'
```