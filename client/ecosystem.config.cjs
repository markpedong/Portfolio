module.exports = {
	apps: [
		{
			name: 'admin',
			script: 'serve',
			env: {
				PM2_SERVE_PATH: './dist',
				PM2_SERVE_PORT: 6602,
				PM2_SERVE_SPA: 'true',
				NODE_ENV: 'production'
			}
		}
	]
}
