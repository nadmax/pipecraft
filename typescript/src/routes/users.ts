import { Router, Request, Response } from 'express';
import { v4 as uuid } from 'uuid';
import { User, users } from '../models/user';

const router = Router();

router.get('/', (_req: Request, res: Response) => {
    res.json({ success: true, data: users, count: users.length });
});

router.get('/:id', (req: Request, res: Response) => {
    const user = users.find(u => u.id === req.params.id);
    if (!user) {
        res.status(404).json({ success: false, message: 'User not found' });

        return;
    }

    res.json({ success: true, data: user });
});

router.post('/', (req: Request, res: Response) => {
    const { name, email } = req.body;
    if (!name || !email) {
        res.status(400).json({ success: false, message: 'Name and email are required' });

        return;
    }

    if (users.find(u => u.email === email)) {
        res.status(409).json({ success: false, message: 'Email already in use' });

        return;
    }

    const newUser: User = {
        id: uuid(),
        name,
        email,
        createdAt: new Date(),
        updatedAt: new Date()
    };

    users.push(newUser);
    res.status(201).json({ success: true, data: newUser });
});

export default router;
