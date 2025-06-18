from fastapi import APIRouter, Response, status
from schemas.user_schema import UserCreate, UserResponse, UserUpdate
from services.user_service import (
    create_user, get_user, get_all_users,
    update_user, delete_user_by_id, get_user_by_email, get_user_count
)

router = APIRouter(prefix="/users", tags=["Users"])

@router.post("/", response_model=UserResponse, status_code=status.HTTP_201_CREATED)
async def create(user: UserCreate):
    return create_user(user)

@router.get("/", response_model=list[UserResponse])
async def list_users():
    return get_all_users()

@router.put("/{user_id}", response_model=UserResponse)
async def update(user_id: str, user: UserUpdate):
    return update_user(user_id, user)

@router.delete("/{user_id}", status_code=status.HTTP_204_NO_CONTENT)
async def delete(user_id: str):
    delete_user_by_id(user_id)

    return Response(status_code=status.HTTP_204_NO_CONTENT)

@router.get("/email/{email}", response_model=UserResponse)
async def by_email(email: str):
    return get_user_by_email(email)

@router.get("/count")
async def count():
    return {"count": get_user_count()}

@router.get("/{user_id}", response_model=UserResponse)
async def get(user_id: str):
    return get_user(user_id)
