
``` 
- District VulnWeb

- David José Cardona Nieves, djcardonan@eafit.edu.co
- Julián Andrés Ramírez Jiménez, jaramirezj@eafit.edu.co
```

---

## Table of contents[](#table-of-contents)
- [Introduction](#introduction)
- [Motivation](#motivation)
- [Composition](#composition)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Cloning the repository](#cloning-the-repository)
  - [Setting up the database](#setting-up-the-database)
  - [Running the webapi](#running-the-webapi)
  - [Running the webapp](#running-the-webapp)
- [Roadmap](#roadmap)
- [Call to action](#call-to-action)

---

## Introduction[](#introduction)

**District** is a complete educational web application designed to provide a simulated environment for students participating in Exfil Security's training program. Developed with a focus on enhancing cybersecurity skills, this project serves as a hands-on learning platform for aspiring security professionals.

## Motivation[](#motivation)

This project was developed as part of the internship project in conjunction with EAFIT University. The general idea was to be able to provide the next groups participating in the Exfil Security training program with the resources to apply their knowledge in a practical setting as they learn.
Our approach to introducing vulnerabilities into the District VulnWeb application was guided by reputable sources and industry-standard testing guides. 

## Composition[](#composition)

Generally speaking, **District** has as main components:

- **Backend API:**
    - **Programming Language:** Golang.
    - **Framework:** Echo.
    - **Database:** PostgreSQL.

- **Frontend Server:**
    - **Programming Language:** JavaScript.
    - **Framework:** VueJS.

We aim for a concise API design structure, comprising the following integral components:
-   **Controllers:** Serving as the interface to external requests, controllers manage the accurate handling of incoming requests and generation of responses. They orchestrate interactions between services, ensuring the smooth execution of business logic and enhancing overall API responsiveness.
-   **Services:** Positioned at the core, these components encapsulate crucial business logic, delivering essential functions to API consumers. This separation of concerns enhances code modularity, scalability, and overall codebase clarity.
-   **Repositories:** Responsible for the management and execution of CRUD operations in the system models. Also in charge of the setup of the tables in the database.

---

## Getting Started[](#getting-started)

### Prerequisites[](#prerequisites)

* **Golang:** The Go language is a prerequisite for running the API and can be obtained from the [following link](https://go.dev/doc/install).
*  **Node Package Manager (npm):** npm is required to be installed in order to correctly setup the web application. [Link here](https://www.npmjs.com/package/npm).
* **Docker-compose:** You need to have docker-compose installed since the database is run through containers.

### Cloning the repository[](#cloning-the-repository)

First, you need to clone the repository from GitHub. To do this, you can open a terminal and run the following command:

```
git clone https://github.com/DavidLakeT/district.git
```

### Setting up the database[](#setting-up-the-database)

```
cd docker
docker-compose up -d
```

### Running the webapi[](#running-the-webapi)

```
cd app/webapi
go run .
```

### Running the webapp[](#running-the-webapp)

```
cd app/webapp
./init.sh
```

---

## Roadmap[](#roadmap)

One of the main objectives of District beyond the implementation of vulnerabilities is to achieve a clean and scalable development that allows future members to easily join the project and contribute to a code base built following best practices and development strategies.

Here are some of the aspects that could be considered in the future for the implementation of the project:
- Implementation of new vulnerabilities.
- Configuration resources that allow to establish the type of vulnerability to be used at the time.

These features are open to public intention to contribute to open-source development.

## Call to action[](#call-to-action)

If there is a feature that you think could be interesting for the application, please open a [new issue](https://github.com/DavidLakeT/district/issues) or start a [discussion](https://github.com/DavidLakeT/district/discussions).

Feel free to contact us if you have any questions or contributions to make to the project.
