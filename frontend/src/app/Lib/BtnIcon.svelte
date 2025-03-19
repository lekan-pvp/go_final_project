<script>
    import Icon from "./Icon.svelte";

    export let fn;
    export let name;
    export let size;
    export let color;
</script>

<button class="btnicon" on:click={fn}><Icon {name} {color} {size} /></button>
