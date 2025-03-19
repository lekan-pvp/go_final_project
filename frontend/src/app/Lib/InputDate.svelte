<script>
    import InputBtn from "./InputBtn.svelte";
    import DatePicker from "../Lib/DatePicker.svelte";
    import DropDownMenu from "../Lib/DropDownMenu.svelte";

    export let placeholder = "";
    export let value = "";
    export let change;

    let show = false;

    function callback() {
        show = !show;
    }
</script>

<div style="position:relative; max-width: 15em;">
    <InputBtn
        bind:value
        {placeholder}
        icon="calendar-month"
        {callback}
        {change}
    />
    <DropDownMenu absolute right slot bind:show>
        <div style="padding: 0.5em 0.5em;font-size: 1.1em">
            <DatePicker
                bind:value
                callback={() => {
                    if (change) {
                        change();
                    }
                    show = false;
                }}
            />
        </div>
    </DropDownMenu>
</div>
