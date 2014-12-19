qfi() {
    local dest
    # normal run by default
    local normal_run=true
    # find path of real qfi command
    local qfi_bin=$(/usr/bin/env which qfi)

    # see if the only argument points to a directory
    if [[ $# -eq 1 && $1 != -* ]]; then
        # cd if the target destination is a directory
        dest=$($qfi_bin -l $1 2>/dev/null)
        if [[ -d "$dest" ]]; then
            # change directories and don't run normally
            normal_run=false
            cd "$dest"
        fi
    fi

    # if none of that worked, run qfi with the supplied arguments
    if [[ $normal_run == true ]]; then
        $qfi_bin $@
    fi
}

# vim: ft=sh:
