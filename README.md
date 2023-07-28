# Intro
This solution uses golang with the Echo API framework, Postgres, Docker, and Kubernetes.


# Scaling
Scaling this solution is done via Kubernetes and Docker. Some of the logic and resource specifics of these parts of the project are missing or otherwise not what they would be in a real world environment.


## Kubernetes & Docker
The Echo server is containerized as a Docker image, which would normally be pushed to Docker hub during the deployment stage of a hypothetical development pipeline.
We have a kubernetes deployment for our server, see `k8s/examples/api-deployment.yml`, which uses our server docker image to run the application as a pod. Increasing or decreasing the number of these pods running simultaneously allows us to scale the application, and is handled via the HorizontalPodAutoscaler defined at `k8s/example/api-autoscalar.yml`.


Similarly, our database is scaled by deploying the Postgres docker image via a kubernetes deployment and controlled via a HorizontalPodAutoscaler. As traffic changes, our autoscaler will create/destroy postgres instances/pods to match. These pods will share a persistent volume claim, allowing our data to exist independently from our database instances.


### Note
This autoscaled deployment solution is effective only to a point. Depending on how high we were expecting to need to scale, a better solution may be to use kubernetes statefulsets instead of deployments. The difference between these two is how they handle the persistent volume claim. In a system using deployments, all database instances will share one volume, meaning there is a bottleneck due to all replicas using the same file system. In a system using deployments, each pod is created with its own volume created using a PVC (persistent claim volume) template. The drawback is that stateful sets treat pods much less ephemeral than in a deployment setup. I went with deployments in my example k8s manifests, mainly because they are a bit easier to read, making it easier to demonstrate how we would be scaling this system.


# How-to's


## Testing
Unit testing can be easily run from the repo root with
`go test ./...`


## Running with Docker
This project works fine running locally on docker images. Use `docker compose up` to start both the database and the Echo API. The database is set up with a raw sql file which sets up our table and loads a little bit of dummy data. The API can be reached at `localhost:1323`, like so:


    curl -s -X PUT "localhost:1323/event/example.com/delivered"
    curl -s -X PUT "localhost:1323/event/example.com/bounced"
    curl -s "localhost:1323/domains/example.com"





