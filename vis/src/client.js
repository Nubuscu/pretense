import { HoudiniClient } from '$houdini';

export default new HoudiniClient({
	// url: `${process.env.BACKEND}/query`
	url: `http://localhost:8081/query`,
});
