# gomesh

GoMeSH is a small library for parsing MeSH data files.

The code is probably best as an an example of how to parse large xml files (using
using only the go standard library). The functions in XMLFileParsers are based on
the design in this article (Parsing Huge XML Files With Go)[]http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/.

## Example
An example of using one of the parser functions would be as followed:
```go
f, err := os.Open("testdata/pa2017.xml")
if err != nil {
  panic(err)
}

pac, errc := gomesh.ParsePharmacologicalActionSet(f)

count := 0
for {
  select {
  case pa := <-pac:
    storeSomewhere(pa)
    count++
  case err := <-errc:
    log.Printf("Pharmacological Actions parsed: %d", count)
    if err == io.EOF {
      close(errc)
      return
    }
    panic(err)
    close(errc)
    return
  }
}
```

## Notes
I (crhntr) do not use this library in any production code nor do I commit to maintaining it.
