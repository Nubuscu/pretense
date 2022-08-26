import { multiTopic } from "$lib/topicProcessing";

let root = `${import.meta.env.VITE_BACKEND_HOST}:${import.meta.env.VITE_BACKEND_PORT
    }`;

async function fetchTopic(id) {
    let resp = await fetch(`${root}/v1/topics/${id}`);
    return await resp.json();
}
async function fetchTopicIds() {
    return await fetch(`${root}/v1/topics`)
        .then((response) => response.json())
        .then((data) => {
            return data.map((val) => val.id);
        })
        .catch((err) => {
            console.error("it went bad", err.code);
            return [];
        });
}

export async function load({ params }) {
    const ids = await fetchTopicIds()
    const topics = await Promise.all(ids.map(async (id) => await fetchTopic(id)));
    const processed = multiTopic(topics);

    return {
        topics: topics,
        processed: processed,
    }
}