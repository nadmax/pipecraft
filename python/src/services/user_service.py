from datetime import datetime, timezone
from fastapi import HTTPException
from schemas.user_schema import UserCreate, UserUpdate
import uuid

db = {}

def create_user(user: UserCreate):
    user_id = str(uuid.uuid4())
    now = datetime.now(timezone.utc)
    db[user_id] = {
        "id": user_id,
        "name": user.name,
        "email": user.email,
        "created_at": now,
        "updated_at": now
    }

    return db[user_id]

def get_user(user_id: str):
    if user_id not in db:
        raise HTTPException(status_code=404, detail="User not found")

    return db[user_id]

def get_all_users():
    return list(db.values())

def update_user(user_id: str, update: UserUpdate):
    user = get_user(user_id)
    data = update.dict(exclude_unset=True)
    for key, value in data.items():
        user[key] = value

    user["updated_at"] = datetime.now(timezone.utc)
    
    return user

def delete_user_by_id(user_id: str):
    if user_id not in db:
        raise HTTPException(status_code=404, detail="User not found")

    del db[user_id]

def get_user_by_email(email: str):
    for user in db.values():
        if user["email"] == email:
            return user

    raise HTTPException(status_code=404, detail="User not found")

def get_user_count():
    return len(db)
