#!/bin/bash

set -uo pipefail

PATTERN="${PATTERN:-validator}"
DRY_RUN="${DRY_RUN:-false}"
VERBOSE="${VERBOSE:-false}"

readonly RED='\033[1;31m'
readonly GRN='\033[0;32m'
readonly YEL='\033[1;33m'
readonly BLU='\033[0;34m'
readonly PUR='\033[0;35m'
readonly NOC='\033[0m'

log_info() {
    echo -e "${BLU}[INFO]${NOC} $1"
}

log_success() {
    echo -e "${GRN}[SUCCESS]${NOC} $1"
}

log_warning() {
    echo -e "${YEL}[WARNING]${NOC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NOC} $1" >&2
}

log_verbose() {
    [[ "$VERBOSE" == "true" ]] && echo -e "${PUR}[VERBOSE]${NOC} $1"
}

show_help() {
    cat <<EOF
Post Validation Cleanup Script

USAGE:
    $0 [OPTIONS]

OPTIONS:
    -p, --pattern PATTERN   Container name pattern to match (default: validator)
    -d, --dry-run           Show what would be removed without actually removing
    -v, --verbose           Enable verbose output
    -h, --help              Show this help message

ENVIRONMENT VARIABLES:
    PATTERN                 Container name pattern (default: validator)
    DRY_RUN                 Set to 'true' for dry run mode
    VERBOSE                 Set to 'true' for verbose output

EXAMPLES:
    $0                      # Remove exited validator containers
    $0 -p "worker"          # Remove exited containers matching "worker"
    $0 --dry-run            # Show what would be removed
    VERBOSE=true $0         # Run with verbose output
EOF
}

parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
        -p | --pattern)
            PATTERN="$2"
            shift 2
            ;;
        -d | --dry-run)
            DRY_RUN="true"
            shift
            ;;
        -v | --verbose)
            VERBOSE="true"
            shift
            ;;
        -h | --help)
            show_help
            exit 0
            ;;
        *)
            log_error "Unknown option: $1"
            show_help
            exit 1
            ;;
        esac
    done
}

check_docker() {
    if ! command -v docker &>/dev/null; then
        log_error "Docker is not installed or not in PATH."
        exit 1
    fi

    if ! docker info &>/dev/null; then
        log_error "Cannot connect to Docker daemon. Is Docker running?"
        exit 1
    fi
}

# Get container information with error handling
get_container_info() {
    local container_id="$1"
    local field="$2"

    # Use a subshell to prevent the main script from exiting on error
    local result
    result=$(docker inspect --format="{{$field}}" "$container_id" 2>/dev/null) || {
        log_verbose "Failed to get $field for container $container_id"
        return 1
    }

    echo "$result"
}

# Remove container with error handling
remove_container() {
    local container_id="$1"

    if [[ "$DRY_RUN" == "true" ]]; then
        log_info "DRY RUN: Would remove container $container_id"
        return 0
    fi

    if docker rm "$container_id" &>/dev/null; then
        log_verbose "Successfully removed container $container_id"
        return 0
    else
        log_error "Failed to remove container $container_id"
        return 1
    fi
}

cleanup_containers() {
    log_info "Searching for exited containers matching pattern: '$PATTERN'"

    # Get all exited containers that match the name pattern
    local containers
    containers=$(docker ps -a --filter "status=exited" --format "{{.ID}} {{.Names}}" | grep "$PATTERN" || true)

    if [[ -z "$containers" ]]; then
        log_warning "No exited containers found matching pattern '$PATTERN'"
        return 0
    fi

    log_info "Found exited containers matching pattern:"
    echo

    # Counters for summary
    local total_found=0
    local total_removed=0
    local total_kept=0
    declare -A exit_code_counts=()
    declare -a removal_failures=()

    # Process each matching exited container
    while IFS= read -r line; do
        [[ -z "$line" ]] && continue

        local container_id container_name exit_code status
        container_id=$(echo "$line" | awk '{print $1}')
        container_name=$(echo "$line" | awk '{print $2}')

        log_verbose "Processing: $container_id ($container_name)"

        ((total_found++))

        # Get container exit code and status with error handling
        if ! exit_code=$(get_container_info "$container_id" ".State.ExitCode"); then
            log_error "Skipping container $container_id ($container_name) - inspection failed"
            continue
        fi

        if ! status=$(get_container_info "$container_id" ".State.Status"); then
            log_error "Skipping container $container_id ($container_name) - inspection failed"
            continue
        fi

        printf "%-12s %-30s %-20s" "$container_id" "$container_name" "$status (exit $exit_code)"

        if [[ "$exit_code" == "0" ]]; then
            if remove_container "$container_id"; then
                echo -e " → ${GRN}Removed${NOC}"
                ((total_removed++))
            else
                echo -e " → ${RED}Failed to remove${NOC}"
                removal_failures+=("$container_id ($container_name)")
            fi
        else
            echo -e " → ${YEL}Kept (non-zero exit)${NOC}"
            ((total_kept++))
            exit_code_counts["$exit_code"]=$((exit_code_counts["$exit_code"] + 1))
        fi
    done <<<"$containers"

    echo
    log_success "Cleanup complete!"
    echo

    log_info "Summary:"
    printf "    Total containers found: %d\n" "$total_found"
    printf "    Containers removed: %d\n" "$total_removed"
    printf "    Containers kept: %d\n" "$total_kept"

    if [[ ${#removal_failures[@]} -gt 0 ]]; then
        echo
        log_warning "Failed to remove ${#removal_failures[@]} container(s):"
        for failure in "${removal_failures[@]}"; do
            echo "   - $failure"
        done
    fi

    if [[ ${#exit_code_counts[@]} -gt 0 ]]; then
        echo
        log_info "Containers kept by exit code:"
        printf "    %-9s | %s\n" "Exit Code" "Count"
        printf "    %s\n" "----------|------"
        for code in $(printf '%s\n' "${!exit_code_counts[@]}" | sort -n); do
            printf "    %-9s | %d\n" "$code" "${exit_code_counts[$code]}"
        done
    else
        echo
        log_success "All validator containers exited cleanly (exit code 0)."
    fi

    if [[ "$DRY_RUN" == "true" ]]; then
        echo
        log_info "This was a dry run. Use without --dry-run to actually remove containers."
    fi
}

main() {
    parse_args "$@"

    log_verbose "Configuration:"
    log_verbose "   Pattern: $PATTERN"
    log_verbose "   Dry run: $DRY_RUN"
    log_verbose "   Verbose: $VERBOSE"

    check_docker
    cleanup_containers
}

# Only run main if script is executed directly (not sourced)
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
