<script>
    import { clickOutside } from "../api.js";
    import Icon from "./Icon.svelte";

    export let show = false;
    export let right = false;
    export let bottom = false;
    export let slot = false;
    export let absolute = false;
    export let items = [];
    export let size = "";

    function action(item) {
        show = false;
        if (item.action) {
            item.action(item);
        }
    }
</script>

{#if show}
    <div
        class="dropdown-menu"
        class:right
        class:bottom100={bottom}
        class:top100={!bottom}
        style:position={absolute ? "absolute" : "relative"}
        style:font-size={size ? size : null}
        use:clickOutside
        on:clickoutside={() => (show = false)}
    >
        {#if slot}
            <slot />
            {#if items.length > 0}
                <hr class="dropdown-divider" />
            {/if}
        {/if}
        {#each items as item}
            {#if item.title == "hr"}
                <hr class="dropdown-divider" />
            {:else}
                <a
                    href={null}
                    class="dropdown-item"
                    on:click={() => action(item)}
                >
                    {#if item.icon}
                        <Icon name={item.icon} size="1.3em" />
                    {/if}<span>{item.title}</span></a
                >
            {/if}
        {/each}
    </div>
{/if}
