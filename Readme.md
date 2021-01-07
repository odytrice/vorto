# DevOps Engineering Challenge

You own a coffee shop. You procure coffee beans from a number of suppliers and have them delivered directly to you. You source 3 types of coffee beans: Arabica, Robusta and Liberica.

Each supplier supplies one or more types of beans.

Drivers work for a carrier, and each carrier can haul one or more types of beans.

Write a SQL query that produces the invalid deliveries.
_(Invalid deliveries are deliveries that a carrier cannot perform due to carrier bean constraints)_

Write a simple Golang server that has an endpoint that will return the results.

You will also need to provide the manifests for deploying the server onto a Kubernetes cluster. (e.g. deployment, service, ingress)

_The SQL dump can be downloaded [here](https://github.com/VortoEng/VortoFiles/blob/master/coffee.sql)_

### Requirements
You will need to deploy your project using Docker/Kubernetes.
Please provide a README along with the link to your project

### Bonus
- Use gRPC for the server and provide a postman endpoint. Please provide the proto files.
- Given that each delivery contains a random bean from the supplier's stock, what is the probability that the delivery is valid. Write an endpoint that will return this result.