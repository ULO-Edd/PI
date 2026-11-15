from fastapi import FastAPI, Depends, HTTPException, status
from fastapi.security import OAuth2PasswordRequestForm, OAuth2PasswordBearer
from sqlalchemy import create_engine, Column, Integer, String, DateTime, ForeignKey, Enum
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, Session, relationship
from pydantic import BaseModel
from datetime import datetime, timedelta
from passlib.context import CryptContext
from jose import JWTError, jwt
import enum
import os

# --- Configuration ---
DATABASE_URL = "sqlite:///./test.db"
SECRET_KEY = "your-secret-key"
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 30

# --- Database Setup ---
engine = create_engine(
    DATABASE_URL, connect_args={"check_same_thread": False}
)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

# --- Database Models ---
class UserRole(str, enum.Enum):
    attendee = "attendee"
    organiser = "organiser"
    admin = "admin"

class EventStatus(str, enum.Enum):
    draft = "draft"
    published = "published"
    cancelled = "cancelled"

class User(Base):
    __tablename__ = "users"
    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True)
    hashed_password = Column(String)
    role = Column(Enum(UserRole))
    events = relationship("Event", back_populates="organiser")
    registrations = relationship("Registration", back_populates="user")

class Event(Base):
    __tablename__ = "events"
    id = Column(Integer, primary_key=True, index=True)
    title = Column(String, index=True)
    description = Column(String)
    location = Column(String)
    start_datetime = Column(DateTime)
    end_datetime = Column(DateTime)
    capacity = Column(Integer)
    status = Column(Enum(EventStatus))
    organiser_id = Column(Integer, ForeignKey("users.id"))
    organiser = relationship("User", back_populates="events")
    registrations = relationship("Registration", back_populates="event")

class Registration(Base):
    __tablename__ = "registrations"
    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(Integer, ForeignKey("users.id"))
    event_id = Column(Integer, ForeignKey("events.id"))
    user = relationship("User", back_populates="registrations")
    event = relationship("Event", back_populates="registrations")

# Create database tables
Base.metadata.create_all(bind=engine)

# --- Pydantic Schemas ---
class UserBase(BaseModel):
    username: str
    role: UserRole

class UserCreate(UserBase):
    password: str

class UserSchema(UserBase):
    id: int
    class Config:
        from_attributes = True

class EventBase(BaseModel):
    title: str
    description: str
    location: str
    start_datetime: datetime
    end_datetime: datetime
    capacity: int
    status: EventStatus

class EventCreate(EventBase):
    pass
    
class EventUpdate(BaseModel):
    title: str | None = None
    description: str | None = None
    location: str | None = None
    start_datetime: datetime | None = None
    end_datetime: datetime | None = None
    capacity: int | None = None
    status: EventStatus | None = None

class EventSchema(EventBase):
    id: int
    organiser_id: int
    organiser: UserSchema
    class Config:
        from_attributes = True

class Token(BaseModel):
    access_token: str
    token_type: str

class TokenData(BaseModel):
    username: str | None = None
    
# --- Security ---
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)

def get_password_hash(password):
    return pwd_context.hash(password)

def create_access_token(data: dict, expires_delta: timedelta | None = None):
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt

def decode_access_token(token: str):
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])
        username: str = payload.get("sub")
        if username is None:
            return None
        return username
    except JWTError:
        return None

# --- Dependency ---
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def get_current_user(token: str = Depends(oauth2_scheme), db: Session = Depends(get_db)):
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Could not validate credentials",
        headers={"WWW-Authenticate": "Bearer"},
    )
    username = decode_access_token(token)
    if username is None:
        raise credentials_exception
    user = db.query(User).filter(User.username == username).first()
    if user is None:
        raise credentials_exception
    return user

# --- FastAPI App ---
app = FastAPI(title="Event Management API")

# --- API Endpoints ---
@app.post("/token", response_model=Token, tags=["Auth"])
def login_for_access_token(db: Session = Depends(get_db), form_data: OAuth2PasswordRequestForm = Depends()):
    user = db.query(User).filter(User.username == form_data.username).first()
    if not user or not verify_password(form_data.password, user.hashed_password):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect username or password",
            headers={"WWW-Authenticate": "Bearer"},
        )
    access_token = create_access_token(data={"sub": user.username})
    return {"access_token": access_token, "token_type": "bearer"}

@app.post("/users/", response_model=UserSchema, tags=["Users"])
def create_user(user: UserCreate, db: Session = Depends(get_db)):
    db_user = db.query(User).filter(User.username == user.username).first()
    if db_user:
        raise HTTPException(status_code=400, detail="Username already registered")
    hashed_password = get_password_hash(user.password)
    db_user = User(username=user.username, hashed_password=hashed_password, role=user.role)
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user

@app.post("/events/", response_model=EventSchema, tags=["Events"])
def create_event(event: EventCreate, db: Session = Depends(get_db), current_user: User = Depends(get_current_user)):
    if current_user.role not in [UserRole.organiser, UserRole.admin]:
        raise HTTPException(status_code=403, detail="Not authorized to create events")
    db_event = Event(**event.dict(), organiser_id=current_user.id)
    db.add(db_event)
    db.commit()
    db.refresh(db_event)
    return db_event

@app.get("/events/", response_model=list[EventSchema], tags=["Events"])
def read_events(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    events = db.query(Event).filter(Event.status == EventStatus.published).offset(skip).limit(limit).all()
    return events

@app.get("/events/{event_id}", response_model=EventSchema, tags=["Events"])
def read_event(event_id: int, db: Session = Depends(get_db)):
    db_event = db.query(Event).filter(Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    return db_event

@app.put("/events/{event_id}", response_model=EventSchema, tags=["Events"])
def update_event(event_id: int, event: EventUpdate, db: Session = Depends(get_db), current_user: User = Depends(get_current_user)):
    db_event = db.query(Event).filter(Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    if db_event.organiser_id != current_user.id and current_user.role != UserRole.admin:
        raise HTTPException(status_code=403, detail="Not authorized to update this event")
    for var, value in vars(event).items():
        if value is not None:
            setattr(db_event, var, value)
    db.commit()
    db.refresh(db_event)
    return db_event

@app.delete("/events/{event_id}", status_code=status.HTTP_204_NO_CONTENT, tags=["Events"])
def delete_event(event_id: int, db: Session = Depends(get_db), current_user: User = Depends(get_current_user)):
    db_event = db.query(Event).filter(Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    if db_event.organiser_id != current_user.id and current_user.role != UserRole.admin:
        raise HTTPException(status_code=403, detail="Not authorized to delete this event")
    db.delete(db_event)
    db.commit()
    return

@app.post("/events/{event_id}/register", status_code=status.HTTP_201_CREATED, tags=["Events"])
def register_for_event(event_id: int, db: Session = Depends(get_db), current_user: User = Depends(get_current_user)):
    db_event = db.query(Event).filter(Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    if db_event.status != EventStatus.published:
        raise HTTPException(status_code=400, detail="Event is not published")
    if len(db_event.registrations) >= db_event.capacity:
        raise HTTPException(status_code=400, detail="Event is full")
    existing_registration = db.query(Registration).filter(Registration.event_id == event_id, Registration.user_id == current_user.id).first()
    if existing_registration:
        raise HTTPException(status_code=400, detail="Already registered for this event")
    db_registration = Registration(user_id=current_user.id, event_id=event_id)
    db.add(db_registration)
    db.commit()
    return {"message": "Successfully registered"}

@app.delete("/events/{event_id}/unregister", status_code=status.HTTP_204_NO_CONTENT, tags=["Events"])
def unregister_from_event(event_id: int, db: Session = Depends(get_db), current_user: User = Depends(get_current_user)):
    db_registration = db.query(Registration).filter(Registration.event_id == event_id, Registration.user_id == current_user.id).first()
    if db_registration is None:
        raise HTTPException(status_code=400, detail="Not registered for this event")
    db.delete(db_registration)
    db.commit()
    return

@app.get("/", tags=["Root"])
def read_root():
    return {"message": "Welcome to the Event Management API!"}
