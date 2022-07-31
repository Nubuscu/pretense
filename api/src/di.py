from dependency_injector import containers, providers
from src.graph import GraphRepository, connection_factory


def generate_conn(host, port, username, password):
    conn = connection_factory(host, port, username, password)
    try:
        yield conn
    finally:
        conn.close()


class Container(containers.DeclarativeContainer):

    wiring_config = containers.WiringConfiguration(
        packages=["src.routes", "src.graph"],
    )

    config = providers.Configuration()
    config.graph.host.from_env("DB_HOST")
    config.graph.port.from_env("DB_PORT")
    config.graph.username.from_env("DB_USER")
    config.graph.password.from_env("DB_PASS")

    # TODO manage graph connection here, pass to the repo
    graph_conn = providers.Resource(
        generate_conn,
        host=config.graph.host,
        port=config.graph.port,
        username=config.graph.username,
        password=config.graph.password,
    )
    graph_repo = providers.ThreadSafeSingleton(
        GraphRepository,
        conn=graph_conn,
    )
