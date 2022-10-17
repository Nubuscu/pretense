import { multiTopic, singleTopic } from "$lib/topicProcessing";
import { createClient, gql } from '@urql/svelte';


let root = `${process.env.BACKEND}`;

const client = createClient({
  url: `${root}/v1/graphql`
})

const ALL_TOPICS_QUERY = gql`
  query allTopics {
    topics {
      id
      name
      albums {
        id
        name
        artists {
          id
          name
        }
      }
    }
  }`


export async function load({ params }) {
  const resp = await client.query(ALL_TOPICS_QUERY).toPromise();
  if (resp.error) {
    console.error(resp.error)
    return;
  }
  const topics = resp.data.topics

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