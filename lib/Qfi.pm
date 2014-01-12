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

    our $VERSION = "v0.2.0";
    our @ISA = qw(Exporter);
    #our @EXPORT = qw(add delete edit list move rename conf_dir);
    #our @EXPORT_OK = qw();
}

##### FIGURE OUT/CREATE $conf_dir #####
my $conf_dir;

sub init {
    my ($xdg_dir, $home_dir);
    # first see if $XDG_CONFIG_HOME is defined
    if (defined($xdg_dir = $ENV{'XDG_CONFIG_HOME'})) {
        $conf_dir = File::Spec->catdir($xdg_dir, "qfi");
    }
    # then try to build conf_dir path from $HOME
    elsif (defined($home_dir = $ENV{'HOME'})) {
        $conf_dir = File::Spec->catdir($home_dir, ".config", "qfi");
    }
    else { return &fwarn("\$HOME not set"); }
    # create configration directory if it doesn't exist
    unless (-d $conf_dir) {
        # try to create directory
        make_path($conf_dir, {error => \my $err});
        # loop through error messages
        for my $diag (@$err) {
            my ($file, $mesg) = %$diag;
            # general error
            if ($file eq '') { &fwarn("error creating `$conf_dir': $mesg"); }
            # error creating specific file
            else             { &fwarn("error creating `$file': $mesg"); }
        }
        # throw error if we couldn't create $conf_dir
        return 0 if @$err;
    }
    else { return 1; }
}


##### UTILITY FUNCTIONS #####
# construct the path to a target link
sub target_link { File::Spec->catfile($conf_dir, $_[0]); }

# print mesg to stderr, returns false
sub fwarn { warn "$0: $_[0]\n"; 0; }

# gets the destination of a target
sub get_target_dest {
    my $file;
    my $target = shift;
    my $link = &target_link($target);
    # make sure target exists
    if (not -l $link) { &fwarn("target `$target' does not exist"); }
    # print error and return false if operation failed
    elsif ($file = readlink $link) { return $file; }
    else { &fwarn("cannot follow link for `$target': $!"); }
}

# checks a target name for invalid characters
sub invalid_target_name {
    my $target = shift;
    if ($target =~ m|/|) { &fwarn("illegal target name: `$target'"); 1; }
    else { return 0; }
}


##### MAIN FUNCTIONS #####
# first arg is name, second is the file; symlink the file to the name
sub add {
    my $target = shift;
    my $dest   = shift;
    # check for valid target name
    return 0 if &invalid_target_name($target);
    my $link = &target_link($target);
    # make sure file does not exist
    if (-l $link) { return &fwarn("target `$target' exists"); }
    # symlinks act funny with relative paths
    my $abs_dest = File::Spec->rel2abs($dest);
    # print error and return false if operation failed
    if (symlink $abs_dest, $link) { return 1; }
    else { &fwarn("error creating link for `$target': $!"); }
}

# delete a target
sub delete {
    my $target = shift;
    my $link = &target_link($target);
    # make sure target exists
    if (not -l $link) { &fwarn("target `$target' does not exist"); }
    # print error and return false if operation failed
    elsif (unlink $link) { return 1; }
    else { &fwarn("error removing target `$_[0]': $!"); }
}

# first arg is target; edit file pointed to by the target
sub edit {
    my $target = shift;
    my ($env_editor, $editor, $file, $uid);
    # get the destination of the target
    $file = &get_target_dest($target) or return 0;
    # the wrapper script should catch this condition, so it must not be loaded
    if (-d $file) {
        return &fwarn( "directory-switching not enabled, see `man 1 qfi'");
    }
    # if it's a file and we can see it
    elsif (-f $file) { $uid = (stat $file)[4]; }
    # if we can't see the file, use sudoedit
    elsif ($!{'EACCES'}) { $uid = 0; }
    # if we can access the location, but the file is not found, use the UID of
    # the parent directory
    else {
        my @dirs = (File::Spec->splitpath($file))[1];
        $uid = (stat File::Spec->catdir(@dirs))[4];
        # warn and return false if stat failed
        if (not defined $uid) {
            return &fwarn("cannot stat parent dir of `$file'");
        }
    }
    # pick editor
    if ($uid != $<)                      { $editor = "sudoedit"; }
    elsif ($env_editor = $ENV{'EDITOR'}) { $editor = $env_editor; }
    else                                 { $editor = "vi"; }
    # run editor
    exec "$editor $file";
}

# list all defined targets, or the dest of a single link, if specified
sub list {
    my $target = shift;
    # print destination if an argument was given
    if (defined $target) {
        my $dest = &get_target_dest($target) or return 0;
        print "$dest\n";
    }
    # otherwise print all of the targets
    else {
        my @files = glob &target_link("*");
        for (sort @files) {
            $target = fileparse($_);
            printf "$target\n";
        }
    }
    # if we got here, we must have succeeded
    return 1;
}

# rename a target
sub rename {
    my $target   = shift;
    my $new_name = shift;
    # check for valid target name
    return 0 if &invalid_target_name($target);
    # construct paths to links
    my $old_link = &target_link($target);
    my $new_link = &target_link($new_name);
    # make sure the old target exists and the new one doesn't
    if (not -l $old_link) { &fwarn("target `$target' does not exist"); }
    elsif (-l $new_link) { &fwarn("target `$new_name' exists"); }
    # print error and return false if operation failed
    elsif (rename $old_link, $new_link) { return 1; }
    else { &fwarn("error renaming `$target"); }
}

# first arg is target, second is filename; move a target's destination
sub move {
    my $target = shift;
    my $dest   = shift;
    # delete old and add new
    &delete($target) and &add($target, $dest);
}

1; # report a successful import
