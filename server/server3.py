from robyn import Robyn

app = Robyn(__file__)

@app.get("/")
async def h(request):
    return "Response from Server 3"

app.start(port=5002)
