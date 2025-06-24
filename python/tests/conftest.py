import pytest
import pytest_asyncio
from httpx import AsyncClient, ASGITransport
from fastapi import FastAPI
from src.app import app

@pytest_asyncio.fixture
async def async_client() -> AsyncClient:
    transport = ASGITransport(app=app)

    async with AsyncClient(transport=transport, base_url="http://test") as client:
        yield client

@pytest_asyncio.fixture
async def created_user(async_client: AsyncClient):
    response = await async_client.post("/users/", json={
        "name": "Test User",
        "email": "test@example.com"
    })

    assert response.status_code == 201, f"Failed to create user: {response.text}"

    data = response.json()

    yield data
