const prisma = require('../../prisma/client');

const createEvent = async (req, res) => {
  try {
    if (req.user.role !== 'ORGANISER') {
      return res.status(403).send({ error: 'Only organisers can create events.' });
    }
    const event = await prisma.event.create({
      data: {
        ...req.body,
        organiserId: req.user.id,
      },
    });
    res.status(201).send(event);
  } catch (error) {
    res.status(400).send(error);
  }
};

const getEvents = async (req, res) => {
  try {
    const events = await prisma.event.findMany();
    res.send(events);
  } catch (error) {
    res.status(500).send();
  }
};

const getEvent = async (req, res) => {
  try {
    const event = await prisma.event.findUnique({
      where: {
        id: parseInt(req.params.id),
      },
    });

    if (!event) {
      return res.status(404).send();
    }

    res.send(event);
  } catch (error) {
    res.status(500).send();
  }
};

const updateEvent = async (req, res) => {
  try {
    const event = await prisma.event.findUnique({
      where: {
        id: parseInt(req.params.id),
      },
    });

    if (!event) {
      return res.status(404).send();
    }

    if (event.organiserId !== req.user.id && req.user.role !== 'ADMINISTRATOR') {
      return res.status(403).send({ error: 'You are not authorized to update this event.' });
    }

    const updatedEvent = await prisma.event.update({
      where: {
        id: parseInt(req.params.id),
      },
      data: req.body,
    });

    res.send(updatedEvent);
  } catch (error) {
    res.status(400).send(error);
  }
};

const deleteEvent = async (req, res) => {
  try {
    const event = await prisma.event.findUnique({
      where: {
        id: parseInt(req.params.id),
      },
    });

    if (!event) {
      return res.status(404).send();
    }

    if (event.organiserId !== req.user.id && req.user.role !== 'ADMINISTRATOR') {
      return res.status(403).send({ error: 'You are not authorized to delete this event.' });
    }

    await prisma.event.delete({
      where: {
        id: parseInt(req.params.id),
      },
    });

    res.send({ message: 'Event deleted successfully.' });
  } catch (error) {
    res.status(500).send();
  }
};

const registerForEvent = async (req, res) => {
  try {
    const event = await prisma.event.findUnique({
      where: {
        id: parseInt(req.params.id),
      },
      include: {
        registrations: true,
      },
    });

    if (!event) {
      return res.status(404).send();
    }

    if (event.registrations.length >= event.capacity) {
      return res.status(400).send({ error: 'Event is full.' });
    }

    const registration = await prisma.registration.create({
      data: {
        userId: req.user.id,
        eventId: parseInt(req.params.id),
      },
    });

    res.status(201).send(registration);
  } catch (error) {
    res.status(400).send(error);
  }
};

const unregisterFromEvent = async (req, res) => {
  try {
    const registration = await prisma.registration.findFirst({
      where: {
        userId: req.user.id,
        eventId: parseInt(req.params.id),
      },
    });

    if (!registration) {
      return res.status(404).send({ error: 'You are not registered for this event.' });
    }

    await prisma.registration.delete({
      where: {
        id: registration.id,
      },
    });

    res.send({ message: 'You have been unregistered from the event.' });
  } catch (error) {
    res.status(500).send();
  }
};

module.exports = {
  createEvent,
  getEvents,
  getEvent,
  updateEvent,
  deleteEvent,
  registerForEvent,
  unregisterFromEvent,
};
