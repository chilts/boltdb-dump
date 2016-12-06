# boltdb-dump #

Command to dump a human-readable BoltDB to stdout. This works with buckets at any level, not just the top-level.

Note: currently this will only be readable with string keys and values such as JSON. If you're using MsgPack or
Protocol Buffers for your values, then this program won't do what you want (yet).

## Install ##

```sh
go get -u github.com/chilts/boltdb-dump
```

## Usage ##

There are (currently) no options, nor anythin fancy. Just pass the db file you want to dump:

```sh
boltdb-dump database.db
```

An example of a blog site with users and domains:

```sh
[users]
  chilts
    {"UserName":"chilts","Email":"andychilton@gmail.com"}
[domains]
  [chilts.org]
    [authors]
      andrew-chilton
    [posts]
      first-post
        {"PostName":"first-post","Title":"First Post","Content":"Hello, World!"}
  [blog.appsattic.com]
    [authors]
      andrew-chilton
    [posts]
```

## Author ##

[Andrew Chilton](https://chilts.org/) - [@andychilton](https://twitter.com/andychilton).

## License ##

[MIT](https://chilts.mit-license.org/2016/).

(Ends)
