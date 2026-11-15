const express = require('express');
const eventController = require('../controllers/eventController');
const auth = require('../middleware/auth');

const router = new express.Router();

/**
 * @swagger
 * /events:
 *   post:
 *     summary: Create a new event
 *     tags: [Events]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *               description:
 *                 type: string
 *               location:
 *                 type: string
 *               startTime:
 *                 type: string
 *                 format: date-time
 *               endTime:
 *                 type: string
 *                 format: date-time
 *               capacity:
 *                 type: integer
 *               status:
 *                 type: string
 *                 enum: [DRAFT, PUBLISHED, CANCELLED]
 *     responses:
 *       201:
 *         description: Event created successfully
 *       400:
 *         description: Bad request
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 */
router.post('/events', auth, eventController.createEvent);

/**
 * @swagger
 * /events:
 *   get:
 *     summary: Get all events
 *     tags: [Events]
 *     responses:
 *       200:
 *         description: A list of events
 *       500:
 *         description: Internal server error
 */
router.get('/events', eventController.getEvents);

/**
 * @swagger
 * /events/{id}:
 *   get:
 *     summary: Get an event by ID
 *     tags: [Events]
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *     responses:
 *       200:
 *         description: The event
 *       404:
 *         description: Event not found
 *       500:
 *         description: Internal server error
 */
router.get('/events/:id', eventController.getEvent);

/**
 * @swagger
 * /events/{id}:
 *   patch:
 *     summary: Update an event
 *     tags: [Events]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *               description:
 *                 type: string
 *               location:
 *                 type: string
 *               startTime:
 *                 type: string
 *                 format: date-time
 *               endTime:
 *                 type: string
 *                 format: date-time
 *               capacity:
 *                 type: integer
 *               status:
 *                 type: string
 *                 enum: [DRAFT, PUBLISHED, CANCELLED]
 *     responses:
 *       200:
 *         description: Event updated successfully
 *       400:
 *         description: Bad request
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Event not found
 */
router.patch('/events/:id', auth, eventController.updateEvent);

/**
 * @swagger
 * /events/{id}:
 *   delete:
 *     summary: Delete an event
 *     tags: [Events]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *     responses:
 *       200:
 *         description: Event deleted successfully
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Event not found
 *       500:
 *         description: Internal server error
 */
router.delete('/events/:id', auth, eventController.deleteEvent);

/**
 * @swagger
 * /events/{id}/register:
 *   post:
 *     summary: Register for an event
 *     tags: [Events]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *     responses:
 *       201:
 *         description: Registered for the event successfully
 *       400:
 *         description: Bad request
 *       401:
 *         description: Unauthorized
 *       404:
 *         description: Event not found
 */
router.post('/events/:id/register', auth, eventController.registerForEvent);

/**
 * @swagger
 * /events/{id}/unregister:
 *   delete:
 *     summary: Unregister from an event
 *     tags: [Events]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         schema:
 *           type: integer
 *     responses:
 *       200:
 *         description: Unregistered from the event successfully
 *       401:
 *         description: Unauthorized
 *       404:
 *         description: Not registered for this event
 *       500:
 *         description: Internal server error
 */
router.delete('/events/:id/unregister', auth, eventController.unregisterFromEvent);

module.exports = router;

