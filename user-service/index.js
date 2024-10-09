'use strict';

const Hapi = require('@hapi/hapi');

const init = async () => {
    const server = Hapi.server({
        port: 3000,
        host: 'localhost'
    });

    server.route({
        method: 'POST',
        path: '/public-api/register',
        handler: (request, h) => {
            const { username, password, email } = request.payload;

            if (!username || !password || !email) {
                return h.response({ error: 'All fields are required' }).code(400);
            }

            const newUser = { username, password, email, createdAt: new Date() };

            return h.response({ message: 'User registered successfully', data: newUser }).code(201);
        }
    });

    server.route({
        method: 'GET',
        path: '/api/users/me',
        handler: (request, h) => {
            const userName = request.headers['x-user-name'];

            if (!userName) {
                return h.response({ error: 'Must Login' }).code(401);
            }

            return h.response({
                username: userName,
                email: "test@mail.com",
                createdAt: new Date()
            });
        }
    });

    // Delete Current User Account (protected route)
    server.route({
        method: 'DELETE',
        path: '/api/users/me',
        handler: (request, h) => {
            const userName = request.headers['x-user-name'];

            if (!userName) {
                return h.response({ error: 'Must Login' }).code(401);
            }

            return h.response({ message: 'Account deleted successfully', data: {username: userName} }).code(200);
        }
    });

    await server.start();
    console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {

    console.log(err);
    process.exit(1);
});

init();
