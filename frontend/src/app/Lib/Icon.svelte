<script>
    export let name = "";
    export let size = "1.5em";
    export let color = "inherit";
    export let style = "";
</script>

<svg style="width:{size};height: {size};fill:{color};{style}">
    <use xlink:href={"#" + name} />
</svg>
