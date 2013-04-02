#! /usr/bin/env python
# -*- coding: utf-8 -*-

#
# A simple-minded style checker for C code.
# This only catches the most obvious style mistakes, and occasionally
# flags stuff that isn't wrong.
#

import sys, string, re, os

#
# Constants.
#

MAX_LINE_LENGTH = 190   # default: 78
OK = 0
ERROR = -1 

#
# Regular expressions corresponding to style violations.
#

tabs                = re.compile(r"\t+")
comma_space         = re.compile(",[^\n\r ]")

# This one is really tough to get right, so we settle for catching the
# most common mistakes.  Add other operators as necessary and/or feasible.
operator_space      = re.compile("(\w(\+|\-|\*|\<|\>|\=)\w)" + \
                                 "|(\w(\=\=|\<\=|\>\=)\w)")
comment_line        = re.compile("^\s*\/\*.*\*\/\s*$")
open_comment_space  = re.compile("\/\*[^ *\n\r]")
close_comment_space = re.compile("[^ *]\*\/")
paren_curly_space   = re.compile("\)\{")
space_before_paren  = re.compile("if\(|while\(|for\(")
c_plus_plus_comment = re.compile("\/\/")
semi_space          = re.compile(";[^ \s]")

def check_line(filename, line, n):
    """
    Check a line of code for style mistakes.
    """
    # Strip the terminal newline.
    line = line[:-1]
    err_cnt = 0

    # 禁止使用TAB，应全部转换为空格
    if tabs.search(line):
        print "File: %s, line %d: [TABS]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    # 限制每行代码长度不得超过MAX_LINE_LENGTH
    if len(line) > MAX_LINE_LENGTH:
        print "File: %s, line %d: [TOO LONG (%d CHARS)]:\n%s" % \
              (filename, n, len(line), line)
        err_cnt = err_cnt + 1

    # 要求所有逗号后面必须有空格，包括注释里的内容
    if comma_space.search(line):
        if not comment_line.search(line):
            print "File: %s, line %d: [PUT SPACE AFTER COMMA]:\n%s" % \
                  (filename, n, line)
            err_cnt = err_cnt + 1

    # 要求在常见操作符(比如+,-,*,/,>,<等)两侧要添加空格
    if operator_space.search(line):
        if not comment_line.search(line):
            # 排除字符串内有类似"mpiag-smsc.fifo"这样的字符串被误匹配
            sections_in_quotes = re.findall(r'"(.*?)"', line)
            operator_in_string = False
            for section in sections_in_quotes:
                if operator_space.search(section):
                    operator_in_string = True
                    break
            if not operator_in_string:
                print "File: %s, line %d: [PUT SPACE AROUND OPERATORS]:\n%s" % \
                      (filename, n, line)
                err_cnt = err_cnt + 1

    # 我们希望/* Comments */而不是/*Comments*/
    if open_comment_space.search(line):
        print "File: %s, line %d: [PUT SPACE AFTER OPEN COMMENT]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    # 我们希望/* Comments */而不是/*Comments*/
    if close_comment_space.search(line):
        print "File: %s, line %d: [PUT SPACE BEFORE CLOSE COMMENT]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1
        
    # 在){之间要求添加一个空格
    if paren_curly_space.search(line):
        print "File: %s, line %d: [PUT SPACE BETWEEN ) AND {]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    # 在if/while/for与(之间要求添加一个空格
    if space_before_paren.search(line):
        print "File: %s, line %d: [PUT SPACE BETWEEN if/while/for AND (]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    # 要求使用传统的C注释方式：/* ... */
    if c_plus_plus_comment.search(line):
        print "File: %s, line %d: [DON'T USE C++ COMMENTS]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    # 要求分号后面放置一个空格或换行符
    if semi_space.search(line):
        print "File: %s, line %d: [PUT SPACE/NEWLINE AFTER SEMICOLON]:\n%s" % \
              (filename, n, line)
        err_cnt = err_cnt + 1

    if err_cnt == 0:
        return OK
    else:
        return ERROR


def check_file(filename):
    file = open(filename, "r")
    lines = file.readlines()
    file.close()

    err_lines_cnt = 0;
    for i in range(len(lines)):
        result = check_line(filename, lines[i], i + 1)  # Start on line 1.
        if result == ERROR:
            err_lines_cnt = err_lines_cnt + 1

    if err_lines_cnt == 0:
        return OK
    else:
        return ERROR

#
# Main body of program.
#

usage = "usage: c_style_check filename1 [filename2 ...]"

def main():
    if len(sys.argv) < 2:
        print usage
        #raise SystemExit
        os._exit(os.EX_USAGE)
    
    for filename in sys.argv[1:]:
        print "=====> checking file: %s" % filename
        check_result = check_file(filename)
        if check_result == ERROR:
            print "=====> checking file: %s Error!" % filename
            os._exit(os.EX_DATAERR)
        else:
            print "=====> checking file: %s OK" % filename

    os._exit(os.EX_OK)

            
if __name__ == '__main__':
    main()

