<script>
    export let title;
    export let value = "";
    export let placeholder = "";
    export let disabled = false;
    export let change;
    export let width = 0;
    export let size = 0;
    export let req = false;
</script>

<div class="form-input" style={width > 0 ? "max-width:" + width + "em" : ""}>
    <div style="display: flex;justify-content: start; align-items: center;">
        {#if title}
            <div class="form-label" style="margin: 0; padding-right: 1em;">
                {title}{#if req}<span class="star">*</span>{/if}
            </div>
        {/if}
        <input
            class="input"
            type="number"
            bind:value
            {placeholder}
            {disabled}
            on:input={change}
            style={size > 0 ? "max-width:" + size + "em" : ""}
        />
    </div>
</div>
