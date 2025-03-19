<script>
    import { getContext } from "svelte";

    export let placeholder = "";
    export let value = "";
    export let callback;
    export let change;
    export let table;
    export let title;
    export let editable;
    export let fromid = "";

    let ctx = getContext("ctx");
    let lng = ctx.lng;

    function selectLink() {
        ctx.fn.selectInTable({
            table: table,
            fromid: fromid,
            title: title ? title : lng.selectitem,
            callback: (action, par) => {
                if (action == "select") {
                    callback(action, par);
                }
            },
        });
    }
</script>

<div style="display:flex;align-items:stretch">
    <div style="flex-grow:1">
        {#if editable}
            <input
                style="width: 100%;min-height: 100%"
                class="input no-right-radius"
                type="text"
                bind:value
                {placeholder}
                on:input={change}
            />
        {:else}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div
                class="input no-right-radius"
                style="cursor: pointer;min-height: 100%;width: 100%"
                on:click={selectLink}
            >
                {#if value}
                    {value}
                {:else}
                    {lng.selectitem}...
                {/if}
            </div>
        {/if}
    </div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="input-right-btn" on:click={selectLink}>
        <svg style="width:1.5em;height: 1.5em;fill:inherit">
            <use xlink:href="#list-alt" />
        </svg>
    </div>
</div>
