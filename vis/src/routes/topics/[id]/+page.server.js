import { multiTopic, singleTopic } from "$lib/topicProcessing";
import { createClient, gql } from '@urql/svelte';


let root = `${process.env.BACKEND}`;

const client = createClient({
  url: `${root}/v1/graphql`
})

const TOPIC_QUERY = gql`
  query singleTopic($id: Int) {
    topics(id: $id) {
      id
      name
      reviews {
        id
        title
        body
      }
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
  const resp = await client.query(TOPIC_QUERY, { id: parseInt(params.id) }).toPromise();
  if (resp.error) {
    console.error(resp.error)
    return;
  }
  const content = resp.data.topics[0]

  return {
    id: params.id,
    title: content.name,
    reviews: content.reviews.map(r => {
      return { id: r.id, title: r.title ?? "", body: r.body ?? "" }
    }),
    graphInput: singleTopic(content),
  }
}
