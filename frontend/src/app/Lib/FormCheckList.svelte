<script>
    import { beforeUpdate } from "svelte";

    export let title;
    export let values = {};
    export let width = 0;
    export let change;
    export let list = [];

    let items = {};
    let loaded;
</script>

<div class="form-input" style={width > 0 ? "max-width:" + width + "em" : ""}>
    <div class="form-label">{title}</div>
    <div style="display: flex; flex-wrap: wrap;">
        {#each list as item}
            <label class="form-label" style="margin-right: 1em;"
                ><input
                    class="form-check-input"
                    type="checkbox"
                    bind:checked={values[item.id]}
                    on:input={change}
                />{item.title}</label
            >
        {/each}
    </div>
</div>
