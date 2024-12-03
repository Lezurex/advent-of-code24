# Advent of Code 2024 🎄

This repository contains my solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges, written
in **Go**. However, they may not be the most concise nor the most optimized.

## Project Structure

```plaintext
advent-of-code/
│
├── go.mod
├── go.sum
├── main.go                # Entry point for running specific days
│
├── solutions/             # Daily solutions
│   ├── day01.go           # Solution for Day 1
│   ├── day02.go           # Solution for Day 2
│   └── ...
│
├── inputs/                # Input files for each day
│   ├── day01.txt          # Input for Day 1
│   ├── day02.txt          # Input for Day 2
│   └── ...
│
└── utils/                 # Reusable utilities
    └── file.go            # Functions for reading input files
```

## Getting Started

1. Clone this repository:
   ```bash
   git clone https://github.com/Lezurex/advent-of-code24.git
   cd advent-of-code24
   ```

2. Prepare your input files:
    - Place your inputs in the `inputs/` directory.
    - Name them as `dayXX.txt` (e.g., `day01.txt` for Day 1).

3. Run a solution:
    - Open `main.go` and set the `day` variable to the desired day.
    - Run the program:
      ```bash
      go run main.go
      ```

Happy coding! 🎅