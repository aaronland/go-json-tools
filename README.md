# go-json-tools

There are many tools for working with JSON. These ones are ours.

## Deprecated

This package has been deprecated. These tools have been moved in to [aaronland/go-tools](https://github.com/aaronland/go-tools).

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

### jv

Ensure JSON is valid (where "valid" means it can be parsed)

```
$> ./bin/jv ~/sfomuseum/go-sfomuseum-media/config/*.json
```

