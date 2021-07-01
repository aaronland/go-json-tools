# go-json-tools

There are many tools for working with JSON. These ones are ours.

## Documentation

Documentation is incomplete.

## Tools

### jf

Flatten JSON.

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