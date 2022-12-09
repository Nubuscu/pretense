import { multiTopic, singleTopic } from "$lib/topicProcessing";
import { createClient, gql } from '@urql/svelte';


let root = `${process.env.BACKEND}`;

const client = createClient({
  url: `${root}/query`
})

const ALL_TOPICS_QUERY = gql`
  query allTopics {
    topics {
      edges {
        node {
          id
          name
          includes {
            id
            name
            by {
              id
              name
            }
          }
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
  const topics = resp.data.topics.edges

  const all_processed = multiTopic(topics);
  return {
    multi: all_processed,
  }
}