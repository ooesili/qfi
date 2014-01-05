#!/usr/bin/env perl

# qfi - quicky edit commonly used files
#
# Copyright (C) 2014 Wesley Merkel <ooesili@gmail.com>
#
# This file is part of qfi
#
# qfi is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# qfi is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with qfi.  If not, see <http://www.gnu.org/licenses/>.

package Qfi;

use strict; use warnings;
use File::Path 'make_path';
use File::Basename;
use File::Spec;

BEGIN {
    require Exporter;

    our $VERSION = 0.1.0;
    our @ISA = qw(Exporter);
    #our @EXPORT = qw(add delete edit list move rename conf_dir);
    #our @EXPORT_OK = qw();
}

our $conf_dir;

# first arg is name, second is the file; symlink the file to the name
sub add {
    # check for slashes in target name
    die "$0: illegal target name: `$_[0]'\n" if ($_[0] =~ m|/|);
    my $link = File::Spec->catfile($conf_dir, $_[0]);
    # symlinks act funny with relative paths
    my $file = File::Spec->rel2abs($_[1]);
    symlink $file, $link or die "$0: error creating link for `$_[0]': $!\n"
}

# first arg is name; delete associated link
sub delete {
    my $link = File::Spec->catfile($conf_dir, $_[0]);
    unlink $link or die "$0: error removing target `$_[0]': $!\n";
}

# first arg is name; edit file pointed to by the link assocated with the name
sub edit {
    my ($env_editor, $editor, $link, $file, $uid);
    # follow link
    $link = File::Spec->catfile($conf_dir, $_[0]);
    if (-l $link) { $file = readlink $link; }
    else { die "$0: cannot find target `$_[0]': $!\n"; }
    # see who owns it
    if (-d $file) {
        die "$0: directory-switching not enabled, see `man 1 qfi' to fix\n";
    }
    elsif (-f $file) { $uid = (stat $file)[4]; }
    else { die "$0: cannot stat `$file': $!\n"; }
    # pick editor
    if ($uid == 0) {
        $editor = "sudoedit";
    }
    elsif (defined($env_editor = $ENV{'EDITOR'})) {
        $editor = $env_editor;
    }
    else {
        $editor = "vi";
    }
    # run editor
    exec "$editor $file";
}

# list all defined targets, or the dest of a single link, if specified
sub list {
    # print destination if an argument was given
    if (defined $_[0]) {
        my ($file, $link);
        $file = File::Spec->catfile($conf_dir, $_[0]);
        if (-l $file) { $link = readlink $file; }
        else { die "$0: cannot find target `$_[0]': $!\n"; }
        print "$link\n";
    }
    # otherwise print all of the targets
    else {
        my @files = glob File::Spec->catfile($conf_dir,"*");
        for (sort @files) {
            my $target = fileparse($_);
            printf "$target\n";
        }
    }
}

# rename a target
sub rename {
    chdir $conf_dir;
    die "$0: cannot find target `$_[0]': $!'\n" if (! -l $_[0]);
    # check for slashes in target name
    die "$0: illegal target name: `$_[0]'\n" if ($_[0] =~ m|/|);
    rename $_[0], $_[1];
}

# first arg is target, second is filename; move a target's destination
sub move {
    &delete($_[0]);     # delete old target
    &add($_[0], $_[1]); # add target with new destination
}

# set and return appropriate configuration directory
# create it if it does not exists
{
    my ($xdg_dir, $home_dir);
    # first see if $XDG_CONFIG_HOME is defined
    if (defined($xdg_dir = $ENV{'XDG_CONFIG_HOME'})) {
        $conf_dir = File::Spec->catdir($xdg_dir, "qfi");
    }
    # then try to build conf_dir path from $HOME
    elsif (defined($home_dir = $ENV{'HOME'})) {
        $conf_dir = File::Spec->catdir($home_dir, ".config", "qfi");
    }
    else {
        warn "$0: \$HOME not set\n";
        exit(1);
    }
    # create configration directory if it doesn't exist
    unless (-d $conf_dir) {
        make_path($conf_dir) or die "$0: cannot create $conf_dir: $!\n";
    }
}

# report a successful import
1;
