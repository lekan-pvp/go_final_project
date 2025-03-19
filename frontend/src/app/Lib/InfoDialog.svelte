<script>
    import MsgDialog from "./MsgDialog.svelte";

    export let lng = {};
    export const info = {
        show(msg, fCallback) {
            callback = fCallback;
            dialog.show(msg);
        },
    };
    let dialog;
    let callback;
</script>

<MsgDialog
    bind:msg={dialog}
    header={lng.info}
    icon={{ name: "information-slab-circle" }}
    btns={[
        {
            title: lng.ok,
            class: "primary",
            callback: () => {
                if (callback) {
                    callback();
                }
            },
        },
    ]}
/>
