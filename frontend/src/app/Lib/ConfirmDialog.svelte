<script>
    import MsgDialog from "./MsgDialog.svelte";

    export let lng = {};
    export const confirm = {
        show(msg, fnYes, fnNo) {
            yesFn = fnYes;
            noFn = fnNo;
            dialog.show(msg);
        },
    };
    let yesFn;
    let noFn;
    let dialog;
</script>

<MsgDialog
    bind:msg={dialog}
    header={lng.confirm}
    icon={{ name: "help-circle" }}
    btns={[
        {
            title: lng.yes,
            class: "primary",
            callback: () => {
                if (yesFn) {
                    yesFn();
                }
            },
        },
        {
            title: lng.no,
            class: "secondary",
            callback: () => {
                if (noFn) {
                    noFn();
                }
            },
        },
    ]}
/>
