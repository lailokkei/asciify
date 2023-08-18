# CLI
![image](https://github.com/toodemhard/asciify/assets/100080774/149cc808-d321-45e6-830c-31dd20b5f8b6)

## Installation
Manual
```sh
git clone https://github.com/toodemhard/asciify.git
cd asciify
scripts/build-cli

# run
bin/asciify -h
```

## Usage
Show help page
```sh
asciify -h
```
```
Usage:
  asciify [OPTIONS]

Application Options:
  -f, --file=         Image file path to source
  -i, --invert        Invert the values of the image
  -c, --charset=      Set of characters to use in output (default: simple)
  -s, --scale=        Width of output in number of characters (default: 20)
  -m, --sampleMethod= Method of converting grid of pixels to character (default: mid)

Help Options:
  -h, --help          Show this help message
```

# Web
Start server
```sh
go run cmd/server/main.go
```
