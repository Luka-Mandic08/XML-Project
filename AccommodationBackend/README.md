# microservices demo

The repository contains an online clothing store microservice application. It serves the purpose of demonstrating microservice architecture and the following commonly used patterns:
- API Gateway
- API Composition
- Saga

The application comprises four microservices (Catalogue, Inventory, Ordering, Shipping) and an API Gateway, all written in Go.

## Quickstart

1. Clone this repository

```
git clone https://github.com/tamararankovic/microservices_demo
cd microservices_demo
```

2. Run the following command

```
docker-compose up --build
```

## Architecture

![](/diagrams/architecture.png)

## Endpoints

- GET http://localhost:8000/catalogue/product - retrieves all products
- GET http://localhost:8000/catalogue/product/{id} - retrieves a product by its id
- GET http://localhost:8000/order - retrieves all orders
- GET http://localhost:8000/order/{id} - retrieves an order by its id
- GET http://localhost:8000/order/{id}/details - retrieves an order's detailed information
- POST http://localhost:8000/order - places a new order
- GET http://localhost:8000/order - retrieves inventory information
- GET http://localhost:8000/shipping/order - retrieves shipping information for all orders
- GET http://localhost:8000/shipping/order/{id} - retrieves shipping information for the specified order

## API Composition

API Gateway plays the role of an API Composer when detailed information on an order is requested. It queries Catalogue, Ordering and Shipping service, after which it aggregates the collected data.

![](/diagrams/composition.png)

## Saga

Orchestration-based saga is implemented in Ordering, Inventory and Shipping services and is triggered when a user tries to place an order. Create Order Orchestrator is part of the Ordering service. Saga orchestrator states based on commands and corresponding replies are displayed below.

![](/diagrams/saga.png)
