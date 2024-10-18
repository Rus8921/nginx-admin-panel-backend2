# Nginx Admin Panel Backend
# Nginx admin panel backend

## Description

This repository of the our project contains the code for frontend part, you can find research part here https://gitlab.pg.innopolis.university/antiddos/research




## How to Run the Service

This project uses Docker to manage the application. To run the application using Docker, enter shell-command:

```sh
docker-compose up
```

## How to Run the Service Using Makefile

This project uses a Makefile to manage building, running, and cleaning the application. Below are the instructions on how to use the Makefile.

### Prerequisites

- Ensure you have `Go` installed on your machine.
- Ensure you have `make` installed on your machine.

### Makefile Targets

- **build**: Compiles the application.
- **run**: Builds and runs the application.
- **clean**: Cleans the build directory.
- **docker-image**: Builds the Docker image.
- **docker-run**: Runs the Docker image using docker-compose.

### Commands

1. **Build the Application**

    To build the application, run:
    ```sh
    make build
    ```

2. **Run the Application**

    To build and run the application, run:
    ```sh
    make run
    ```

3. **Clean the Build Directory**

    To clean the build directory, run:
    ```sh
    make clean
    ```

### Example

```sh
# Build the application
make build

# Run the application
make run

# Clean the build directory
make clean
```

## Support

You can contact with us in:
email: a.antonian@innopolis.university
telegram: @artom_antonyn

## Project status
work in progres
