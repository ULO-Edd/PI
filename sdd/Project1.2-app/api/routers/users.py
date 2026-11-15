from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from src.db import models, database
from src.schemas import user as user_schema
from src.core.security import get_password_hash

router = APIRouter(
    prefix="/users",
    tags=["users"],
)

@router.post("/", response_model=user_schema.User)
def create_user(user: user_schema.UserCreate, db: Session = Depends(database.get_db)):
    db_user = db.query(models.User).filter(models.User.username == user.username).first()
    if db_user:
        raise HTTPException(status_code=400, detail="Username already registered")
    hashed_password = get_password_hash(user.password)
    db_user = models.User(username=user.username, hashed_password=hashed_password, role=user.role)
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user
