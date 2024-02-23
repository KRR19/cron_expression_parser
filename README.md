# Cron Expression Parser

Cron Expression Parser is a command-line utility designed to parse and validate cron expressions. Cron expressions are used to schedule tasks at specific times, and this tool helps in parsing and validating those expressions, ensuring they are correctly formatted and can be used to schedule tasks as intended.

## Features

- **Parsing**: The tool can parse cron expressions and break them down into their components (minutes, hours, day of the month, month, day of the week, and command).
- **Validation**: It validates the syntax of cron expressions to ensure they are correctly formatted.
- **Console Interaction**: Users can input cron expressions directly into the console, and the tool will parse and validate them on the fly.
- **Argument Support**: The tool can also accept cron expressions as command-line arguments.

## Installation

To install the Cron Expression Parser, clone the repository and build the project using the provided Makefile.

```bash
git clone https://github.com/KRR19/cron_expression_parser.git
cd cron_expression_parser
make build
```

This will compile the project and create an executable named `cron_expression_parser` in the `./bin/` directory.

## Usage

### Running the Parser

To run the parser, execute the following command:

```bash
./bin/cmd
```

### Parsing Cron Expressions

You can input cron expressions directly into the console, and the tool will parse and validate them. For example:

```
Enter a cron expression in correct format.
Example '*/15  0  1,15 *  1-5 /usr/bin/find'.
Press 'q' to quit.
*/15  0  1,15 *  1-5 /usr/bin/find
```

### Using Command-Line Arguments

Alternatively, you can pass a cron expression as a command-line argument:

```bash
./bin/cmd '*/15  0  1,15 *  1-5 /usr/bin/find'
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.