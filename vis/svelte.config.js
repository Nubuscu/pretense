import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		alias: {
			$houdini: './$houdini',
		},
		prerender: {
			handleHttpError:({path, referrer, message}) => {
				console.error("Failed to build path", path, referrer)
				throw new Error(message)
			}
		}
	}
};

export default config;
