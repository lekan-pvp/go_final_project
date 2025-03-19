<script>
    export let title;
    export let value = "";
    export let width = 0;
    export let change;
</script>

<div class="form-input" style={width > 0 ? "max-width:" + width + "em" : ""}>
    <label class="form-label"
        ><input
            class="form-check-input"
            type="checkbox"
            bind:checked={value}
            on:input={change}
        />{title}</label
    >
</div>
