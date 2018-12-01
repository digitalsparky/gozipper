Zipping files via Go the 'complete' way

Usage:

```
zipFile, err := zip.NewZipFile("zip-file-path.zip")
if err != nil {
    return err
}

zipFile.AddFile("file") // run for each file
zipFile.AddFolder("folder") // add a folder

zipFile.Write() // Write the zip file
```