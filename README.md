# SQLite to JSON
This program allows you to run queries on a SQLite database and output each row as JSON. This is a very basic program, could be modified for other purposes.

### Usage
```console
./stj mydatabase.db "SELECT * FROM mytable" > outputfile
```

### Goal
The goal of this program is to be as simple as possible. It also uses no CGO SQLite to enable cross compilation. See [supported platforms](https://pkg.go.dev/modernc.org/sqlite#hdr-Supported_platforms_and_architectures). **This isn't a replacement for sqlite3**.

#### Sources
I used/plagiarized parts of [@SchumacherFM's code](https://gist.github.com/SchumacherFM/69a167bec7dea644a20e)