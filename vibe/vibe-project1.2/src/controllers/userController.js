const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const prisma = require('../../prisma/client');

const register = async (req, res) => {
  try {
    const { name, email, password, role } = req.body;
    const hashedPassword = await bcrypt.hash(password, 8);
    const user = await prisma.user.create({
      data: {
        name,
        email,
        password: hashedPassword,
        role,
      },
    });
    res.status(201).send({ user });
  } catch (error) {
    res.status(400).send(error);
  }
};

const login = async (req, res) => {
  try {
    const { email, password } = req.body;
    const user = await prisma.user.findUnique({
      where: {
        email,
      },
    });

    if (!user) {
      return res.status(400).send({ error: 'Unable to login' });
    }

    const isMatch = await bcrypt.compare(password, user.password);

    if (!isMatch) {
      return res.status(400).send({ error: 'Unable to login' });
    }

    const token = jwt.sign({ id: user.id.toString() }, process.env.JWT_SECRET);
    res.send({ user, token });
  } catch (error) {
    res.status(400).send(error);
  }
};

module.exports = {
  register,
  login,
};
