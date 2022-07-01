<script>
    import { Nav, NavItem, NavLink } from "sveltestrap";
    import { Styles } from "sveltestrap";
    export let topics = [];
    let options = [];
    $: if (topics) {
        topics.forEach((t) => {
            let opt = t;
            opt.uri = `/topics/${t.id}`;
            options = [opt, ...options];
        });
    }
</script>

<Styles />

<h2>Pretense</h2>
<Nav vertical>
    {#each options as opt}
        <NavItem>
            <NavLink href={opt.uri}>{opt.name}</NavLink>
        </NavItem>
    {/each}
</Nav>
