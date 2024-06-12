# The Rent App

## An intuitive and feature-rich eCommerce web application with all the pages tailored for renting a diverse range of items online. 

### Platform offers following features - 

* Flexible rent payment options, allowing users to choose from daily, weekly, monthly, or custom rental periods as set by the business owners. 

* A secure payment gateway is crucial for seamless transactions. 

* Tracking of deliveries and maps to help users locate shops from where rental items are purchased. 

* Capability for Businesses to list products for both rent and sale, with customers having the flexibility to rent items as per their preference. 

* Admin panel provides comprehensive inventory management tools for businesses to efficiently manage their product listings. 

* An user-centric interface that promotes trust and facilitates transactions for both businesses and customers.

---

### Tech Stack -

* Frontend - React (18.3.1)
* Backend - Go (1.22.1), PostgreSQL (16.2)
* Others - Docker

---

### Setup - 

* #### Prerequisite 
    * [Docker](https://docs.docker.com/get-docker/)
    * [Node](https://nodejs.org/en/download/package-manager)

* #### Starting the server
    * Navigate to server folder 
    * Run the following command to build and start postgres container in detach mode
        ```bash
        docker compose up rent_db -d
        ```

    * Set DB_MIGRATION_FLAG to Y in .env
    * Run the following command to build and start golang container in detach mode for db migration
        ```bash
        docker compose up rent_api -d
        ```

    * Set DB_MIGRATION_FLAG to N in .env
    * Run the following command to restart golang container in detach mode for running the API
        ```bash
        docker compose up rent_api -d
        ```

* #### Starting the client

    * Navigate to client folder
    * Run the following command to install all dependencies
        ```bash
        npm install
        ```

    * Run the following command to run the react client
        ```bash
        npm start
        ```