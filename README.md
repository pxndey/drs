# drs

Command Line Application (CLI) to view Formula 1 Championship Data (WDC & WCC)

## Features
- Fetch driver standings for a given year.
- Fetch constructor standings for a given year.
- Display race data based on the race number and year.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/pxndey/drs.git
   cd drs
2. Build The Executable:
    ```bash
    go build -o drs.exe
3. Run
    ```bash
    .\drs.exe --flags

## Usage

Flags:

- `-year`        : Specify the race year, leave blank for the current year.
- `-race`        : Specify the race number.
- `--drivers`    : Display the driver's championship standings for the given year.
- `--constructors`: Display the constructor's championship standings for the given year.
- `--help`       : Show usage instructions.