# go-spider

This is a simple but very well done crawler project, implemented entirely in Go.

The input is a basic seed file, specifically a URL, from which the crawler starts crawling, parses the page, extracts the specified web elements (e.g. HTML suffixes) and saves the results to a local file. These crawling processes are ultimately controlled by the crawl depth parameter in the configuration file, which determines the crawler's crawl depth to the target.

This crawler project is essentially a distributed project with the following specific features, and arguably functions.

1. support configuration of the maximum depth of crawling
2. supports configuring the crawl interval and timeout time.
3. support multi-threaded execution (multiple goroutines)
4. support for configuring multiple crawl sources (starting point), support for configuring the output path of crawl results.


## Prerequisites
You'll need to install golang locally first

## Getting Started
```bash
# clone repo
git clone https://github.com/vinsec/go-spider.git

# cd to project directory
cd go-spider

# create the required directories
mkdir -p {output,bin}

# build the project
cd src && go build -o ../bin/go-spider .

# change the seed file (the starting site for crawling) to the site you want to crawl
cd - && vim data/seed

# run the spider
cd bin && ./go-spider

```

The default configuration file is conf/spider.conf, and you can change these parameters at will before running

When you see "request queue nil, all sub spiders are idle, Spider exit.", the crawling process will end and the crawling results will be stored in the output directory


## Contributing

Please read [CONTRIBUTING.md](#) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors 

* **vinsec** - *Go-Spider* - [vinsec](https://github.com/vinsec)

## License

This project uses [Apache License 2.0](LICENSE.md) 
