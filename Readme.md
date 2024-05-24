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
- `-limit`: Specify the number of questions (optional, default is 1).

### Example

```sh
go run main.go -a your_api_key -c Programming -d Easy -limit 5
```

<!-- ## Code Structure

- **main.go**: The main file containing the entry point of the application and all function implementations.

### Functions

- `Config(key string) string`: Loads configuration from the `.env` file.
- `main()`: The entry point of the application. Reads arguments, fetches questions, starts the quiz, and displays the score.
- `read_arguments() (string, string, string, int)`: Reads command-line arguments.
- `begin_quiz(questions []QuizResponse) int`: Starts the quiz and returns the user's score.
- `get_questions(api_key string, category string, difficulty string, limit int) []QuizResponse`: Fetches quiz questions from QuizAPI.
- `show_options(answers map[string]string)`: Displays answer options for a question.
- `exit(msg string)`: Exits the application with a message.
- `check_answers_value(correct_answers map[string]string) int`: Checks the user's answer and returns the score for the question.

### Types

- `QuizResponse`: Represents the structure of a quiz question response.
- `ErrorResponse`: Represents the structure of an error response. -->

<!-- ## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. -->

<!-- ## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. -->

## Acknowledgements
- [QuizAPI](https://quizapi.io/) for providing the quiz questions API.
