#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <syslog.h>

extern void daemonize(const char *cmd);

int main(int argc, char* argv[]) 
{
	char *cmd;
	if ((cmd = strrchr(argv[0], '/')) == NULL)
		cmd = argv[0];
	else
		cmd++;
	
	daemonize(cmd);
	
	FILE *filep = fopen("guixuan.log", "wb+");
	int i = 0;
	while(1) {
		fprintf(filep, "BEEP [%04d]...\n", i++);
		syslog(LOG_ERR, "BEEP [%04d]......\n", i);
		fflush(filep);
		sleep(2);
	}
}