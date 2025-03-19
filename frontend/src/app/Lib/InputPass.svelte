<script>
    export let placeholder = "";
    export let value = "";

    let icon = "#eye";
    let input;

    function oneye() {
        if (icon == "#eye") {
            icon = "#eye-off";
            input.type = "text";
        } else {
            icon = "#eye";
            input.type = "password";
        }
    }
</script>

<div style="display:flex;align-items:stretch">
    <div style="flex-grow:1">
        <input
            style="width: 100%;min-height: 100%"
            bind:this={input}
            class="input no-right-radius"
            type="password"
            bind:value
            {placeholder}
        />
    </div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="input-right-btn" on:click={oneye}>
        <svg style="width:1.5em;height: 1.5em;fill:inherit">
            <use xlink:href={icon} />
        </svg>
    </div>
</div>
