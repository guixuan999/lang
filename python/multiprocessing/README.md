这个例子展示了神奇的multiprocessing.set_executable()和multiprocessing.freeze_support()

1. pip install py2exe
2. 生成mp.exe
    python setup_sp.py py2exe
3. 生成sp.exe
    python setup_mp.py py2exe
4. 运行mp.exe
(py3.9) D:\>dist\mp
Parent start pid 14720
Child  start pid 8728
Child goes here, pid 8728
Child counting 0
Child counting 1
Child counting 2
Child counting 3
Child counting 4
Child counting 5
Child counting 6
Child counting 7
Child counting 8
Child counting 9
Child exiting, pid 8728
Parent existing pid 14720

通过上面的输出，可体会下子进程sp.exe的神奇流程