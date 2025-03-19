<script>
    export let placeholder = "";
    export let value = "";
    export let icon;
    export let callback;
    export let change;
</script>

<div style="display:flex;align-items:stretch">
    <div style="flex-grow:1">
        <input
            style="width: 100%;min-height: 100%"
            class="input no-right-radius"
            type="text"
            bind:value
            {placeholder}
            on:input={change}
        />
    </div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="input-right-btn" on:click={callback}>
        <svg style="width:1.5em;height: 1.5em;fill:inherit">
            <use xlink:href={"#" + icon} />
        </svg>
    </div>
</div>
