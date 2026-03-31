from fastapi import FastAPI
from contextlib import asynccontextmanager
from app.api.routes import router
from app.models.load_model import load_all_models
import uvicorn

# This function runs before the server accepts any web traffic
@asynccontextmanager
async def lifespan(app: FastAPI):
    load_all_models()
    yield
    print("Server shutting down. Cleaning up memory.")

app = FastAPI(
    title="Python AI Threat Engine",
    description="AI-driven log analysis and threat detection microservice",
    version="1.0.0",
    lifespan=lifespan # Attach the startup logic here
)

app.include_router(router)

@app.get("/")
async def health_check():
    return {"status": "online", "service": "Python AI Engine"}

if __name__ == "__main__":
    uvicorn.run("app.main:app", host="0.0.0.0", port=8000, reload=True)