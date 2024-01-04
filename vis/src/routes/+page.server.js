export const prerender = true;

import { AllTopicsStore } from '$houdini';

/* @type { import('./$houdini').PageLoad } */
export const load = async (event) => {
	const myQuery = new AllTopicsStore();
	const { data } = await myQuery.fetch({ event });

	return data;
};
