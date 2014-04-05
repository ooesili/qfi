function qfi
    set -l dest
    # normal run by default
    set -l normal_run true
    # find path of real qfi command
    set -l qfi_bin (/usr/bin/env which qfi)

    # see if the only argument points to a directory
    if [ (count $argv) -eq 1 ]
        if echo $argv[1] | grep -qv '^-'
            # cd if the target destination is a directory
            set dest (eval $qfi_bin -l $argv[1] ^&-)
            if test -d "$dest"
                # change directories and don't run normally
                set normal_run false
                cd $dest
            end
        end
    end

    # if none of that worked, run qfi with the supplied arguments
    if [ $normal_run = 'true' ]
        eval $qfi_bin $argv
    end
end
