# Microservices E-commerce 

## Overview

This is a microservices based E-commerce platform built using Go, gRPC and PostgreSQL. This includes multiple microservices like Order and Payment service handling specific functions of the e-commerce feature.

## Features

- **Order Service** : Manage orders and order items. Communicates with the Payment Service to create payment after order.

- **Payment Service** : Process payments and updates payment status. Communicates with the Order Service to update order status after payment.