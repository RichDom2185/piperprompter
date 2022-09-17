# piperprompter

> A simple command-line utility to redirect STDOUT to a user-specified file.

I made this because I got tired of having to manually edit the output redirect `>` of a command each time. E.g.:

```
$ fbgrab - > 1.png
$ fbgrab - > 2.png
$ fbgrab - > 3.png
$ fbgrab - > 4.png
# ...
```

Using piperpromter, I can just keep repeating the same command by hitting the up arrow key in the terminal:

```
$ piperpromter fbgrab -
Enter filename: 1.png
Saved to 1.png!
$ piperpromter fbgrab -
Enter filename: 2.png
Saved to 2.png!
$ piperpromter fbgrab -
Enter filename: 3.png
Saved to 3.png!
# ...
```

This may or may not be useful in scripts, but mainly useful as a time-saver when one needs to repeatedly run the same command again and again.

# Installing

```bash
go install github.com/RichDom2185/piperprompter@latest
```
