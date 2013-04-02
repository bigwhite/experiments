#
#
#

"""reactor.py

"""

try:
    from select import epoll
    EVENT_NOTIFY_FACILITY = 'epoll'
except ImportError:
    from select import poll
    EVENT_NOTIFY_FACILITY = 'poll'
    
class Reactor(object):
    """

    """

    def __init__(self, 


