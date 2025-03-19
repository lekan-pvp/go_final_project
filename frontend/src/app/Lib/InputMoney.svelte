<script>
    import { beforeUpdate } from "svelte";
    import Input from "./Input.svelte";

    export let placeholder = "";
    export let value = "";
    export let style = "";
    export let disabled = false;
    export let change;
    export let float = 0;

    function validate(before) {
        let out = "";

        let digits = ["", ""];
        let d = 0;
        let dot = -1;

        for (let i = 0; i < value.length; i++) {
            let ch = value[i];
            if (i == 0 && ch == "0") {
                continue;
            }
            if (ch >= "0" && ch <= "9") {
                digits[d] += ch;
            } else if (ch == "," || ch == ".") {
                if (dot == -1) {
                    d = 1;
                    dot = i;
                }
            } else if (ch == "-" && d == 0 && digits[d].length == 0) {
                digits[0] = "-";
            }
        }
        if (digits[0].length > 13) {
            digits[0] = digits[0].substring(0, 13);
        }
        if (!digits[0]) {
            digits[0] = "0";
        }
        out = digits[0];
        if (float > 0) {
            if (digits[1].length > float) {
                // let last = digits[1][digits[1].length - 1];
                //                if (last == "0") {
                digits[1] = digits[1].substring(0, float);
                /*              } else {
                    digits[1] = digits[1].substring(
                        digits[1].length - float,
                        digits[1].length
                    );
                }*/
            } else {
                digits[1] += "0".repeat(float - digits[1].length);
            }
            out += "." + digits[1];
        }
        if (out !== value) {
            value = out;
        }
        if (change && !before) {
            change();
        }
    }

    beforeUpdate(() => {
        validate(true);
    });

    function onchange() {
        validate(false);
    }
</script>

<Input
    class="input"
    type="text"
    bind:value
    {placeholder}
    {style}
    {disabled}
    change={onchange}
/>
