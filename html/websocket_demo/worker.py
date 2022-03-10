from sys import stdout, stderr, stdin
from time import sleep
import sys,os

# for Environment variables, like REQUEST_URI below, refer to https://github.com/joewalnes/websocketd/wiki/Environment-variables
stderr.write("Websocket requested URI: " + os.environ["REQUEST_URI"] + "\n")
stderr.flush()

# Count from 1 to 10 with a sleep
for count in range(0, 10):
  print(count + 1)
  stdout.flush()
  sleep(0.5)
  line = sys.stdin.readline()
  stderr.write(line)
  stderr.flush()