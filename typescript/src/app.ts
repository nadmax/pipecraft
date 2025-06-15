import express from 'express';
import userRoutes from './routes/users';
import { errorHandler } from './middlewares/errorHandler';

const app = express();

app.use(express.json());
app.use('/', (_req, res) => {
    res.json({ success: true, message: 'Hello, World!'});
});
app.use('/api/users', userRoutes);
app.get('/api/health', (_req, res) => {
    res.json({ success: true, message: 'API is running', timestamp: new Date() });
});
app.use((_req, res, next) => {
    res.status(404).json({ success: false, message: 'Oops... Page not found' });
    next();
});
app.use(errorHandler);

export default app;
