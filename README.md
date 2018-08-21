# go-json-tools

There are many tools for working with JSON. These ones are ours.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tools

### jp

Make JSON pretty.

```
./bin/jp -h
Usage of ./bin/jp:
  -rewrite
    	Rewrite the JSON file in place.
```

For example:

```
./bin/jp -rewrite example.json
```

Or to read from `STDIN`:

```
cat example.json | ./bin/jp -
```