# imgurCrawler
URL bruteforcer for imgur.com

# Install

Download and extract release from https://github.com/The-Eye-Team/imgurCrawler/releases

# Usage

```
./imgurCrawler -h

usage: imgurCrawler [-h|--help] [-j|--concurrency <integer>]

                    Crawl imgur.com IDs

Arguments:

  -h  --help         Print help information
  -j  --concurrency  Concurrency. Default: 4
```

# Example

How to crawl imgur using 8 concurrent workers:

`./imgurCrawler -j 8`

