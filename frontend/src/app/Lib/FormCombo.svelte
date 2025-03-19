<script>
    import { onMount } from "svelte";
    export let value = 0;
    export let list = [];
    export let width = 0;
    export let title = "";
    export let change;
    export let req = false;

    onMount(() => {
        if (typeof value == "string") {
            for (let i = 0; i < list.length; i++) {
                if (value == list[i].id) {
                    value = list[i].id;
                    break;
                }
                if (value == list[i].title) {
                    value = list[i].id;
                    break;
                }
            }
        }
    });
</script>

<div class="form-input" style={width > 0 ? "max-width:" + width + "em" : ""}>
    {#if title}
        <div class="form-label">
            {title}{#if req}<span class="star">*</span>{/if}
        </div>
    {/if}
    <select
        class="form-select"
        bind:value
        on:change={change}
        style="font-size: 1em"
    >
        {#each list as col}
            <option value={col.id}>{col.title}</option>
        {/each}
    </select>
</div>
