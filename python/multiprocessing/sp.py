import multiprocessing, os
if __name__=="__main__":
   print("Child  start pid %d" % os.getpid())
   multiprocessing.freeze_support()
   print("return from freeze\n")  # You will never see this, cause freeze_support will not return.
