import { sveltekit } from '@sveltejs/kit/vite';
import { optimizeDeps } from 'vite';

const config = {
	plugins: [sveltekit()],
	optimizeDeps: {
		entries: []
	}
};

export default config;
