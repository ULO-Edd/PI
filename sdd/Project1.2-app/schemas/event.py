from pydantic import BaseModel
from datetime import datetime
from src.db.models import EventStatus
from src.schemas.user import User

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

class Event(EventBase):
    id: int
    organiser_id: int
    organiser: User

    class Config:
        orm_mode = True
