# Go-Dict

Simple dictionary in Golang with persisted data to [BadgerDB](https://dgraph.io/docs/badger/get-started/).

## Usage

Build dictionary

```
go build -o dict
```

Add a word to the dictionary

```
./dict -action add <word> <definition>
```

Remove a word from the dictionary

```
./dict -action remove <word>
```

Get the definition from a word in the dictionary

```
./dict -action define <word>
```

List the words from the dictionary

```
./dict -action list
```
