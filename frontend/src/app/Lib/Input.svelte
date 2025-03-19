<script>
    export let placeholder = "";
    export let value = "";
    export let style = null;
    export let disabled = false;
    export let change;
    export let size = 0;
</script>

<input
    class="input"
    type="text"
    bind:value
    maxlength={size > 0 ? size : null}
    {placeholder}
    {style}
    {disabled}
    on:input={change}
/>
