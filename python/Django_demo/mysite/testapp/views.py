from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse

import logging
# Get an instance of a logger
logger = logging.getLogger(__name__) # __name__ : "testapp.views"

import json

def fuck(request):
    logger.info(request.GET)

    # 下面两行代码也可以改成一行（需要 from django.http import JsonResponse）
    #  ret = JosnResponse({"a":"AA", "b":"bb"})
    ret = HttpResponse(json.dumps({"a":"AA", "b":"bb"}))
    ret["Content-Type"] = "application/json"

    return ret
    #return HttpResponse("Hello, world. You're at the polls index.")