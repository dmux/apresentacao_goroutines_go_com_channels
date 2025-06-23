from fastapi import FastAPI
import asyncio

app = FastAPI()

@app.get('/')
async def read_root():
    await asyncio.sleep(0.001)
    return {"message": "ok"}

if __name__ == '__main__':
    import uvicorn
    uvicorn.run(app, host='0.0.0.0', port=8080)