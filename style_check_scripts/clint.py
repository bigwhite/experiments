#!/usr/bin/env python 

import getopt
import os
import re
import sys

USAGE = """
 Usage: clint.py [--verbose=#] [--output=vs7] [--filter=-x,+y,...]
                    [--counting=total|toplevel|detailed]
                            <file> [file] ...


"""

def print_usage(message):
    """Prints a brief usage string, optionally with an error message.

    Args:
    message: The optional error message.
    """

    sys.stderr.write(USAGE)
    if message:
        sys.stderr.write('\nFATAL ERROR: %s\n' % message)


def parse_arguments(args):
    """ Parse the command line arguments
    
    Args:
        args: The command line arguments:

    Returns:
        The list of filenames to lint.
    """

    try:
        (opts, filenames) = getopt.getopt(args, '', ['help', 'output=', 
                                                     'verbose=','filter='])
    except getopt.GetoptError:
        print_usage('Invalid arguments.')
        sys.exit(1)

    if not filenames:
        print_usage('No files were specified')


def main():
    filenames = parse_arguments(sys.argv[1:])

if __name__ ==  '__main__':
    main()
