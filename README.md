# YAML to JSON Converter CLI

## Overview
This is a command-line tool built in Go using the spf13 Cobra framework to convert YAML files to JSON format.

## Features
- Converts YAML files to JSON
- Supports conversion of single files

## Installation
1. Make sure you have Go installed on your system. If not, you can download and install it from [here](https://golang.org/doc/install).
2. Clone this repository:

   ```bash
   git clone git@github.com:manik23/YAML2JSON.git
   ```

3. Change to the directory of the cloned repository:
    ```   
    cd YAML2JSON
    ```

4. Build the CLI tool:
    ```bash
    make build
    ```

5. Move the built binary to your $PATH to make it accessible system-wide:

    ```bash
    sudo mv yamlToJson /usr/local/bin/
    ```

## Usage
```bash
yamlToJson parse --path pod.yaml --o output.json
``` 

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please feel free to open an issue or submit a pull request.