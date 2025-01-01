import react from '@vitejs/plugin-react-swc'
import path from 'path'
import { defineConfig } from 'vite'
import sassDts from 'vite-plugin-sass-dts'
import tsConfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config
export default defineConfig({
	base: "/",
	server: {
		port: 6602,
	},
	plugins: [react(), sassDts(), tsConfigPaths()],
	resolve: {
		alias: {
			'@': path.resolve(__dirname, './src')
		}
	}
})
