# fend
Check for Newline at **F**ile **End**

# Why might you want to do this?

This StackOverflow post captures the *why?* better than I could:

https://stackoverflow.com/questions/729692/why-should-text-files-end-with-a-newline

It will mean no more GitHub warnings for 'No newline at EOF' on Pull Request.

TODO: Image

By enforcing this check using the GitHub Action you can automate a basic Standard in your project.

Consistent is clean. Clean is good. Don't overthink it. :ok_hand:

# TODO
## Docs
* Tests doc.
* Readme docs. release.
* GH Action release
## Code
* Finish basic testing.
* Better algorithm for different categories.
* Format strings everywhere
* Format output inc. errors and edge cases.
* Define and use data structures for ScanConfig and ScanResults.

# Usage
You can use the 
```
go get github.com/njgibbon/fend
# ensure binary can be foud in $PATH
cd <dir-to-scan>
fend
```
# Configuration
TODO

# GitHub Action
TODO

# Meta
Project used as a vehicle to help learn about basic tool development using GoLang and also to explore GitHub Actions from a development point of view as I have been having a really positive experience with them as a User.

# Similar tools
TODO: write nice things about the other tool I found.
