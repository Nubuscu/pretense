import { multiTopic, singleTopic } from "$lib/topicProcessing";

let root = `${process.env.BACKEND}`;

async function fetchTopic(id) {
    let resp = await fetch(`${root}/v1/topics/${id}`);
    return await resp.json();
}
async function fetchTopicIds() {
    let url = `${root}/v1/topics`
    return await fetch(url)
        .then((response) => response.json())
        .then((data) => {
            return data.map((val) => val.id);
        })
        .catch((err) => {
            console.error("it went bad", url, err);
            return [];
        });
}

export async function load({ params }) {
    const ids = await fetchTopicIds()
    const topics = await Promise.all(ids.map(async (id) => await fetchTopic(id)));
    const singles = {}
    topics.map(t => [t.id, singleTopic(t)]).forEach(x => {
        let [id, graphData] = x
        singles[id] = graphData
    })
    const all_processed = multiTopic(topics);
    return {
        topics: topics,
        singles: singles,
        multi: all_processed,
    }
}