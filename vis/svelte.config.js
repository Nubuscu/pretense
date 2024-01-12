import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		alias: {
			$houdini: './$houdini',
		},
		// paths: {
		// 	base: process.argv.includes('dev') ? '' : process.env.BASE_PATH
		// },
		prerender: {
			handleHttpError:({path, referrer, message}) => {
				console.error("Failed to build path", path, referrer)
				throw new Error(message)
			}
		}
	}
};

export default config;
