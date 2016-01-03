qfi [![Build Status](https://travis-ci.org/ooesili/qfi.svg?branch=master)](https://travis-ci.org/ooesili/qfi)
====================

### quickly access commonly used files and directories


Introduction
------------

`qfi` is a command line tool for UNIX-like systems, for quickly editing
commonly used files, and switching to commonly used directories.  This is
accomplished through the use of **targets**, which are short aliases for files
and directories.


Installation
------------

### Pre-built
The easiest way to install qfi is to download a pre-built binary from the
[releases page][3] for your OS and architecture.  Unzip the binary, and move it
to a directory in your `$PATH`.  If you can run `qfi` and see the usage message
, you're good to go.  After that, you'll probably want read below to setup
directory switching and tab completion for your shell.

### From source
If you have [Go][4] installed, running `go get github.com/ooesili/qfi` will
download and install program, although the version will not be set.  To do a
proper build from source cd into `$GOPATH/src/github.com/ooesili/qfi` and run
`make`. You'll then see an executable called `qfi` that directory.


Description
-----------

### Motivation
The reason I created `qfi` is that I found myself editing certain files on my
system quite often, and typing out the paths to these files over and over again
seemed awfully redundant.  `qfi` removes some of this redundancy by creating
shorter names for commonly used files, which I call **targets**.  I also wanted
to make `qfi` simple and easy to learn so that it wouldn't take more time to
learn than it would save.


### Targets
`qfi` maintains a list of **targets** as symbolic links inside of its
configuration directory.  To add a **target**, we call `qfi` with the `-a`
option, the name of the **target**, and the name of the file/directory:

```bash
$ qfi -a php /etc/php/php.ini
```

Now, whenever we want to open that file (or switch to that directory), we just
call `qfi` with the name of the **target**:

```bash
$ qfi php
```

which will open up the given **target** using the following rules:

 *  If the file is a directory, `cd` into it.

 *  If the file is writable by you, open it with the program specified in the
    `$EDITOR` environment variable, or `vi` if `$EDITOR` is not defined.

 *  If the file is not writable by you, open the file with `sudoedit`.

So in this case, the file will be opened with `sudoedit`.


### Configuration Directory
The configuration directory defaults to `$HOME/.config/qfi`, but
`$QFI_CONFIGDIR` will be used if it is defined.


### Options
 *  The `-a` option will add a **target** pointing to the specified
    file/directory.  If only one argument is given, the basename of the file
    wil be used as the target name:
    ```bash
    $ qfi -a [<target>] <filename>
    ```

 *  The `-m` option will move the destination of a **target** to a new
    file/directory:
    ```bash
    $ qfi -m <target> <filename>
    ```

 *  The `-d` option will delete the listed **targets**:
    ```bash
    $ qfi -d <target1> [<target2> [...]]
    ```

 *  The `-r` option will rename a **target**:
    ```bash
    $ qfi -r <target> <newname>
    ```

 *  The `-l` option will list all **targets**, or where the specified
    **target** points to, if one was given:
    ```bash
    $ qfi -l [target]
    ```

 *  The `-s` option will show the status and destination of each **target**.
    ```bash
    $ qfi -s
    ```
    Listings are in the following format:
    ```
    target *> /path/to/destination
    ```
    where the line's color and `*>' are one of the following:

    Symbol | Meaning           | Color
    ------ | ----------------- | ------
    ->     | writable file     | green
    />     | directory         | blue
    #>     | non-writable file | yellow
    ?>     | unknown file      | purple

    Red lines indicate a non-existent or non-accessible file, or a non chdir-able
    directory.  These files may still be used, but you may encounter problems
    saving them if their parent directories do not exist.

 *  The `--shell` prints wrapper and completion scripts for the supported
    shells:
    ```bash
    $ qfi --shell (zsh|fish|bash) (comp|wrapper)
    ```

 *  The `--help` and `--version` options will display usage and version
    information, respectively.


### Tab Completion
`qfi` comes with bash, zsh, and fish tab completion support for **targets** and
options.  For bash completion to work, [bash-completion][2] must be installed
on the system.

#### bash
Put this line into your `~/.bashrc`:
```bash
eval "$(qfi --shell bash comp)"
```

#### fish
Put this line into your `~/.config/fish/config.fish`:
```fish
eval (qfi --shell fish comp)
```

#### zsh
Installing zsh is a little more complicated.  Write the completion script to a
file called `_qfi`, and put that file inside of a directory in your `$FPATH`.
For example:
```zsh
$ qfi --shell zsh comp > /usr/share/zsh/site-functions/_qfi
```


### Directory Switching
`qfi` allows targets to point to directories, in which case your shell will
`cd` into the directory.  This feature was largely inspired by [autojump][1],
which is a very cool piece of software with similar functionality. The
following commands will setup directory switching for your shell:

bash:
```bash
eval "$(qfi --shell bash wrapper)"
```

fish:
```fish
eval (qfi --shell fish wrapper)
```

zsh:
```zsh
eval "$(qfi --shell zsh wrapper)"
```


[1]: https://github.com/joelthelion/autojump
[2]: http://bash-completion.alioth.debian.org/
[3]: https://github.com/ooesili/qfi/releases
[4]: https://golang.org
