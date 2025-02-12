# tmdb-cli-tool

A command line interface (CLI) application that fetches and displays data from The Movie Database (TMDB) API. Users can specify the type of movies they want to view, including currently playing, popular, top-rated, and upcoming movies.

This project is part of the backend section of the [roadmap.sh](https://roadmap.sh/projects/tmdb-cli) website.

## Features

- Fetch and display movie data from TMDB API.
- Command line arguments to specify movie types:
  - `playing`
  - `popular`
  - `top`
  - `upcoming`

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/alielmi98/tmdb-cli-tool.git
   ```
2. Navigate to the project directory:
   ```
   cd tmdb-cli-tool
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```

## Configuration

Before running the application, you need to set up your TMDB API key as an environment variable. You can do this by running the following command in your terminal:

### On Windows
```
set TMDB_API_KEY=your_api_key_here
```

### On macOS/Linux
```
export TMDB_API_KEY=your_api_key_here
```

Replace `your_api_key_here` with your actual TMDB API key.

## Usage

To run the application, use the following command:

```
go run main.go [movie_type]
```

Replace `[movie_type]` with one of the following options:
- `playing`
- `popular`
- `top`
- `upcoming`

### Example

To fetch popular movies, run:

```
go run main.go popular
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.