from fastapi import APIRouter
from schemas.user_schema import UserCreate, UserResponse, UserUpdate
from services.user_service import (
    create_user, get_user, get_all_users,
    update_user, delete_user_by_id, get_user_by_email, get_user_count
)

router = APIRouter(prefix="/users", tags=["Users"])

@router.post("/", response_model=UserResponse)
def create(user: UserCreate):
    return create_user(user)

@router.get("/", response_model=list[UserResponse])
def list_users():
    return get_all_users()

@router.get("/{user_id}", response_model=UserResponse)
def get(user_id: str):
    return get_user(user_id)

@router.put("/{user_id}", response_model=UserResponse)
def update(user_id: str, user: UserUpdate):
    return update_user(user_id, user)

@router.delete("/{user_id}")
def delete(user_id: str):
    return delete_user_by_id(user_id)

@router.get("/email/{email}", response_model=UserResponse)
def by_email(email: str):
    return get_user_by_email(email)

@router.get("/count")
def count():
    return {"count": get_user_count()}
