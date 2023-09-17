# CLI
![image](https://github.com/toodemhard/asciify/assets/100080774/0f447370-6916-4166-b15f-95109f462489)

## Install
Clone and cd into repository.
```sh
git clone https://github.com/toodemhard/asciify.git
cd asciify
```
Build and install to `/usr/local/bin/asciify`.
- Requires go version 1.21
```sh
./install
```

## Uninstall
```sh
rm /usr/local/bin/asciify
```
or use the uninstall script
```sh
cd asciify
./uninstall
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

# Library
## Install
```sh
go get -u github.com/toodemhard/asciify
```
