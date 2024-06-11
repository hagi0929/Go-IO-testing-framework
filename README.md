# Golang Testing Framework for Text-Based Input and Output

This project is a comprehensive Golang testing framework designed to streamline the testing of text-based input and output for programs. The framework is equipped with several utility scripts that simplify common tasks, enhancing the ease of use and efficiency for developers. This framework was created to make the CS240 programming assignments easier and more efficient.

## Configuration

The configuration is stored in a JSON file (`setup.json`) with the following structure:

```json
{
	"source-folder": "src",
	"test-folder": "test_cases",
	"test-format-type": "manual-pair",
	"test-format": {
		"io-pairs": [
			{
				"input": "input1.txt",
				"output": "output1.txt"
			},
			{
				"input": "input2.txt",
				"output": "output2.txt"
			}
		]
	}
}
```

## Usage

1. **Clone the repository** and navigate to the project directory.
2. **Edit the `setup.json`** file to match your source folder and test cases.
3. **Run the application**:

    ```sh
    go run main.go
    ```

## Implementation Details

### config/config.go

Contains the configuration logic for loading and parsing the JSON config file.

### main.go

The entry point of the application.

### models/io_file.go

Defines the `IOFile` model.

### test/manual_pair.go and test/pattern.go

Define the test formats and their implementations.

### utils/file_utils.go and utils/json_utils.go

Utility functions for file operations and JSON formatting.
