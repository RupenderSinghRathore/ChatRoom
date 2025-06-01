![2025-06-01-230531_hyprshot](https://github.com/user-attachments/assets/a48f167e-1e1f-4ce6-851f-97ea732fc6da)
![image](https://github.com/user-attachments/assets/d1d37dbd-954e-4235-884e-d54831bdeb21)


# Go CLI ChatRoom

A simple command-line interface (CLI) chat room application built with Go, utilizing the `github.com/gorilla/websocket` package for real-time communication between multiple users.

## Features

* Real-time messaging between multiple clients.
* Simple CLI-based interaction.
* Easy to set up and run locally.

## Technology Stack

* **Go (Golang)**
* **WebSockets:** `github.com/gorilla/websocket` package

## Prerequisites

* [Go](https://golang.org/dl/) installed on your system (version 1.24.3 or higher recommended).
* [Git](https://git-scm.com/downloads) for cloning the repository.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/RupenderSinghRathore/ChatRoom.git
    ```

2.  **Navigate to the project directory:**
    ```bash
    cd ChatRoom
    ```

## Usage

You'll need to start the server first, and then one or more clients can connect to it.

1.  **Start the Server:**
    Open a terminal window and run the following command from the project's root directory (`ChatRoom`):
    ```bash
    go run ./cmd/cli/ start
    ```
    You should see a confirmation message that the server has started (e.g., "Server started on :8080" or similar).

2.  **Connect a Client:**
    Open a new terminal window for each client you want to connect. From the project's root directory (`ChatRoom`), run the following command, replacing `<username>` with the desired nickname for the client:
    ```bash
    go run ./cmd/cli/ connect -u <username>
    ```
    For example:
    ```bash
    go run ./cmd/cli/ connect -u Alice
    ```
    Once connected, you can start sending and receiving messages with other connected users.

## Next Steps (Development Roadmap)

* Implement robust unit tests.
* Explore various deployment technologies.

## Contributing

Contributions are welcome! If you have suggestions or want to improve the application, please feel free to open an issue or submit a pull request.
