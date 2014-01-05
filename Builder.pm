#!/usr/bin/env perl

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

package Builder;

use strict; use warnings;
use Module::Build;
use ExtUtils::Install;
use File::Spec;
use IO::Compress::Gzip qw(gzip $GzipError);
use parent 'Module::Build';

# overridden to gzip the man files
sub ACTION_docs {
    my $self = shift;
    my ($dh, $file);
    # generate documentation
    $self->SUPER::ACTION_docs;
    # change into bindoc directory
    chdir File::Spec->catdir('blib', 'bindoc');
    # gzip every man file
    opendir($dh, File::Spec->curdir);
    while ($file = readdir($dh)) {
        # gzip if it is a perl man file
        if ($file =~ /\.\dp$/) {
            # compress file and remove old version
            gzip $file => "$file.gz"
                or die "$0: gzip failed: $GzipError\n";
            unlink $file or die "$0: could not delete `$file': $!\n";
        }
    }
    # go back the base directory
    chdir $self->base_dir;
};

# set relative installation paths for util functions
sub ACTION_build {
    my $self = shift;
    # get installation prefix
    my $dirs = $self->installdirs();
    my $prefix = $self->_default_install_paths->{'original_prefix'}{$dirs};
    # set path based on whether install_base is specified
    my $base = $self->install_base();
    if (defined $base) {
        # set relative paths
        $self->install_base_relpaths(
            'bashcomp' => 'share/bash-completion/completions');
        $self->install_base_relpaths(
            'zshcomp'  => 'share/zsh/site-functions');
        $self->install_base_relpaths(
            'share'    => 'share/qfi');
    }
    else {
        # set absolute paths
        $self->install_path(
            'bashcomp' => "$prefix/share/bash-completion/completions");
        $self->install_path(
            'zshcomp'  => "$prefix/share/zsh/site-functions");
        $self->install_path(
            'share'    => "$prefix/share/qfi");
    }
    # run super function
    $self->SUPER::ACTION_build;
}

1;

# uninstallation action
sub ACTION_uninstall {
    my $self = shift;
    # get path of .packlist file
    my $packlist = $self->install_map->{'write'};
    # uninstall package, verbosely, without faking
    &ExtUtils::Install::uninstall($packlist, 1, 0);
}
