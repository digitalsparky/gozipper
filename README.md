Zipping files via Go the 'complete' way

Usage:

```
zip, err := NewZipFile("zip-file-path.zip")
if err != nil {
    return err
}

zip.AddFile("file") // run for each file

zip.Write() // Write the zip file
```