# Go Quiz Application

This is a simple command-line quiz application written in Go.
It reads questions and answers from a CSV file and presents them to the user.
The user has a specified time limit to answer each question, after which the quiz ends and the user's score is displayed.

## Usage

### Prerequisites

Make sure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://golang.org/).

### Installation

Clone this repository to your local machine:

```bash
git clone <repository-url>
cd <repository-directory>
```

### Running the Application

To run the quiz application, use the following command:

```bash
go run main.go -csv <csv-file> -timelimit <time-limit>
```

Replace <csv-file> with the path to your CSV file containing questions and answers,
and <time-limit> with the desired time limit for the quiz in seconds.

Example:

```bash
go run main.go -csv problems.csv -timelimit 30
```

### CSV File Format

The CSV file should have two columns: the first column containing questions and the second column containing corresponding answers.
Make sure there are no headers in the CSV file.

Example CSV file content:

```bash
What is 2+2?,4
Who is the president of the USA?,Joe Biden
```

### Note

- Ensure that the CSV file path is correct and accessible.
- The application will wait for user input for each question. Type in the answer and press Enter to submit.
- If the time limit for the quiz is exceeded, the quiz will end automatically.
- Your score will be displayed at the end of the quiz.

### License

This project is licensed under the MIT License.

```bash
Feel free to adjust and expand upon this README file based on your specific project needs.
```
