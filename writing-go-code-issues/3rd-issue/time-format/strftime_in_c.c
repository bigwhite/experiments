#include <stdio.h>
#include <time.h>

int main() {
	time_t now = time(NULL);

	struct tm *localTm;
	localTm = localtime(&now);

	char strTime[100];
	strftime(strTime, sizeof(strTime),  "%Y-%m-%d %H:%M:%S", localTm);
	printf("%s\n", strTime);

	return 0;
}
