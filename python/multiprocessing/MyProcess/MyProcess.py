import multiprocessing
import time, os
class MyProcess(multiprocessing.Process):
	def run(self):
		print("Child goes here, pid %d" % os.getpid())
		try:
			for i in range(10):
				print("Child counting " +  str(i))
				time.sleep(1)
		except Exception as e:
			print(e)
			
		print("Child exiting, pid %d" % os.getpid())