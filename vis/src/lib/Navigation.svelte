<script>
    import { Nav, NavItem, NavLink } from "sveltestrap";
    import { Styles } from "sveltestrap";
    import { topics } from "$lib/stores";
    let options = [{ uri: "/", name: "Home", sort_index: 0 }];
    $: if ($topics) {
        $topics.forEach((t) => {
            let opt = t;
            opt.uri = `/topics/${t.id}`;
            opt.sort_index = t.id;
            options = [opt, ...options];
        });
        options = [...options.sort((a, b) => a.sort_index - b.sort_index)];
    }
</script>

<Styles />

<h2>Pretense</h2>
<Nav vertical>
    {#each options as opt}
        <NavItem>
            <NavLink sveltekit:prefetch href={opt.uri}>{opt.name}</NavLink>
        </NavItem>
    {/each}
</Nav>
