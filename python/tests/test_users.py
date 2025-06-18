import pytest

test_user = {
    "name": "Test User",
    "email": "test@example.com"
}

async def test_create_user(async_client, created_user):
    assert created_user["name"] == test_user["name"]
    assert created_user["email"] == test_user["email"]


async def test_get_user(async_client, created_user):
    user_id = created_user["id"]
    response = await async_client.get(f"/users/{user_id}")

    assert response.status_code == 200
    assert response.json()["id"] == user_id


async def test_update_user(async_client, created_user):
    user_id = created_user["id"]
    updated_data = {"name": "Updated User", "email": "updated@example.com"}
    response = await async_client.put(f"/users/{user_id}", json=updated_data)

    assert response.status_code == 200
    assert response.json()["name"] == "Updated User"
    assert response.json()["email"] == "updated@example.com"


async def test_list_users(async_client):
    response = await async_client.get("/users/")

    assert response.status_code == 200
    assert isinstance(response.json(), list)


async def test_get_user_by_email(async_client):
    response = await async_client.get(f"/users/email/{test_user['email']}")

    assert response.status_code == 200
    assert response.json()["email"] == test_user["email"]


async def test_get_user_count(async_client):
    response = await async_client.get("/users/count")

    assert response.status_code == 200
    assert "count" in response.json()


async def test_delete_user(async_client, created_user):
    user_id = created_user["id"]
    response = await async_client.delete(f"/users/{user_id}")

    assert response.status_code == 204
    assert response.text == ""
