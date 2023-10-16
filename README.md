markdown

# Website Monitor in Go

This is a sample project that demonstrates how to create a website monitor in Go using goroutines and channels. The program checks the status of a list of URLs in parallel and records the results, including the status, URL, and response time.

## Features

- Parallel status monitoring of multiple websites.
- Real-time result logging.
- Control over the monitoring interval.
- Detection of changes in website status.

## How to Use

Follow the steps below to run the program:

1. Clone this repository:

```bash
git clone https://github.com/Gabriel-Jeronimo/go-url-monitoring
```

2. Navigate to the project directory:

```bash
cd go-url-monitoring
```

3. Edit the `main.go` file to add the URLs you want to monitor.

4. Run the program:

```bash
go run main.go
```

The program will start monitoring the websites and logging the results in real-time.

## Requirements

- Go (must be installed on your system).
- Internet connection to check the websites.

## How It Works

The program creates a goroutine for each URL in the list of websites to be monitored. Each goroutine periodically checks the website's status, measures the response time, and sends the results to a channel. A logging goroutine reads the results from the channel and logs them.

## Contribution

Feel free to contribute improvements, bug fixes, or additional features. Simply fork this repository, make the desired changes, and submit a pull request.

## License

This project is licensed under the MIT License. Please refer to the LICENSE file for details.
