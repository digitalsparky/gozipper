Zipping files via Go the 'complete' way

This allows you to add folders (recusive) or files while still maintaining dates and more importantly file permissions, even in windows.

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

# Like my stuff?

Would you like to buy me a coffee or send me a tip?
While it's not expected, I would really appreciate it.

[![Paypal](https://www.paypalobjects.com/webstatic/mktg/Logo/pp-logo-100px.png)](https://paypal.me/MattSpurrier) <a href="https://www.buymeacoffee.com/digitalsparky" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/white_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
