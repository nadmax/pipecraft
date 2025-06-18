import request from 'supertest';
import app from '../src/app';
import { users } from '../src/models/user';

describe('API Integration Tests', () => {
    beforeEach(() => {
        users.length = 0;
    });

    it('GET / should return hello world', async () => {
        const res = await request(app).get('/');

        expect(res.statusCode).toBe(200);
        expect(res.body.message).toBe('Hello, World!');
    });

    it('GET /api/health should return API health info', async () => {
        const res = await request(app).get('/api/health');

        expect(res.statusCode).toBe(200);
        expect(res.body.message).toBe('API is running');
        expect(new Date(res.body.timestamp)).toBeInstanceOf(Date);
    });

    it('POST /api/users should create user', async () => {
        const res = await request(app).post('/api/users').send({
            name: 'Test User',
            email: 'test@example.com'
        });
    
        expect(res.statusCode).toBe(201);
        expect(res.body.data.name).toBe('Test User');
    });

    it('GET /api/users should return list', async () => {
        await request(app).post('/api/users').send({ 
            name: 'User', 
            email: 'u@e.com' 
        });

        const res = await request(app).get('/api/users');

        expect(res.body.count).toBe(1);
    });
});
