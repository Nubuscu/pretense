from argparse import ArgumentParser
import os
import importlib.util
import sys

from python_graphql_client import GraphqlClient
from requests import HTTPError

TOPIC_DIR = os.path.join(os.getcwd(), "topics")
KNOWN_TOPICS = [f.split(".")[0] for f in os.listdir(TOPIC_DIR) if f.endswith(".py")]

CLIENT = GraphqlClient(endpoint="http://localhost:8081/query", verify=False)

MUTATION = """
mutation createTopicWithReview($topicName: String!, $reviewName: String!, $reviewBody: String!, $albumNames: [String!]!, $metaLabels: [String!]!) {
    createTopicWithReview(topicName: $topicName, reviewName: $reviewName, reviewBody: $reviewBody, albumNames: $albumNames, metaLabels: $metaLabels) {
        id
    }
}
"""


def get_args():
    parser = ArgumentParser()
    parser.add_argument(
        "-f",
        "--file",
        dest="filename",
        required=True,
        help="file name to write. use `all` to write everything in ./topics",
        choices=KNOWN_TOPICS + ["all"],
    )
    return parser.parse_args()


def main():
    args = get_args()
    to_write = KNOWN_TOPICS if args.filename == "all" else [args.filename]
    to_write = [os.path.join(TOPIC_DIR, f"{fname}.py") for fname in to_write]
    for i, topic_file in enumerate(to_write):
        module_name = f"topic={i}"
        spec = importlib.util.spec_from_file_location(module_name, topic_file)
        module = importlib.util.module_from_spec(spec)
        sys.modules[module_name] = module
        spec.loader.exec_module(module)
        albums = module.ALBUMS
        topic = module.TOPIC
        title = module.TITLE
        body = module.BODY
        meta_labels = module.META

        variables = {
            "topicName": topic,
            "reviewName": title,
            "reviewBody": body,
            "albumNames": albums,
            "metaLabels": meta_labels,
        }
        try:
            CLIENT.execute(query=MUTATION, variables=variables)
        except HTTPError as err:
            print(err.response.json())
            breakpoint()
            raise err


if __name__ == "__main__":
    main()
