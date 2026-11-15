from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from typing import List
from src.db import models, database
from src.schemas import event as event_schema
from src.schemas import user as user_schema
from src.api import dependencies

router = APIRouter(
    prefix="/events",
    tags=["events"],
)

@router.post("/", response_model=event_schema.Event)
def create_event(
    event: event_schema.EventCreate, 
    db: Session = Depends(database.get_db), 
    current_user: models.User = Depends(dependencies.get_current_user)
):
    if current_user.role not in [models.UserRole.organiser, models.UserRole.admin]:
        raise HTTPException(status_code=403, detail="Not authorized to create events")
    
    db_event = models.Event(**event.dict(), organiser_id=current_user.id)
    db.add(db_event)
    db.commit()
    db.refresh(db_event)
    return db_event

@router.get("/", response_model=List[event_schema.Event])
def read_events(skip: int = 0, limit: int = 100, db: Session = Depends(database.get_db)):
    events = db.query(models.Event).filter(models.Event.status == models.EventStatus.published).offset(skip).limit(limit).all()
    return events

@router.get("/{event_id}", response_model=event_schema.Event)
def read_event(event_id: int, db: Session = Depends(database.get_db)):
    db_event = db.query(models.Event).filter(models.Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    return db_event

@router.put("/{event_id}", response_model=event_schema.Event)
def update_event(
    event_id: int, 
    event: event_schema.EventUpdate, 
    db: Session = Depends(.get_db), 
    current_user: models.User = Depends(dependencies.get_current_user)
):
    db_event = db.query(models.Event).filter(models.Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    
    if db_event.organiser_id != current_user.id and current_user.role != models.UserRole.admin:
        raise HTTPException(status_code=403, detail="Not authorized to update this event")

    for var, value in vars(event).items():
        if value is not None:
            setattr(db_event, var, value)
            
    db.commit()
    db.refresh(db_event)
    return db_event

@router.delete("/{event_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_event(
    event_id: int, 
    db: Session = Depends(database.get_db), 
    current_user: models.User = Depends(dependencies.get_current_user)
):
    db_event = db.query(models.Event).filter(models.Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")

    if db_event.organiser_id != current_user.id and current_user.role != models.UserRole.admin:
        raise HTTPException(status_code=403, detail="Not authorized to delete this event")

    db.delete(db_event)
    db.commit()
    return {"ok": True}

@router.post("/{event_id}/register", status_code=status.HTTP_201_CREATED)
def register_for_event(
    event_id: int, 
    db: Session = Depends(database.get_db), 
    current_user: models.User = Depends(dependencies.get_current_user)
):
    db_event = db.query(models.Event).filter(models.Event.id == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")

    if db_event.status != models.EventStatus.published:
        raise HTTPException(status_code=400, detail="Event is not published")

    if len(db_event.registrations) >= db_event.capacity:
        raise HTTPException(status_code=400, detail="Event is full")

    existing_registration = db.query(models.Registration).filter(
        models.Registration.event_id == event_id,
        models.Registration.user_id == current_user.id
    ).first()
    if existing_registration:
        raise HTTPException(status_code=400, detail="Already registered for this event")

    db_registration = models.Registration(user_id=current_user.id, event_id=event_id)
    db.add(db_registration)
    db.commit()
    return {"message": "Successfully registered"}

@router.delete("/{event_id}/unregister", status_code=status.HTTP_204_NO_CONTENT)
def unregister_from_event(
    event_id: int,
    db: Session = Depends(database.get_db),
    current_user: models.User = Depends(dependencies.get_current_user)
):
    db_registration = db.query(models.Registration).filter(
        models.Registration.event_id == event_id,
        models.Registration.user_id == current_user.id
    ).first()

    if db_registration is None:
        raise HTTPException(status_code=400, detail="Not registered for this event")

    db.delete(db_registration)
    db.commit()
    return {"ok": True}
