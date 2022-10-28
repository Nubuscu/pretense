import { createClient, gql } from '@urql/svelte';

let root = `${process.env.BACKEND}`;
const client = createClient({
  url: `${root}/v1/graphql`
})

const TOPIC_TITLES_QUERY = gql`
  query topicTitles {
    topics {
      id
      name
    }
  }
`
export async function load() {
  let retval = []
  const resp = await client.query(TOPIC_TITLES_QUERY).toPromise();
  if (resp.error) {
    console.error(resp.error)
  } else {
    retval = resp.data.topics
  }
  return {
    topics: retval
  };
}