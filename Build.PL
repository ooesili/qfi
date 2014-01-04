use strict; use warnings;
use Builder;

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

my $build = Builder->new (
    module_name        => 'Qfi',
    dist_author        => 'Wesley Merkel <ooesili@gmail.com>',
    dist_abstract      => 'quickly edit commonly used files',
    license            => 'gpl',
    dist_version       => '0.1.0',
    create_makefile_pl => 'traditional',
    # add install action and paths for completion functions
    comp_files         => {'completion/bash.comp' => 'bashcomp/qfi',
                           'completion/zsh.comp'  => 'zshcomp/_qfi',
                          },
);

$build->add_build_element('comp');
$build->create_build_script;
