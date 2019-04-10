# Jnana

Alfred workflow to index eBook (EPUB & PDF) bookmarks. Filter bookmarks of current EPUB (calibre ebook-viewer) or PDF (Preview, Skim, and Acrobat) and go to page. Search index of all previously opened ebooks and open any of them to proper section.

Alfred 3.

## Usage

### Workflow commands

- `.jna` - search and go to bookmarks for currently opened ePub or PDF.
- `.jnaall` - bookmarks of any type

## Building

### Compile-time options

`export CGO_CFLAGS="-DSQLITE_THREADSAFE=0 -DSQLITE_DEFAULT_MEMSTATUS=0 -DSQLITE_DEFAULT_WAL_SYNCHRONOUS=1 -DSQLITE_MAX_EXPR_DEPTH=0 -DSQLITE_OMIT_DEPRECATED -DSQLITE_OMIT_PROGRESS_CALLBACK -DSQLITE_OMIT_SHARED_CACHE -DSQLITE_USE_ALLOCA -DSQLITE_TEMP_STORE=3 -DSQLITE_ENABLE_ATOMIC_WRITE -DSQLITE_ENABLE_BATCH_ATOMIC_WRITE -DSQLITE_ENABLE_NULL_TRIM -DSQLITE_DEFAULT_FOREIGN_KEYS=1 -DSQLITE_DEFAULT_SYNCHRONOUS=0 -DSQLITE_DEFAULT_AUTOVACUUM=2 -DSQLITE_DEFAULT_PAGE_SIZE=4096"`

## Build
`go build -o jnana --tags "stat4 foreign_keys vacuum_incr introspect fts5" *.go`