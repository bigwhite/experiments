#!/usr/bin/env python

#
# trigger_sim.py
#

#
# usage;
#   trigger_sim.py tablename oper_type id  
#
#

import socket  
import sys

if __name__ == '__main__':
    # usage: 
    #   trigger_sim.py table oper_type id
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)  
    sock.connect(('localhost', 8877))  
    str = '^' + sys.argv[1] + '|' + sys.argv[2] + '|' + sys.argv[3] + '$'
    sock.send(str)  
    sock.close()  

