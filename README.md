# Jnana

Alfred workflow to index eBook (EPUB & PDF) bookmarks. Filter bookmarks of current EPUB (calibre 4 ebook-viewer) or PDF (Preview, Skim, and Acrobat) and go to page. Search index of all previously opened eBooks and open any of them to proper section.

Alfred 3.

## Usage

### Workflow commands

- `.jna` - search and go to bookmarks for currently opened ePub or PDF.
- `.jnaall` - bookmarks of any type

## Building

## Build

`go build -o jnana --tags "stat4 foreign_keys vacuum_incr introspect fts5" *.go`