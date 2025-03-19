<script>
    import Input from "./Input.svelte";
    import InputPass from "./InputPass.svelte";
    import InputDate from "./InputDate.svelte";
    import InputMoney from "./InputMoney.svelte";
    import InputLink from "./InputLink.svelte";

    export let title;
    export let value = "";
    export let placeholder = "";
    export let disabled = false;
    export let change;
    export let type;
    export let link;
    export let linkedit;
    export let width = 0;
    export let size = 0;
    export let req = false;
    export let callback;
</script>

<div class="form-input" style={width > 0 ? "max-width:" + width + "em" : ""}>
    {#if title}
        <div class="form-label">
            {title}{#if req}<span class="star">*</span>{/if}
        </div>
    {/if}
    {#if type == "password"}
        <InputPass {placeholder} bind:value />
    {:else if type == "date"}
        <InputDate {placeholder} bind:value {change} />
    {:else if type == "money"}
        <InputMoney {placeholder} bind:value float={size} {change} />
    {:else if type == "link"}
        <InputLink
            {placeholder}
            bind:value
            {change}
            table={link}
            {callback}
            fromid={size}
            editable={linkedit}
        />
    {:else}
        <Input {placeholder} bind:value {disabled} {change} {size} />
    {/if}
</div>

<style>
    :global(.form-input) {
        display: flex;
        flex-direction: column;
        margin: 0.5em 0.3em 0em;
    }
    :global(.form-label) {
        color: var(--gray-700);
        margin-bottom: 0.5em;
        font-size: 0.9em;
        font-weight: 600;
    }
</style>
