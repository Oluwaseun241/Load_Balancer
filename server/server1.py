from robyn import Robyn

app = Robyn(__file__)

@app.get("/")
async def h(request):
    return "Response from Server 1"

app.start(port=5000)
