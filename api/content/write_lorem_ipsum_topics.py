from src.graph import ContentWriter

title = "My Test Topic"

body = """
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras et consectetur lacus. Fusce ante ligula, viverra in sapien aliquam, imperdiet iaculis dui. Vestibulum accumsan urna ac arcu varius, ac maximus tellus egestas. Maecenas accumsan arcu vulputate, interdum odio quis, tristique nunc. Mauris eleifend ex nec lobortis dictum. Donec ut tellus a lorem rhoncus venenatis. Nunc condimentum purus eu magna cursus, a luctus felis euismod. Nam dignissim, magna non condimentum aliquet, urna risus pulvinar nibh, nec venenatis dui nisi non risus. Quisque efficitur leo non mauris lacinia egestas.
Cras a porttitor risus, in euismod nisi. Duis imperdiet ut est vel gravida. Maecenas a egestas ex. Donec mollis non diam in efficitur. Praesent rutrum turpis eget felis feugiat faucibus. Maecenas sit amet dictum velit, sit amet faucibus lorem. Vestibulum sit amet sem tellus. Nullam elit orci, commodo ac cursus at, ullamcorper ac velit. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.
Pellentesque ultricies vehicula accumsan. Fusce sollicitudin pellentesque aliquam. Nunc mattis ante at faucibus bibendum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque a porttitor elit. Curabitur venenatis viverra vulputate. Maecenas id urna non tortor ultricies dapibus id eu nulla. Donec mattis sodales justo nec gravida. Nam sodales auctor sodales. Aenean luctus orci euismod lectus bibendum pretium. Curabitur sit amet diam facilisis, sagittis velit a, rutrum mauris. Nam laoreet sem a malesuada tincidunt. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Pellentesque feugiat pulvinar convallis. Aliquam eget justo ex. Suspendisse potenti.
Proin imperdiet ante velit, sed sodales sem ultricies vel. Curabitur at viverra dui. Vestibulum leo ligula, aliquet eu aliquet nec, fringilla nec est. Nullam eleifend non nisi id viverra. Praesent eget tortor a felis laoreet consequat eu eu metus. Etiam vel diam eu sem gravida malesuada vitae vel nisi. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae;
Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Ut consequat sapien id euismod ornare. In ultricies nibh id neque consectetur efficitur. Ut tempor finibus nulla et porta. Duis a massa non enim molestie dictum vehicula at lacus. Phasellus maximus tincidunt neque vel ullamcorper. Praesent orci arcu, mollis vitae tincidunt at, tempor ut sem. Integer turpis eros, elementum eget vulputate eget, pellentesque in libero. Aenean sit amet massa non lacus luctus rutrum.
"""

test_albums = [
    [
        "Koloss",
        "I Let It in and It Took Everything",
        "Thegodmachine",
        "A Tear in the Fabric of Life",
    ],
    [
        "A Tear in the Fabric of Life",
        "Awakened",
        "måsstaden under vatten",
        "Oh What The Future Holds",
    ],
    [
        "Oh What The Future Holds",
        "May Our Chambers Be Full",
        "Oxidized",
        "Absolute",
        "Menschenmühle",
    ],
    [
        "Lifeblood",
        "Melancholy",
        "Laugh Tracks",
    ],
]
for i, subset_albums in enumerate(test_albums):
    title_id = f"test{i}"
    ContentWriter().write_topic(
        title=title_id,
        review_content=body,
        album_names=subset_albums,
    )
