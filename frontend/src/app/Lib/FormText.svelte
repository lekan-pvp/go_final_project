<script>
    export let title;
    export let value = "";
    export let size = 5;
    export let req = false;
    export let placeholder = "";
    export let change;
</script>

<div class="form-input" style="width: 100%;">
    {#if title}
        <div class="form-label">
            {title}{#if req}<span class="star">*</span>{/if}
        </div>
    {/if}
    <textarea
        bind:value
        rows={size}
        contenteditable
        class="input"
        {placeholder}
        style="width: 100%"
        on:input={change}
    />
</div>
