import multiprocessing
import time, os, sys
from MyProcess import MyProcess

if __name__ == "__main__":
	print("Parent start pid %d" % os.getpid())
	basedir = os.path.dirname(sys.argv[0])
	multiprocessing.set_executable(os.path.join(basedir, "sp.exe"))

	p = MyProcess.MyProcess()
	p.start()
	
	p.join()
	print("Parent existing pid %d" % os.getpid())
