# The Botory Project

## About This Project
* * *
- Botory (“Bot + Factory”) is an automation tool that generates and deploys chatbots from a single YAML file. 
Provide your options and comments in YAML, and Botory takes care of building and deploying the chatbot.
- It helps developers integrate chatbots quickly without spending time on repetitive non-core tasks.

## Participants
* * *
- Developer: Jo Sangyun
- :school: **School**: Seoultech University
- :globe_with_meridians: **Languages**: Go, C, C++
- :computer: **Skills**: Docker, Linux, Git 
- :email: **Email**: sengyun8908@naver.com
- :octocat: **Github**: https://github.com/joseng8908

## Motivation
* * *
- In real projects, adopting chatbots often involves too many non-core tasks (configuration, deployment, and operational setup).
Botory automates these steps so developers can focus on core logic.  
 
## Used Technologies
* * *
- **Go**: Core implementation
- **Docker**: Packaging and deployment 

## Project Goals
* * *
- Provide a declarative, YAML-based workflow to automate chatbot creation, build, and deployment.
- Enable production-ready chatbots from minimal input (YAML only).
- Gain proficiency in Go and Docker, and release as open source.

## Manual

### 1. Quick Start (Recommended: Using Docker)

This is the easiest and recommended way to get your chatbot server up and running.

**Prerequisites:**
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) must be installed and running on your system.

**Steps:**
1.  Clone this repository to your local machine:
    ```bash
    git clone https://github.com/joseng8908/botory_project.git
    cd botory_project
    ```

2.  Customize your chatbot's behavior by editing the `configs/chatbot.yaml` file.

3.  Run the following command in your terminal. Docker will build the image and start the container.
    ```bash
    docker compose up --build
    ```
    
4.  That's it! Your chatbot is now running and accessible at `http://localhost:8080/chat`. You can test it with `curl`:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"message": "안녕"}' http://localhost:8080/chat
    ```

### 2. Running without Docker (Native Binary)

If you prefer not to use Docker, you can run the pre-compiled binary directly.

**Prerequisites:**
- None!

**Steps:**
1.  Go to the [**Releases**](https://github.com/joseng8908/botory_project/releases) page of this repository.
2.  Download the binary that matches your operating system (e.g., `botory_windows_amd64.exe` or `botory_linux_amd64`).
3.  Download the `configs` folder (or just the `chatbot.yaml` file) and place it in the same directory as the downloaded binary.
4.  Customize `configs/chatbot.yaml` to your needs.
5.  Open a terminal in that directory and run the program:

    *   **On Windows:**
        ```bash
        .\botory_windows_amd64.exe start
        ```
    *   **On macOS / Linux:**
        ```bash
        # You might need to grant execute permission first
        chmod +x ./botory_darwin_amd64
        ./botory_darwin_amd64 start
        ```

### 3. YAML File Format

The `chatbot.yaml` file has a simple structure:

-   **`botName`**: The name of your bot.
-   **`defaultResponse`**: The response sent when no keyword matches.
-   **`dialogs`**: A list of keyword-response pairs.
    -   **`keyword`**: The user's input message to trigger the response.
    -   **`response`**: The bot's reply for the corresponding keyword.
    -   **`matchType`**: The type of matching algorithm to use.
        -   **`exact`**: Matches the keyword exactly.
        -   **`contains`**: Matches if the keyword is a substring of the user's input.
```yaml
botName: "봇토리"
defaultResponse: "무슨 말씀이신지 잘 모르겠어요."

dialogs:
  - keyword: "안녕"
    response: "안녕하세요! 만나서 반가워요."
  
  - keyword: "이름이 뭐야"
    response: "제 이름은 봇토리입니다."
    matchType: exact

  - keyword: "가격"
    response: "N$ 입니다"
    matchType: contains
```
Thank you!😊
