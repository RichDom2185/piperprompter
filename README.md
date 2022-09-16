# piperprompter

> A simple command-line utility to redirect STDOUT to a user-specified file.

I made this because I got tired of having to manually edit the output redirect `>` of a command each time. E.g.:

```bash
fbgrab - > 1.png
fbgrab - > 2.png
fbgrab - > 3.png
fbgrab - > 4.png
# ...
```

Using piperpromter, I can just keep repeating the same command by hitting the arrow key:

```bash
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

# Installing

```bash
go install github.com/RichDom2185/piperprompter@latest
```
