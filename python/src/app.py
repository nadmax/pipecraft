from fastapi import FastAPI
from middlewares.error_handler import add_middlewares
from routes.users import router as user_router

app = FastAPI(
    title="CRUD API with FastAPI",
    description="A simple CRUD API",
    version="1.0.0"
)

add_middlewares(app)

app.include_router(user_router)

@app.get("/")
async def root():
    return {"message": "Hello, World!"}
