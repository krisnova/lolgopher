# lolgopher
Rainbow Writer implementation for Golang

Based on https://github.com/busyloop/lolcat.

## TODO:

* [ ] Detect tty and turn off color there
* [ ] Have helpers for LOLing directly to stdout or other writers.  Similar to log packages
* [ ] Handle terminals that aren't true color.  Do some detection if we can
* [ ] Exapand tabs and strip pre-existing ansi codes
* [ ] Support inverse (set background, not text)
* [ ] Support animation
* [ ] Implement lolcat