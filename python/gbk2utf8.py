import sys, os

def gbk2utf8(fn):
	try:
		with open(fn, "rb+") as f:
			content = f.read().decode("gbk").encode("utf-8")
			f.truncate(0)
			f.seek(0)
			f.write(content)
	except Exception as e:
		print("%s failed. ERR=%s" % (fn, e))
		return
	print("%s ok." % fn)

def walk(fn, exts=None):
	if os.path.isfile(fn):
		if not exts:
			return [fn]
		else:
			_, ext = os.path.splitext(fn)
			if ext in exts:
				return [fn]
			else:
				return []
	fns = []
	for root, dirs, files in os.walk(fn):
		if exts:
			fns = fns + [os.path.join(root, f) for f in files if os.path.splitext(f)[1] in exts]
		else:
			fns = fns + [os.path.join(root, f) for f in files ]
			
		for d in dirs:
			fns = fns + walk(d)
	return fns
	
if __name__ == "__main__":
	files = walk(sys.argv[1], (".cpp", ".h", ".txt"))
	for file in files:
		gbk2utf8(file)