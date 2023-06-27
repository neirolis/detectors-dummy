#!/usr/bin/python3 


import io
import time
import json
from aiohttp import web
from PIL import Image



version = "0.0.1"
item_types = ['dummy']

# GET /status 
async def status(req):
    return web.json_response({
        "name": "name of detector",
        "type": "detector",
        "path": "/",
        "version": version,
        "output": {
            "types": item_types,
        },
    })

# POST /
async def parser(req):
    at = time.time()

    data = await req.post()
    file = data['frame'].file
    imgdata = io.BytesIO(file.read())
    img = Image.open(imgdata)

    items = json.loads(data["json"])
    for item in items:
        print(item)





    # w, h = img.size
    print(img.size)

    items = []

    # # do something with `img` and store detected objects to the items array
    # for obj in DETECT(img):
    #     items.append({
    #         'type': obj.type,
    #         'name': obj.name,
    #         'bbox': [
    #             round(obj.x, 3),
    #             round(obj.y, 3),
    #             round(obj.w, 3),
    #             round(obj.h, 3),
    #         ]
    #     })

    # print(items, time.time() - at)
    
    resp = {
        'predicts': items,
        'version': version,
        'time': time.time() - at,
    }
   
    return web.json_response(resp)




app = web.Application()
app.add_routes([
        web.get('/status', status),
        web.post('/', parser),
    ])



if __name__ == '__main__':
    web.run_app(app, port=64465)
