# QuizApp README

## Overview

This is a Goland command-line implimentation of [QuizAPI-BASH](https://github.com/QuizAPI/QuizAPI-BASH) by the [QuizAPI](https://quizapi.io/) team.It fetches quiz questions from the QuizAPI and quizzes the user. The application allows users to specify the category, difficulty, and number of questions they want to be quizzed on. It then scores the user's responses and displays the final score.

## Features

- Fetch quiz questions from the QuizAPI.
- Customize quiz by category, difficulty, and number of questions.
- Interactive command-line interface for answering questions.
- Scoring system to evaluate user performance.

## Requirements
- An API key from [QuizAPI](https://quizapi.io/).

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/MogboPython/quizapp.git
   cd quizapp
   ```

2. **Install dependencies**:
   ```sh
   go get -u github.com/joho/godotenv
   ```

3. **Set up the environment**:
   - Create a `.env` file in the root directory.
   - Add your QuizAPI key to the `.env` file:
     ```env
     API_KEY=your_api_key_here
     ```

## Usage

To run the application, use the following command-line arguments:

```sh
go run main.go -a API_KEY [-c Category] [-d Difficulty] [-limit NumberOfQuestions]
```

### Arguments

- `-a`: Your QuizAPI key (required).
- `-c`: Specify a category (optional). Valid options include: `Linux`, `DevOps`, `Networking`, `Programming`, `Cloud`, `Docker`, `Kubernetes`.
- `-d`: Specify a difficulty level (optional). Valid options are: `Easy`, `Medium`, `Hard`.
- `-limit`: Specify the number of questions (optional, default is 1, maximum of 20).
- `-time`: Specify the time limit for the quiz in seconds (optional, default is 30).

### Example

```sh
go run main.go -a your_api_key -c Programming -d Easy -limit 5 -time 60
```

## Acknowledgements
- [QuizAPI](https://quizapi.io/) for providing the quiz questions API.
