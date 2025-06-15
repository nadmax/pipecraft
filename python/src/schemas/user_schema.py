from pydantic import BaseModel, EmailStr
from typing import Optional
from datetime import datetime

class UserBase(BaseModel):
    name: str
    email: EmailStr
    age: int

class UserCreate(UserBase):
    pass

class UserUpdate(BaseModel):
    name: Optional[str]
    email: Optional[EmailStr]
    age: Optional[int]

class UserResponse(UserBase):
    id: str
    created_at: datetime
    updated_at: datetime
