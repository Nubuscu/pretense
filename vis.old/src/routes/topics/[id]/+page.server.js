import { singleTopic } from "$lib/topicProcessing";
import { createClient, gql } from '@urql/svelte';


let root = `${process.env.BACKEND}`;

const client = createClient({
  url: `${root}/query`
})

const TOPIC_QUERY = gql`
  query singleTopic($id: ID) {
    topics(where: {id: $id}) {
      edges {
        node {
          id
          name
          reviewedBy {
            id
            name
            body
          }
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
  const resp = await client.query(TOPIC_QUERY, { id: parseInt(params.id) }).toPromise();
  if (resp.error) {
    console.error(resp.error)
    return;
  }
  const content = resp.data.topics.edges[0]?.node

  return {
    id: params.id,
    title: content.name,
    reviews: content.reviewedBy.map(r => {
      return { id: r.id, title: r.name ?? "", body: r.body ?? "" }
    }),
    graphInput: singleTopic(content),
  }
}
