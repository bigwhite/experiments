
"""event.py

   A wrapper of epoll or poll events

"""

import select

try:
    from select import epoll
    EVENT_IN = select.EPOLLIN
    EVENT_PRI = select.EPOLLPRI
    EVENT_OUT = select.EPOLLOUT
    EVENT_ERR = select.EPOLLERR
    EVENT_HUP = select.EPOLLHUP
    EVENT_NVAL = 0

except ImportError:

    from select import poll
    EVENT_IN = select.POLLIN
    EVENT_PRI = select.POLLPRI
    EVENT_OUT = select.POLLOUT
    EVENT_ERR = select.POLLERR
    EVENT_HUP = select.POLLHUP
    EVENT_NVAL = select.POLLNVAL
