const express = require('express');
const swaggerUi = require('swagger-ui-express');
const swaggerSpec = require('./config/swagger');
require('dotenv').config();

const userRouter = require('./routes/userRoutes');
const eventRouter = require('./routes/eventRoutes');

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());
app.use(userRouter);
app.use(eventRouter);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec));

app.get('/', (req, res) => {
  res.send('Event Management API');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
