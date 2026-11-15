from sqlalchemy import Column, Integer, String, DateTime, ForeignKey, Enum
from sqlalchemy.orm import relationship
from .database import Base
import enum

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
