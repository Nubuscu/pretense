from argparse import ArgumentParser
import os
import importlib.util
import sys
from src.graph import ContentWriter, GraphRepository

TOPIC_DIR = os.path.join(os.getcwd(), "content", "topics")
KNOWN_TOPICS = [f.split(".")[0] for f in os.listdir(TOPIC_DIR) if f.endswith(".py")]


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
    with GraphRepository() as repo:
        for i, topic_file in enumerate(to_write):
            module_name = f"topic={i}"
            spec = importlib.util.spec_from_file_location(module_name, topic_file)
            module = importlib.util.module_from_spec(spec)
            sys.modules[module_name] = module
            spec.loader.exec_module(module)
            albums = module.ALBUMS
            title = module.TITLE
            body = module.BODY
            ContentWriter(repo).write_topic(title, body, albums)


if __name__ == "__main__":
    main()
