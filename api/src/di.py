from dependency_injector import containers, providers
from src.graph import GraphRepository


class Container(containers.DeclarativeContainer):

    wiring_config = containers.WiringConfiguration(
        packages=["src.routes", "src.graph"],
    )

    config = providers.Configuration()
    config.graph.host.from_env("DB_HOST")
    config.graph.port.from_env("DB_PORT")
    config.graph.username.from_env("DB_USER")
    config.graph.password.from_env("DB_PASS")

    graph_repo = providers.ThreadSafeSingleton(
        GraphRepository,
        host=config.graph.host,
        port=config.graph.port,
        username=config.graph.username,
        password=config.graph.password,
    )
