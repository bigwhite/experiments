
"""source.py

   Definitions of the source which trigger the events.
   There are several sorts of source, such as File Descriptor 
   source, Timer source etc.
"""

class Source(object):
    """class Source

       three source object could be created:
         * file source -- fs = Source(sock)
         * file source with timeout -- fs = Source(sock, 500)
         * pure timer source -- ts = Source(None, 1000)

    """

    __slots__ = ('fileobj', 'timeval')

    def __init__(self, fileobj, timeval = None):
        """ timeval in microseconds """
        self.fileobj = fileobj
        self.timeval = timeval

    @property
    def fileno(self):
        if not self.fileobj:
            return -1
        else:
            return self.fileobj.fileno()

