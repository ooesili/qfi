qfi [![Build Status](https://travis-ci.org/ooesili/qfi.svg?branch=master)](https://travis-ci.org/ooesili/qfi)
====================

### quickly edit commonly used files


Introduction
--------------------

`qfi` is a command line tool for UNIX-like systems, for quickly editing
commonly used files, and switching to commonly used directories.  This is
accomplished through the use of **targets**, which are short aliases for files
and directories.  For a more terse explanation of how `qfi` works, consult the
manual page.


Description
--------------------

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

 *  If the file does not exist, use the longest existing part of it's path, and
    the rules above, to determine which program to use.

So in this case, the file will be opened with `sudoedit`.

### Configuration Directory
The configuration directory defaults to `$HOME/.config/qfi`, but
`$XDG_CONFIG_HOME/qfi` will be used if `$XDG_CONFIG_HOME` is defined.

### Options
 *  The `-a` option will add a **target** pointing to the specified
    file/directory:
    ```bash
    $ qfi -a <target> <filename>
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
    target ?> /path/to/destination
    ```
    where the line's color and `?>' are one of the following:

    Symbol | Meaning           | Color
    ------ | ----------------- | ------
    ->     | writable file     | green
    />     | directory         | blue
    #>     | non-writable file | orange

    Red lines indicate a non-existent or non-accessible file, or a non chdir-able
    directory.  These files may still be used, but you may encounter problems
    saving them if their parent directories do not exist.


 *  The `--help` and `--version` options will display usage and version
    information, respectively.

### Tab-Completion
`qfi` comes with bash, zsh, and fish tab-completion support for **targets** and
options.  For bash completion to work, additional software must be installed,
as specified below.

### Directory-Switching
`qfi` allows targets to point to directories, in which case your shell will
`cd` into the directory.  To properly enable this feature, view the
instructions in the `INSTALL` file, or manual page.  This feature was largely
inspired by [autojump][1], which is a very cool piece of software with similar
functionality.


Dependencies
--------------------

#### Perl
`qfi` depends on Perl, as it is written the language.

#### sudo
`qfi` depends on `sudoedit` to open files owned by other users.  In the future,
I *could* make this an optional dependency and just open the file with the
normal editor if `sudoedit` couldn't be found.  Tell me if you would like this
feature.

#### bash-completion
This one is optional.  `qfi` can provide bash tab-completion support for
**targets** and options if the [bash-completion][2] software is installed.

Distribution Packages
---------------------

 *  Arch Linux: [qfi][3] (AUR)


[1]: https://github.com/joelthelion/autojump
[2]: http://bash-completion.alioth.debian.org/
[3]: https://aur.archlinux.org/packages/qfi/
