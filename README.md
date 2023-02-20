<a name="readme-top"></a>

[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![made-with-golang](https://img.shields.io/badge/Go-v1.19-blue)](https://go.dev/)


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><<a href="#installation">Installation</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This Project is a quiz game where the user participates in a mathematical quiz which is time managed using concurrency concepts such as channels and go routines. The user can pass in their own csv file or set the time of the quiz instance where all the questions are randomly shuffled before the quiz begins. The purpose of this project is to practice golang's concurrency concepts.


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

Languages:
* Golang
* Shell Script

Libraries:
* encoding/csv
* math/rand
* time
* io
* os
* strings
* log
* flag

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

Clone the repository, via HTTP or SSH above, download golang v1.19 if you havent already, finally start the program by running the shell script file `./run.sh` or manually putting in values for the optional argument flags when running `main.go`. The optional arguments are `-csv` which is the csv you're passing in, if you make your own it should be formatted as `problem,answer` per line, as well as `-limit` which is amount of time in seconds the programming will run.

### Installation

1. Clone the repo
  ```sh
  git clone https://github.com/dnnysoftware/GolangMathQuiz
  ```
2. Give file execute permissions to run.sh
  ```sh
  chmod +x run.sh
  ```
3. Run Program
  * In root directory run by typing in CLI
  ```sh
  ./run.sh
  ```
  * Or you can manually run the program and providing values to optional arguments by doing:
  ```sh
  go build .
  go run main.go -csv=[csvFilenameInRoot] -limit=[integerValueInSeconds]
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>