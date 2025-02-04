# linebreak
Text wrapper algorithms

## Algorithms

### Knuth-Plass
Knuth-Plass implements a simplified version of the [Knuth-Plass line breaking algorithm](https://en.wikipedia.org/wiki/Knuth%E2%80%93Plass_line-breaking_algorithm). It wraps text (provided as a slice of words) into lines that do not exceed a specified maximum width. The algorithm uses dynamic programming to determine optimal breakpoints based on a cubic penalty function that discourages ragged lines.

## Features

- **Dynamic Programming:** Efficiently computes optimal line breaks.
- **Cubic Penalty:** Uses a cubic penalty for extra spaces (except on the last line) to promote balanced text.
- **Robust Input Handling:** Validates inputs and handles edge cases like empty word slices or non-positive maximum widths.
- **Preallocated Slices:** Optimized memory usage by preallocating slices where appropriate.
- **Unit Tested:** Includes comprehensive unit tests to ensure functionality remains correct through future changes.

## Installation

To use the package in your project, you can get it using Go modules:

```bash
go get github.com/gqgs/linebreak
```

## Contributing
Contributions are welcome! If you have suggestions for improvements or bug fixes, please open an issue or submit a pull request. When contributing, please ensure that your changes are well-tested and that they adhere to the project's coding style.

## License
This project is licensed under the Unlicense License. See the LICENSE file for details.

## Acknowledgments
This project was inspired by the original Knuth-Plass line breaking algorithm, which is widely used in typesetting systems to produce visually appealing text layouts.
